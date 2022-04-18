package phpcereal

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type Parser struct {
	Unescape bool
	Head     []byte
	Bytes    []byte
	Type     TypeFlag
	LastRune rune
	Err      error
	//	Pos      int
}

func NewParser[C Chars](s C) *Parser {
	b := []byte(s)
	return &Parser{
		Head:  b,
		Bytes: b,
	}
}

func (p *Parser) EatNext() rune {
	if p.EOF() {
		goto end
	}
	p.advance()
end:
	return p.LastRune
}

func (p *Parser) PeekNext() (r rune) {
	if p.EOF() {
		goto end
	}
	r, _ = p.decode()
end:
	return r
}

func (p *Parser) EatIntUpTo(r rune) (i int) {
	bytes := p.EatUpTo(r)
	if p.Err != nil {
		goto end
	}
	i, p.Err = strconv.Atoi(string(bytes))
	if p.Err != nil {
		p.Err = fmt.Errorf("expected int prior to '%s'; %w", string(r), p.Err)
	}
end:
	return i
}

func (p *Parser) EatLength() (bytes []byte, length int) {
	var err error
	bytes = p.EatUpTo(':') // TODO Change to int,err = p.GetDigitsUpTo(rune)
	length, err = strconv.Atoi(string(bytes))
	if err != nil {
		p.Err = fmt.Errorf("unable to parse '%s'", string(bytes))
	}
	return bytes, length
}

func (p *Parser) MatchTerminatingSemicolonFor(what string) bool {
	if !p.Match(';') {
		p.Err = fmt.Errorf("expected terminating semicolon; got %q", p.LastRune)
		p.Err = fmt.Errorf("error parsing %s; %w", what, p.Err)
		return false
	}
	return true
}

func (p *Parser) Match(r rune, when ...string) (matched bool) {
	var msg string
	ch := p.EatNext()
	if ch == r {
		matched = true
		goto end
	}
	msg = fmt.Sprintf("expected %q but got %q", r, ch)
	if len(when) == 0 {
		goto end
	}
	p.Err = fmt.Errorf("%s %s", msg, when[0])
end:
	return matched
}

func (p *Parser) MatchClosingBraceFor(what string) bool {
	if !p.Match('}') {
		p.Err = fmt.Errorf("expected terminating brace for %s; %w", what, p.Err)
		return false
	}
	return true
}

func (p *Parser) EOF() bool {
	return len(p.Bytes) == 0
}

func (p *Parser) EatUpTo(r rune) (bytes []byte) {
	start := p.Bytes
	count := 0
	for {
		if p.EOF() {

			goto end
		}
		count += p.advance()
		if p.LastRune == r {
			count--
			break
		}
	}
end:
	if p.LastRune != r {
		p.Err = fmt.Errorf("expected %q but reached end", r)
	} else {
		bytes = make([]byte, count)
		copy(bytes, start[:count])
	}
	return bytes
}

func (p *Parser) EatQuotedString(length int, quote rune, quotesEscaped ...bool) (bytes []byte) {
	var inEsc bool
	var quotePos int
	var qe bool

	start := p.Bytes

	count := 0
	if length == 0 {
		goto end
	}
	if len(quotesEscaped) > 0 {
		qe = quotesEscaped[0]
	}
	quotePos = length + 1
	for {
		count += p.advance()
		switch {
		case p.LastRune == quote:
			switch {
			case count == quotePos:
				if qe && !inEsc {
					p.Err = fmt.Errorf("quote found at %d when backslash was expected: %s",
						quotePos, leftTrunc(start, count))
					goto end
				}
				count--
				goto end
			case p.PeekNext() == quote:
				quotePos++
				continue
			}
			//			inQuote = false
			inEsc = false
		case inEsc:
			switch p.LastRune {
			case BackSlash:
				inEsc = true
				continue
			case 'n', 'r', 't':
				count++
			}
			inEsc = false
		case p.LastRune == BackSlash:
			if count == quotePos {
				if !qe {
					p.Err = fmt.Errorf("backslash found at %d when quote was expected: %s",
						quotePos, leftTrunc(start, count))
					goto end
				}
				count--
			} else {
				quotePos++
			}
			inEsc = true
		case count >= quotePos:
			count--
			goto end
		default:
			if p.EOF() {
				p.Err = fmt.Errorf("no closing quote for string: %s", leftTrunc(start, count))
				goto end
			}
		}
	}
end:
	if p.Err == nil {
		bytes = make([]byte, count)
		copy(bytes, start[:count])
	}
	return bytes
}

