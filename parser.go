package phpcereal

import (
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"
)

type Parser struct {
	opts     CerealOpts
	CountCR  bool // Count Carriage Return
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

func (p *Parser) GetOpts() CerealOpts {
	return p.opts
}
func (p *Parser) SetOpts(opts CerealOpts) {
	p.opts = opts
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

func (p *Parser) handleZeroLengthString(quote rune) (err error) {

	advance := func() error {
		size := p.advance()
		if size != 1 {
			err = fmt.Errorf("expected a rune size of 1, got %d", size)
		}
		return err
	}

	err = advance()
	if err != nil {
		goto end
	}
	if p.opts.Escaped {
		if p.LastRune != BackSlash {
			err = fmt.Errorf("expected an escaping backslash, got %q", p.LastRune)
			goto end
		}
		err = advance()
		if err != nil {
			goto end
		}
	}
	if p.LastRune != quote {
		p.Err = fmt.Errorf("expected closing quote but got %q", p.LastRune)
		goto end
	}

end:
	if err != nil {
		p.Err = fmt.Errorf("expected closing quote type %q; %w",
			quote,
			err,
		)
	}
	return err
}

func (p *Parser) EatQuotedString(length int, quote rune, quotesEscaped ...bool) (bytes []byte) {
	var inEsc bool
	var strEnd int
	var qe bool
	var crCount = 0
	var hasEndEsc bool

	start := p.Bytes

	count := 0
	if length == 0 {
		p.Err = p.handleZeroLengthString(quote)
		goto end
	}
	if len(quotesEscaped) > 0 {
		qe = quotesEscaped[0]
	}
	// Add 1 to account for closing quote
	strEnd = length + 1
	for {
		count += p.advance()
		if p.Err != nil {
			goto end
		}
		switch {
		case p.LastRune == quote:
			switch {
			case inEsc:
				inEsc = false
				if hasEndEsc && count+1 == strEnd {
					// Counteract the final count-- below
					count++
					goto end
				}
				if strEnd <= count+2 {
					// Subtract 1 to get rid of `"` from the end of the string
					count--
				}
				if count+1 == strEnd {
					goto end
				}
				if count == strEnd {
					goto end
				}

			case count == strEnd:
				if quote != CloseBrace {
					if qe && !inEsc {
						p.Err = fmt.Errorf("quote found at %d when backslash was expected: %s",
							strEnd, leftTrunc(start, count))
						goto end
					}
					count--
				}
				goto end

			case p.PeekNext() == quote:
				if CloseBrace != quote {
					// CloseBrace occurs in Custom Object
					strEnd++
				}
				continue
			}

		case inEsc:
			switch p.LastRune {
			case BackSlash, 'a', 'e', 'f', 'n', 'R', 't':
				print("")
				switch {
				// For an escaped value at end of escaped string
				// Or for an escaped value at end of non-escaped string
				case p.opts.Escaped && count+1 == strEnd, !p.opts.Escaped && count == strEnd:
					hasEndEsc = true

				}

			case 'r':
				if !p.opts.CountCR {
					// Ignore carriage returns which Adminer adds (e.g. '\r\n') where for
					// the same file mysqldump outputs '\n' for serializations that were
					// created with \n so the string length does not include the \r.
					strEnd++
					crCount++
				}

			case 'c', 'p', 'P', 'x', '0':
				_, _ = p.EatEscapeSequence(p.LastRune)

			}
			inEsc = false

		case p.LastRune == BackSlash:
			inEsc = true
			switch {
			case count == strEnd:
				if !qe {
					p.Err = fmt.Errorf("backslash found at %d when quote was expected: %s",
						strEnd, leftTrunc(start, count))
					goto end
				}
				if hasEndEsc {
					count -= 2
				} else {
					count--
				}
			default:
				strEnd++
			}

		//case count >= strEnd:
		case count > strEnd:
			//			count--
			goto end

		default:
			if p.EOF() {
				p.Err = fmt.Errorf("no closing quote for string: %s", leftTrunc(start, count))
				goto end
			}
		}
	}
end:
	if p.Err == nil && 0 < count {
		if hasEndEsc {
			count--
		}
		bytes = make([]byte, count)
		copy(bytes, start[:count])
	}
	return bytes
}

func (p *Parser) EatControl() (b []byte, count int) {
	return nil, 0
}
func (p *Parser) EatWithUnicodeProperty() (b []byte, count int) {
	return nil, 0
}
func (p *Parser) EatWithoutUnicodeProperty() (b []byte, count int) {
	return nil, 0
}
func (p *Parser) EatHex() (b []byte, count int) {
	return nil, 0
}
func (p *Parser) EatOctal() (b []byte, count int) {
	return nil, 0
}

// EatEscapeSequence eats PHP escape sequences and returns true if it ate one.
// See https://www.php.net/manual/en/regexp.reference.escape.php
// And https://www.php.net/manual/en/regexp.reference.unicode.php
func (p *Parser) EatEscapeSequence(r rune) (b []byte, count int) {
	switch r {
	case 'c':
		b, count = p.EatControl()
	case 'p':
		b, count = p.EatWithUnicodeProperty()
	case 'P':
		b, count = p.EatWithoutUnicodeProperty()
	case 'x':
		b, count = p.EatHex()
	case '0':
		b, count = p.EatOctal()
	}
	return b, count
}

type void struct{}

var validNodeTypes = map[TypeFlag]void{
	CustomObjectTypeFlag: {},
	NULLTypeFlag:         {},
	ObjectTypeFlag:       {},
	VarRefTypeFlag:       {},
	PHP6StringTypeFlag:   {},
	ArrayTypeFlag:        {},
	BoolTypeFlag:         {},
	FloatTypeFlag:        {},
	IntTypeFlag:          {},
	PHP3ObjTypeFlag:      {},
	ObjRefTypeFlag:       {},
	StringTypeFlag:       {},
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
	pf, err := p.GetParseFunc(tf)
	if err != nil {
		p.Err = err
		goto end
	}
	cv = pf(p)

	if pcs, ok := cv.(TypeFlagSetter); ok {
		pcs.SetTypeFlag(tf)
	}
	if p.Err != nil {
		msg := "%w; invalid format at position %d, possible data corruption for value type %q in: %s; "
		p.Err = fmt.Errorf(msg, p.Err, p.Pos(), tf, string(p.Head))
	}
end:
	return cv, p.Err
}

type ParseFunc func(p *Parser) CerealValue

func (p *Parser) GetParseFunc(tf TypeFlag) (pf ParseFunc, err error) {
	switch tf {
	case ArrayTypeFlag:
		pf = func(p *Parser) CerealValue {
			return ArrayValue{opts: p.GetOpts()}.Parse(p)
		}

	case NULLTypeFlag:
		pf = func(p *Parser) CerealValue {
			return NullValue{opts: p.GetOpts()}.Parse(p)
		}

	case StringTypeFlag:
		pf = func(p *Parser) CerealValue {
			return StringValue{opts: p.GetOpts()}.Parse(p)
		}

	case PHP6StringTypeFlag:
		pf = func(p *Parser) CerealValue {
			return StringValue{opts: p.GetOpts()}.Parse(p)
		}

	case BoolTypeFlag:
		pf = func(p *Parser) CerealValue {
			return BoolValue{opts: p.GetOpts()}.Parse(p)
		}

	case IntTypeFlag:
		pf = func(p *Parser) CerealValue {
			return IntValue{opts: p.GetOpts()}.Parse(p)
		}

	case FloatTypeFlag:
		pf = func(p *Parser) CerealValue {
			return FloatValue{opts: p.GetOpts()}.Parse(p)
		}

	case ObjectTypeFlag:
		pf = func(p *Parser) CerealValue {
			return ObjectValue{opts: p.GetOpts()}.Parse(p)
		}

	case CustomObjectTypeFlag:
		pf = func(p *Parser) CerealValue {
			ov := ObjectValue{opts: p.GetOpts()}
			return CustomObjectValue{ObjectValue: ov}.Parse(p)
		}

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
	if len(p.Bytes) == 0 {
		p.Err = errors.New("unexpected end of parse buffer")
		goto end
	}
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
end:
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
