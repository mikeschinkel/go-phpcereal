package phpcereal

import (
	"fmt"
	"strconv"
	"strings"
)

var _ CerealValue = (*CustomObjectValue)(nil)
var _ StringReplacer = (*CustomObjectValue)(nil)

type CustomObjectValue struct {
	ObjectValue
	ArrayValue *ArrayValue
}

func (c CustomObjectValue) GetTypeFlag() TypeFlag {
	return CustomObjectTypeFlag
}
func (c CustomObjectValue) String() string {
	builder := strings.Builder{}
	builder.WriteString(string(c.Value.ClassName))
	builder.WriteString(c.ArrayValue.String())
	return builder.String()
}

func (c *CustomObjectValue) Serialized() string {
	if !c.BytesSet() {
		builder := strings.Builder{}
		builder.WriteByte(byte(CustomObjectTypeFlag))
		builder.WriteByte(':')
		name := string(c.Value.ClassName)
		builder.WriteString(strconv.Itoa(len(name)))
		if c.escaped {
			builder.WriteString(`:\"`)
		} else {
			builder.WriteString(`:"`)
		}
		builder.WriteString(name)
		if c.escaped {
			builder.WriteString(`\":`)
		} else {
			builder.WriteString(`":`)
		}
		c.ArrayValue.SetEscaped(c.escaped)
		builderWriteInt(&builder, c.ArrayValue.SerializedLen())
		builder.WriteString(":{")
		builder.WriteString(c.ArrayValue.Serialized())
		builder.WriteByte('}')
		c.SetBytes([]byte(builder.String()))
	}
	return string(c.bytes)
}

func (c CustomObjectValue) SerializedLen() int {
	return unescapedLength(c.Serialized())
}

func (c CustomObjectValue) Parse(p *Parser) (_ CerealValue) {
	var arrBytes []byte
	var ap *Parser
	var err error
	var cv CerealValue
	var ok bool
	var r rune

	length, _ := c.ParseHeader(p)
	if p.Err != nil {
		goto end
	}

	r = p.PeekNext()
	if !p.Match(OpenBrace, "containing custom object array") {
		p.Err = fmt.Errorf("error parsing custom object array; expected open brace, got %q", r)
		goto end
	}

	arrBytes = p.EatQuotedString(length, CloseBrace, p.Escaped)
	if p.Err != nil || len(arrBytes) == 0 {
		p.Err = fmt.Errorf("error parsing custom object array as string; %w", p.Err)
		goto end
	}
	ap = NewParser(arrBytes)
	ap.Escaped = p.Escaped
	cv, err = ap.Parse()
	if err != nil {
		p.Err = fmt.Errorf("error parsing custom object array; %w", err)
		goto end
	}
	c.ArrayValue, ok = cv.(*ArrayValue)
	if !ok {
		p.Err = fmt.Errorf("error parsing custom object; expected 'a' got %q", arrBytes[0])
		goto end
	}
end:
	if p.Err != nil {
		if c.Value.ClassName == "" {
			p.Err = fmt.Errorf("error parsing object; %w", p.Err)
		} else {
			p.Err = fmt.Errorf("error parsing class %s; %w", c.Value.ClassName, p.Err)
		}
	}
	return &c // This must be a pointer for StringReplacer to work
}