type void struct{}

var validNodeTypes = map[TypeFlag]void{
	CustomObjTypeFlag:  {},
	NULLTypeFlag:       {},
	ObjectTypeFlag:     {},
	VarRefTypeFlag:     {},
	PHP6StringTypeFlag: {},
	ArrayTypeFlag:      {},
	BoolTypeFlag:       {},
	FloatTypeFlag:      {},
	IntTypeFlag:        {},
	PHP3ObjTypeFlag:    {},
	ObjRefTypeFlag:     {},
	StringTypeFlag:     {},
}

func (p *Parser) EatTypeFlag() TypeFlag {
	var r rune

	tf := TypeFlag(p.EatNext())
	if _, ok := validNodeTypes[tf]; !ok {
		p.Err = fmt.Errorf("invalid node type '%s'", string(tf))
		goto end
	}
	r = p.EatNext()
	switch {
	case r == ':':
		goto end

	case r == ';' && tf == NULLTypeFlag:
		goto end

	default:
		p.Err = fmt.Errorf("expected colon or semi-colon but got '%s'", string(r))
		goto end
	}
end:
	p.Type = tf
	return tf
}

func (p *Parser) Pos() int {
	return len(p.Head) - len(p.Bytes)
}

func (p *Parser) Parse() (cv CerealValue, _ error) {

	tf := p.EatTypeFlag()
	pf, err := GetParseFunc(tf)
	if err != nil {
		p.Err = err
		goto end
	}
	cv = pf(p)

	if pcs, ok := cv.(TypeFlagSetter); ok {
		pcs.SetTypeFlag(tf)
	}
	if p.Err != nil {
		msg := "invalid format at position %d, possible data corruption for value type %q in: %s; %w"
		p.Err = fmt.Errorf(msg, p.Pos(), tf, string(p.Head), p.Err)
	}
end:
	return cv, p.Err
}

type ParseFunc func(p *Parser) CerealValue

func GetParseFunc(tf TypeFlag) (pf ParseFunc, err error) {
	switch tf {
	case ArrayTypeFlag:
		pf = func(p *Parser) CerealValue {
			return ArrayValue{}.Parse(p)
		}

	case NULLTypeFlag:
		pf = func(p *Parser) CerealValue {
			return NullValue{}.Parse(p)
		}

	case StringTypeFlag:
		pf = func(p *Parser) CerealValue {
			return StringValue{}.Parse(p)
		}

	case PHP6StringTypeFlag:
		pf = func(p *Parser) CerealValue {
			return StringValue{}.Parse(p)
		}

	case BoolTypeFlag:
		pf = func(p *Parser) CerealValue {
			return BoolValue{}.Parse(p)
		}

	case IntTypeFlag:
		pf = func(p *Parser) CerealValue {
			return IntValue{}.Parse(p)
		}

	case FloatTypeFlag:
		pf = func(p *Parser) CerealValue {
			return FloatValue{}.Parse(p)
		}

	case ObjectTypeFlag:
		pf = func(p *Parser) CerealValue {
			return ObjectValue{}.Parse(p)
		}

	case CustomObjTypeFlag:
		//pf = func(p *Parser)CerealValue{
		//	return NullValue{}.Parse(p)
		//}

	case ObjRefTypeFlag:
		//pf = func(p *Parser)CerealValue{
		//	return NullValue{}.Parse(p)
		//}

	case VarRefTypeFlag:
		//pf = func(p *Parser)CerealValue{
		//	return NullValue{}.Parse(p)
		//}

	//case PHP3ObjTypeFlag:

	default:

	}
	if pf == nil {
		err = fmt.Errorf("parsing not yet implemented for PHP type %q", tf)
	}

	return pf, err
}

func (p *Parser) decode() (r rune, size int) {
	r, size = rune(p.Bytes[0]), 1
	if r >= utf8.RuneSelf {
		r, size = utf8.DecodeRune(p.Bytes)
		if r == utf8.RuneError {
			switch size {
			case 0:
			case 1:
				p.Err = fmt.Errorf("invalid encoding for UTF-8: got %q", r)
			default:
				p.Err = fmt.Errorf("unexpected size for utf8.DecodeRune() error: %d", size)
			}
			size = 0
			p.Bytes = []byte{}
		}
	}
	return r, size
}

func (p *Parser) advance() int {
	var size int
	p.LastRune, size = p.decode()
	if size == 0 {
		goto end
	}
	p.Bytes = p.Bytes[size:]
end:
	return size
}
