package phpcereal

import "fmt"

var _ CerealValue = (*FloatValue)(nil)

type FloatValue struct {
	Integer  int
	Fraction int
}

func (v FloatValue) GetValue() interface{} {
	return float64(v.Integer) + float64(v.Fraction)/100
}

func (v FloatValue) GetType() PHPType {
	return "float"
}

func (v FloatValue) GetTypeFlag() TypeFlag {
	return FloatTypeFlag
}

func (v FloatValue) String() string {
	return fmt.Sprintf("%d.%d", v.Integer, v.Fraction)
}

func (v FloatValue) Serialized() string {
	return fmt.Sprintf("d:%s;", v.String())
}
func (v FloatValue) SerializedLen() int {
	return 3 + numDigits(v.Integer) + numDigits(v.Fraction)
}

func (v FloatValue) Parse(p *Parser) (_ CerealValue) {
	var i, f int

	i = p.EatIntUpTo('.')
	if p.Err != nil {
		goto end
	}
	f = p.EatIntUpTo(';')
	if p.Err != nil {
		goto end
	}
	v.Integer = i
	v.Fraction = f
end:
	return v
}
