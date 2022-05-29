package phpcereal

import (
	"fmt"
	"strings"
)

var _ CerealValue = (*ArrayValue)(nil)
var _ StringReplacer = (*ArrayValue)(nil)
var _ SQLSerializedGetter = (*ArrayValue)(nil)

type boolMappedBytes map[bool][]byte
type ArrayValue struct {
	Value    Array
	LenBytes []byte // Array Length but as []byte vs. int
	bytes    boolMappedBytes
	escaped  bool
}

func (v *ArrayValue) Length() int {
	return v.Value.Length()
}

func (v *ArrayValue) ReplaceString(from, to string, times int) {
	for _, e := range v.Value {
		sr, ok := e.Value.(StringReplacer)
		if !ok {
			continue
		}
		sr.ReplaceString(from, to, times)
	}
}

func (v ArrayValue) GetBytes() []byte {
	if v.bytes == nil {
		return nil
	}
	return v.bytes[v.escaped]
}

func (v ArrayValue) GetValue() interface{} {
	return v.Value
}

func (v ArrayValue) GetType() PHPType {
	return "array"
}

func (v ArrayValue) GetTypeFlag() TypeFlag {
	return ArrayTypeFlag
}

func (v ArrayValue) String() string {
	return v.Value.String()
}

func (v *ArrayValue) SQLSerialized() string {
	return v.serialized(true)
}

func (v *ArrayValue) Serialized() (s string) {
	return v.serialized(false)
}

func (v *ArrayValue) serialized(escaped bool) (s string) {
	var builder strings.Builder
	if v.bytes == nil {
		v.bytes = boolMappedBytes{
			true:  nil,
			false: nil,
		}
		v.escaped = escaped
	}
	if v.bytes[escaped] != nil {
		goto end
	}
	builder = strings.Builder{}
	builder.WriteByte(byte(ArrayTypeFlag))
	builder.WriteByte(':')
	builderWriteInt(&builder, v.Value.Length())
	builder.WriteString(":{")
	for _, element := range v.Value {
		if escaped {
			// If escaped==true, e.g. generate serialized for SQL
			// then the element *may* need to be serialized.
			builder.WriteString(MaybeGetSQLSerialized(element.Key))
			builder.WriteString(MaybeGetSQLSerialized(element.Value))
		} else {
			// If escaped==false, e.g. do not
			// generate serialized for SQL.
			builder.WriteString(element.Key.Serialized())
			builder.WriteString(element.Value.Serialized())
		}
	}
	builder.WriteByte('}')
	v.bytes[escaped] = []byte(builder.String())
end:
	return string(v.bytes[escaped])
}

func (v ArrayValue) SerializedLen() int {
	return len(v.Serialized())
}

func (v ArrayValue) Parse(p *Parser) (_ CerealValue) {
	var elements []ArrayElement
	var length int
	var lenBytes []byte

	lenBytes, length = p.EatLength()
	if p.Err != nil {
		p.Err = fmt.Errorf("invalid array length; %w", p.Err)
		goto end
	}

	elements = make(Array, length)
	if !p.Match('{') {
		goto end
	} else {
		for index, element := range elements {
			pf, err := GetParseFunc(p.EatTypeFlag())
			if err != nil {
				p.Err = err
				goto end
			}
			element.Key = pf(p)
			if p.Err != nil {
				p.Err = fmt.Errorf("error parsing array key #%d; %w", index, p.Err)
				goto end
			}
			pf, err = GetParseFunc(p.EatTypeFlag())
			if err != nil {
				p.Err = err
				goto end
			}
			element.Value = pf(p)
			if p.Err != nil {
				p.Err = fmt.Errorf("error parsing array value #%d; %w", index, p.Err)
				goto end
			}
			elements[index] = element
		}
		if !p.MatchClosingBraceFor("array") {
			goto end
		}
	}
	v.Value = elements
	v.LenBytes = lenBytes
end:
	return &v // This must be a pointer for StringReplacer to work
}
