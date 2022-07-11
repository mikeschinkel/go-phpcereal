package phpcereal

import (
	"fmt"
	"strconv"
	"strings"
)

var _ CerealValue = (*StringValue)(nil)
var _ TypeFlagSetter = (*StringValue)(nil)
var _ StringReplacer = (*StringValue)(nil)

type StringValue struct {
	escaped      bool
	TypeFlag     TypeFlag
	Value        string
	LengthBytes  []byte
	Length       int
	quoteType    rune
	escapedValue string
}

func (v *StringValue) ReplaceString(from, to string, times int) {
	if len(v.Value) >= len(from) {
		v.Value = strings.Replace(v.Value, from, to, times)
	}
}

func (v StringValue) GetValue() interface{} {
	return v.Value
}

func (v StringValue) GetType() PHPType {
	var t PHPType
	switch v.GetTypeFlag() {
	case StringTypeFlag:
		t = "string"
	case PHP6StringTypeFlag:
		t = "6string"
	}
	return t
}

func (v *StringValue) SetTypeFlag(pc TypeFlag) {
	v.TypeFlag = pc
}

func (v StringValue) GetTypeFlag() TypeFlag {
	if v.TypeFlag == 0 {
		v.TypeFlag = StringTypeFlag
	}
	return v.TypeFlag
}

func (v StringValue) GetEscaped() bool {
	return v.escaped
}

func (v *StringValue) SetEscaped(e bool) {
	v.escaped = e
}

func (v StringValue) String() string {
	var pattern = `"%s"`
	if v.escaped {
		pattern = `\"%s\"`
	}
	return fmt.Sprintf(pattern, v.getEscapedValue())
}

func (v *StringValue) getEscapedValue() string {
	if len(v.Value) == 0 {
		return v.Value
	}
	if v.escapedValue == "" {
		v.escapedValue = escape(v.Value)
	}
	return v.escapedValue
}

func (v StringValue) Serialized() string {
	var pattern string
	if v.escaped {
		pattern = `%c:%d:\"%s\";`
	} else {
		pattern = `%c:%d:"%s";`
	}
	return fmt.Sprintf(pattern,
		byte(v.GetTypeFlag()),
		unescapedLength(v.Value),
		v.getEscapedValue())
}

func (v StringValue) SerializedLen() int {
	length := len(v.getEscapedValue())
	return 2 + numDigits(length) + 3 + length
}

func (v StringValue) Parse(p *Parser) (_ CerealValue) {
	var bytes []byte
	var err error
	var r rune
	var quotesEscaped bool

	v.LengthBytes = p.EatUpTo(':')
	v.Length, err = strconv.Atoi(string(v.LengthBytes))
	if err != nil {
		p.Err = fmt.Errorf("invalid string length; got '%s'", string(v.LengthBytes))
		goto end
	}
	r = p.EatNext()
	if r == BackSlash {
		r = p.EatNext()
		quotesEscaped = true
	}

	v.quoteType = r
	switch v.quoteType {
	case SingleQuote, DoubleQuote:
		bytes = p.EatQuotedString(v.Length, v.quoteType, quotesEscaped)
	default:
		p.Err = fmt.Errorf("expected opening quote for string; got %q", v.quoteType)
		goto end
	}
	quotesEscaped = false
	if !p.MatchTerminatingSemicolonFor("string") {
		goto end
	}
	v.Value = string(bytes)
end:
	// This is a pointer so that cv.(TypeFlagSetter) will work
	return &v
}
