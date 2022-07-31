package phpcereal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var _ CerealValue = (*ObjectValue)(nil)
var _ StringReplacer = (*ObjectValue)(nil)

type ObjectValue struct {
	opts     CerealOpts
	escaped  bool
	Value    Object
	LenBytes []byte // Array Length but as []byte vs. int
	bytes    []byte
}

func (v ObjectValue) GetOpts() CerealOpts {
	return v.opts
}
func (v ObjectValue) SetOpts(opts CerealOpts) {
	v.opts = opts
}
func (v ObjectValue) ReplaceString(from, to string, times int) {
	v.Value.ReplaceString(from, to, times)
	v.bytes = nil
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

func (v ObjectValue) GetEscaped() bool {
	return v.opts.Escaped
}

func (v *ObjectValue) SetEscaped(e bool) {
	if v.GetEscaped() != e {
		v.opts.Escaped = e
		v.bytes = nil
	}
}

func (v ObjectValue) GetBytes() []byte {
	return v.bytes
}

func (v *ObjectValue) SetBytes(b []byte) {
	v.bytes = b
}

func (v *ObjectValue) BytesSet() bool {
	return v.bytes != nil
}

func (v ObjectValue) Serialized() string {
	writeQuote := func(parts *strings.Builder) {
		if v.GetEscaped() {
			parts.WriteString(`\"`)
		} else {
			parts.WriteByte('"')
		}
	}
	if v.bytes == nil {
		parts := strings.Builder{}
		parts.WriteByte(byte(ObjectTypeFlag))
		parts.WriteByte(':')
		name := string(v.Value.ClassName)
		parts.WriteString(strconv.Itoa(len(name)))
		parts.WriteByte(':')
		writeQuote(&parts)
		parts.WriteString(name)
		writeQuote(&parts)
		parts.WriteByte(':')
		builderWriteInt(&parts, v.Value.Size())
		parts.WriteString(":{")
		for _, prop := range v.Value.Properties {
			parts.WriteString(fmt.Sprintf(`%c:%d:%s;`,
				byte(StringTypeFlag),
				prop.GetNameLength(),
				prop.GetQuotedName(),
			))
			parts.WriteString(prop.Value.Serialized())
		}
		parts.WriteByte('}')
		v.bytes = []byte(parts.String())
	}
	return string(v.bytes)
}

func (v ObjectValue) SerializedLen() (length int) {
	return unescapedLength(v.Serialized(), v.opts)
}

func (v *ObjectValue) ParseHeader(p *Parser) (length int, lenBytes []byte) {
	var r rune
	var nameLen int
	var nameBytes []byte

	nameLen = p.EatIntUpTo(':')
	if p.Err != nil {
		p.Err = fmt.Errorf("invalid object class name length; %w", p.Err)
		goto end
	}

	r = p.EatNext()
	if v.GetEscaped() {
		if r != BackSlash {
			p.Err = fmt.Errorf("expected backslash to escape quoted class name, got %q", r)
			goto end
		}
		r = p.EatNext()
	}

	if r != DoubleQuote {
		p.Err = fmt.Errorf("expected enquoted object class name, got %q", r)
		goto end
	}

	nameBytes = p.EatQuotedString(nameLen, DoubleQuote, v.GetEscaped())
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

	if !p.Match('{') {
		goto end
	} else {
		v.bytes = nil
		props = make(ObjectProperties, length)
		for index, prop := range props {
			prop.SetEscaped(v.GetEscaped())
			prop.Parse(p)
			if p.Err != nil {
				p.Err = fmt.Errorf("error parsing property '%s' [#%d] of class %s; %w",
					prop.GetName(), index+1, v.Value.ClassName, p.Err)
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
