package phpcereal

import "fmt"

var _ CerealValue = (*BoolValue)(nil)

type BoolValue struct {
	Value bool
}

func (v BoolValue) GetValue() interface{} {
	return v.Value
}

func (v BoolValue) GetType() PHPType {
	return "bool"
}

func (v BoolValue) GetTypeFlag() TypeFlag {
	return BoolTypeFlag
}

func (v BoolValue) String() string {
	if v.Value {
		return "true"
	}
	return "false"
}

func (v BoolValue) Serialized() string {
	var tf int
	if v.Value {
		tf = 1
	}
	return fmt.Sprintf(`b:%d;`, tf)
}
func (v BoolValue) SerializedLen() int {
	return 3
}

func (v BoolValue) Parse(p *Parser) (_ CerealValue) {
	b := p.EatUpTo(';')
	if p.Err != nil {
		p.Err = fmt.Errorf("expected boolean '1' or '0'; %w", p.Err)
		goto end
	}
	v.Value = b[0] == '1'
end:
	return v
}
