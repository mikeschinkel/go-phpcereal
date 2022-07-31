package phpcereal

import (
	"math"
	"strconv"
	"strings"
)

func MightBeCereal[C []byte | string](c C) bool {
	return c[1] == ':'
}

func numDigits(n int) int {
	return int(math.Floor(math.Log10(float64(n)) + 1))
}

func escape(s string) string {
	s = strings.Replace(s, `\"`, `"`, -1)
	return strings.Replace(s, `"`, `\"`, -1)
}

func builderWriteInt(b *strings.Builder, i int) {
	if i < 10 {
		b.WriteByte(byte(i + '0'))
	} else {
		b.WriteString(strconv.Itoa(i))
	}
}

func stringLengthIgnoreNulls(s string) (n int) {
	n = len(s)
	for _, ch := range s {
		if ch != '0' {
			continue
		}
		n--
	}
	return n
}

func leftTrunc[C Chars](s C, n int) string {
	_s := string(s)
	if len(_s) >= n {
		return _s
	}
	return _s[:n]
}

func unescapedLength(value string, opts ...CerealOpts) int {
	length := len(value)
	lastEle := length - 1
	escaped := false
	if opts == nil {
		opts = make([]CerealOpts, 1)
	}
	for i := lastEle; i >= 0; i-- {
		if value[i] == '\\' {
			if !escaped {
				length--
				escaped = true
			}
			if i < lastEle && value[i+1] == 'r' && !opts[0].CountCR {
				length--
				escaped = false
			}
			continue
		}
		escaped = false
	}
	return length
}
