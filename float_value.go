package phpcereal

import (
	"fmt"
)

var _ CerealValue = (*FloatValue)(nil)

type FloatValue struct {
	escaped      bool
	Integer      int
	Fraction     int
	omitFraction bool
	serializing  bool
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

func (v FloatValue) GetEscaped() bool {
	return v.escaped
}

func (v *FloatValue) SetEscaped(e bool) {
	v.escaped = e
}

func (v FloatValue) String() string {
	if v.omitFraction && v.serializing {
		return fmt.Sprintf("%d", v.Integer)
	}
	return fmt.Sprintf("%d.%d", v.Integer, v.Fraction)
}

func (v FloatValue) Serialized() (s string) {
	v.serializing = true
	s = fmt.Sprintf("d:%s;", v.String())
	v.serializing = false
	return
}
func (v FloatValue) SerializedLen() (n int) {
	n = 3 + numDigits(v.Integer)
	if !v.omitFraction {
		n += numDigits(v.Fraction)
	}
	return n
}

func (v FloatValue) Parse(p *Parser) (_ CerealValue) {
	var i, f int

	b := p.Bytes
	i = p.EatIntUpTo('.')
	if p.Err != nil {
		// Sometimes floats are just `d:<i>;` instead of `d:<i>:<f>;`
		p.Err = nil
		p.Bytes = b
		iv, _ := IntValue{}.Parse(p).(*IntValue)
		if p.Err != nil {
			goto end
		}
		v.Integer = iv.Value
		v.omitFraction = true
		goto end
	}
	f = p.EatIntUpTo(';')
	if p.Err != nil {
		goto end
	}
	v.Integer = i
	v.Fraction = f
end:
	return &v // This is a pointer to allow return values to call pointer interfaces
}
