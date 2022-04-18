package phpcereal

import "strings"

type Array []ArrayElement

func (a Array) Length() int {
	return len(a)
}

func (a Array) String() (s string) {
	b := strings.Builder{}
	b.WriteByte('[')
	for _, e := range a {
		_s := e.String()
		b.WriteString(_s)
	}
	b.WriteByte(']')
	return b.String()
}
