package phpcereal

import (
	"fmt"
	"strconv"
)

var _ CerealValue = (*IntValue)(nil)

type IntValue struct {
	escaped bool
	Value   int
}

func (v IntValue) GetEscaped() bool {
	return v.escaped
}

func (v *IntValue) SetEscaped(e bool) {
	v.escaped = e
}

func (v IntValue) GetValue() interface{} {
	return v.Value
}

func (v IntValue) GetType() PHPType {
	return "int"
}

func (v IntValue) GetTypeFlag() TypeFlag {
	return IntTypeFlag
}

func (v IntValue) String() string {
	return strconv.Itoa(v.Value)
}

func (v IntValue) Serialized() string {
	return fmt.Sprintf("i:%d;", v.Value)
}
func (v IntValue) SerializedLen() int {
	return 2 + numDigits(v.Value)
}

func (v IntValue) Parse(p *Parser) (_ CerealValue) {
	i := p.EatIntUpTo(';')
	if p.Err != nil {
		p.Err = fmt.Errorf("expected integer value; %w", p.Err)
		goto end
	}
	v.Value = i
end:
	return &v // This is a pointer to allow return values to call pointer interfaces
}
