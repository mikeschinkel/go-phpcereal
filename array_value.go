package phpcereal

import (
	"fmt"
	"strings"
)

var _ ValueAccessor = (*ArrayValue)(nil)

type ArrayValue struct {
	Value    Array
	LenBytes []byte // Array Length but as []byte vs. int
	Bytes    []byte
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

func (v ArrayValue) Serialized() (s string) {
	if v.Bytes == nil {
		parts := strings.Builder{}
		parts.WriteByte(byte(ArrayTypeFlag))
		parts.WriteByte(':')
		builderWriteInt(&parts, v.Value.Length())
		parts.WriteString(":{")
		for _, element := range v.Value {
			parts.WriteString(element.Key.Serialized())
			parts.WriteString(element.Value.Serialized())
		}
		parts.WriteByte('}')
		v.Bytes = []byte(parts.String())
	}
	return string(v.Bytes)
}

func (v ArrayValue) SerializedLen() int {
	return len(v.Serialized())
}

func (v ArrayValue) Parse(p *Parser) (_ ValueAccessor) {
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
	return v
}
