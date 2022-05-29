package phpcereal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var _ CerealValue = (*ObjectValue)(nil)
var _ StringReplacer = (*ObjectValue)(nil)
var _ SQLSerializedGetter = (*ObjectValue)(nil)

type ObjectValue struct {
	Value    Object
	LenBytes []byte // Array Length but as []byte vs. int
	Bytes    []byte
}

func (v ObjectValue) ReplaceString(from, to string, times int) {
	v.Value.ReplaceString(from, to, times)
}

func (v ObjectValue) GetValue() interface{} {
	return v.Value
}

func (v ObjectValue) GetType() PHPType {
	return v.Value.ClassName
}

func (v ObjectValue) GetTypeFlag() TypeFlag {
	return ObjectTypeFlag
}

func (v ObjectValue) String() string {
	return v.Value.String()
}

func (v ObjectValue) GetValueType() TypeFlag {
	return ObjectTypeFlag
}

func (v ObjectValue) Serialized() string {
	return v.serialized(false)
}

func (v *ObjectValue) SQLSerialized() string {
	return v.serialized(true)
}

func (v ObjectValue) serialized(sql bool) string {
	if v.Bytes == nil {
		parts := strings.Builder{}
		parts.WriteByte(byte(ObjectTypeFlag))
		parts.WriteByte(':')
		name := string(v.Value.ClassName)
		parts.WriteString(strconv.Itoa(len(name)))
		parts.WriteString(`:"`)
		parts.WriteString(name)
		parts.WriteString(`":`)
		builderWriteInt(&parts, v.Value.Size())
		parts.WriteString(":{")
		for _, prop := range v.Value.Properties {
			parts.WriteString(fmt.Sprintf(`%c:%d:"%s";`,
				byte(StringTypeFlag),
				stringLengthIgnoreNulls(escape(prop.Name)),
				prop.maybeGetSQLName(sql)))
			if sql {
				// If sql==true, e.g. generate serialized for SQL
				// then the element *may* need to be serialized.
				parts.WriteString(MaybeGetSQLSerialized(prop.Value))
			} else {
				// If sql==false, e.g. do not
				// generate serialized for SQL.
				parts.WriteString(prop.Value.Serialized())
			}
		}
		parts.WriteByte('}')
		v.Bytes = []byte(parts.String())
	}
	return string(v.Bytes)
}

func (v ObjectValue) SerializedLen() int {
	return len(v.Serialized())
}

func (v *ObjectValue) ParseHeader(p *Parser) (length int, lenBytes []byte) {
	var r rune
	var quotesEscaped bool
	var nameLen int
	var nameBytes []byte

	nameLen = p.EatIntUpTo(':')
	if p.Err != nil {
		p.Err = fmt.Errorf("invalid object class name length; %w", p.Err)
		goto end
	}

	r = p.PeekNext()
	if r == BackSlash {
		r = p.EatNext()
		quotesEscaped = true
	}

	if !p.Match('"', "enquoting object class name") {
		goto end
	}

	nameBytes = p.EatQuotedString(nameLen, '"', quotesEscaped)
	if nameBytes == nil {
		p.Err = errors.New("error; empty object class name")
		goto end
	}

	if nameLen != len(nameBytes) {
		p.Err = fmt.Errorf("mismatch in class name length; %d vs. %d (possible data corruption)",
			nameLen, len(nameBytes))
		goto end
	}

	v.Value.ClassName = PHPType(nameBytes)

	if !p.Match(':', "after object class name") {
		goto end
	}

	lenBytes, length = p.EatLength()
	if p.Err != nil {
		p.Err = fmt.Errorf("invalid object size; %w", p.Err)
		goto end
	}
end:
	return length, lenBytes
}

func (v ObjectValue) Parse(p *Parser) (_ CerealValue) {
	var props ObjectProperties

	length, lenBytes := v.ParseHeader(p)
	if p.Err != nil {
		goto end
	}

	props = make(ObjectProperties, length)
	if !p.Match('{') {
		goto end
	} else {
		for index, prop := range props {
			prop.Parse(p)
			if p.Err != nil {
				p.Err = fmt.Errorf("error parsing property #%d of class %s; %w",
					index, v.Value.ClassName, p.Err)
				goto end
			}
			props[index] = prop
		}
		if !p.MatchClosingBraceFor("object") {
			goto end
		}
	}
	v.Value.Properties = props
	v.LenBytes = lenBytes
end:
	if p.Err != nil {
		if v.Value.ClassName == "" {
			p.Err = fmt.Errorf("error parsing object; %w", p.Err)
		} else {
			p.Err = fmt.Errorf("error parsing class %s; %w", v.Value.ClassName, p.Err)
		}
	}
	return &v // This must be a pointer for StringReplacer to work
}
