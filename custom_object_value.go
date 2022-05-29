package phpcereal

import (
	"fmt"
	"strconv"
	"strings"
)

var _ CerealValue = (*CustomObjectValue)(nil)
var _ StringReplacer = (*CustomObjectValue)(nil)
var _ SQLSerializedGetter = (*CustomObjectValue)(nil)

type CustomObjectValue struct {
	ObjectValue
	ArrayValue *ArrayValue
}

func (c CustomObjectValue) GetTypeFlag() TypeFlag {
	return CustomObjectTypeFlag
}
func (c CustomObjectValue) GetValueType() TypeFlag {
	return CustomObjectTypeFlag
}
func (c CustomObjectValue) Serialized() string {
	return c.serialized(false)
}

func (c CustomObjectValue) String() string {
	builder := strings.Builder{}
	builder.WriteString(string(c.Value.ClassName))
	builder.WriteString(c.ArrayValue.String())
	return builder.String()
}

func (c *CustomObjectValue) SQLSerialized() string {
	return c.serialized(true)
}

func (c CustomObjectValue) serialized(escaped bool) string {
	panic("TODO: use the escaped parameter")
	if c.Bytes == nil {
		builder := strings.Builder{}
		builder.WriteByte(byte(CustomObjectTypeFlag))
		builder.WriteByte(':')
		name := string(c.Value.ClassName)
		builder.WriteString(strconv.Itoa(len(name)))
		builder.WriteString(`:"`)
		builder.WriteString(name)
		builder.WriteString(`":`)
		builderWriteInt(&builder, c.ArrayValue.SerializedLen())
		builder.WriteString(":{")
		builder.WriteString(c.ArrayValue.Serialized())
		builder.WriteByte('}')
		c.Bytes = []byte(builder.String())
	}
	return string(c.Bytes)
}

func (c CustomObjectValue) SerializedLen() int {
	return len(c.Serialized())
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

	arrBytes = p.EatQuotedString(length, CloseBrace, p.Unescape)
	if p.Err != nil || len(arrBytes) == 0 {
		p.Err = fmt.Errorf("error parsing custom object array as string; %w", p.Err)
		goto end
	}
	ap = NewParser(arrBytes)
	ap.Unescape = p.Unescape
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
