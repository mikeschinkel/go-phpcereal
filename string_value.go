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
	TypeFlag    TypeFlag
	Value       string
	LengthBytes []byte
	Length      int
	quoteType   rune
	quotedValue string
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

func (v StringValue) String() string {
	return fmt.Sprintf(`"%s"`, v.getEscapedValue())
}

func (v *StringValue) getEscapedValue() string {
	if v.quotedValue == "" {
		v.quotedValue = escape(v.Value)
	}
	return v.quotedValue
}

func (v StringValue) Serialized() string {
	return v.serialized(false)
}

func (v StringValue) SQLSerialized() string {
	return v.serialized(true)
}

func (v StringValue) serialized(sql bool) string {
	pattern := `%c:%d:"%s";`
	if sql {
		pattern = `%c:%d:\"%s\";`
	}
	return fmt.Sprintf(pattern,
		byte(v.GetTypeFlag()),
		len(v.Value),
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
