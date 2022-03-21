package phpcereal

import "strings"

type Array []ArrayElement

func (a Array) Length() int {
	return len(a)
}

func (a Array) String() (s string) {
	b := strings.Builder{}
	b.WriteByte('[')
	end := len(a) - 1
	for i, e := range a {
		_s := e.String()
		if i == end && end >= 0 {
			b.WriteString(_s[:len(_s)-1])
		} else {
			b.WriteString(_s)
		}
	}
	b.WriteByte(']')
	return b.String()
}
