package phpcereal

import "strings"

var _ StringReplacer = (*Object)(nil)

type ObjectProperties []ObjectProperty

type Object struct {
	ClassName  PHPType
	Properties ObjectProperties
}

func (o *Object) ReplaceString(from, to string, times int) {
	for _, p := range o.Properties {
		sr, ok := p.Value.(StringReplacer)
		if !ok {
			continue
		}
		sr.ReplaceString(from, to, times)
	}
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
	s = b.String()
	return s
}
