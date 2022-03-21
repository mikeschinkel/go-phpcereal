package phpcereal

import "strings"

type Object struct {
	ClassName  PHPType
	Properties []ObjectProperty
}

func (o Object) Size() int {
	return len(o.Properties)
}
func (o Object) String() (s string) {
	b := strings.Builder{}
	b.WriteString(string(o.ClassName))
	b.WriteByte('{')
	end := len(o.Properties) - 1
	for i, p := range o.Properties {
		_s := p.String()
		if i == end && end >= 0 {
			b.WriteString(_s[:len(_s)-1])
		} else {
			b.WriteString(_s)
		}
	}
	b.WriteByte('}')
	return b.String()
}
