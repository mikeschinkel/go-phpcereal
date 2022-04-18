package phpcereal

import (
	"fmt"
	"strconv"
)

type Char interface {
	[]byte | string
}

func IsCereal[C Char](c C) (is bool) {
	is, _ = _IsCereal(c)
	return is
}

func _IsCereal[C Char](c C) (is bool, n int) {
	var b []byte
	var ok bool
	var p string

	if len(c) <= 2 {
		goto end
	}
	if !isColon(c[1]) {
		goto end
	}
	b = []byte(c)
	p, ok = patterns[TypeFlag(b[0])]
	if !ok {
		goto end
	}
	if isDash(p[0]) {
		panic(fmt.Sprintf("PHP TypeFlag %q not IsCereal() yet", b[0]))
	}
	if isNULL(b) {
		is = true
		goto end
	}
	is, n = isCereal(b[2:], []byte(p))
end:
	return is, n
}

//	#: Repeat
//	L: length(digits)
//	D: digits
//	B: 1 or 0
//	V: CerealValue
//	$: Double-quoted string

var patterns = map[TypeFlag]string{
	NULLTypeFlag:       `N;`,
	ObjectTypeFlag:     `L:$:L:{#}`,
	PHP6StringTypeFlag: `L:$;`,
	ArrayTypeFlag:      "L:{#}",
	BoolTypeFlag:       "B;",
	FloatTypeFlag:      "D.D;",
	IntTypeFlag:        "D;",
	StringTypeFlag:     `L:$;`,
	PHP3ObjTypeFlag:    "-",
	ObjRefTypeFlag:     "-",
	VarRefTypeFlag:     "-",
	CustomObjTypeFlag:  "-",
}

func isCereal(buf, pat []byte) (is bool, bytes int) {
	var index, length, pos, start, n int
	var c byte
	start = pos

	for index, c = range pat {
		if allConsumed(buf, pos) {
			goto end
		}
		switch c {

		case '$': // Double-quoted string
			hasBackslash := false
			if isBackSlash(buf[pos]) {
				hasBackslash = true
				pos++
			}
			if !isDoubleQuote(buf[pos]) {
				goto end
			}
			pos++
			if allConsumed(buf, pos) {
				goto end
			}
			for i := 0; i < length; i++ {
				if isBackSlash(buf[pos]) {
					pos += 2
					continue
				}
				if pos >= len(buf) {
					break
				}
				pos++
				if allConsumed(buf, pos) {
					break
				}
			}
			if hasBackslash {
				if !isBackSlash(buf[pos]) {
					goto end
				} else {
					pos++
				}
			}
			if isDoubleQuote(buf[pos]) {
				pos += 2
				is = true
				goto end
			}

		case 'B': // 1 or 0
			if !isBool(buf[pos]) {
				goto end
			}
			pos++

		case 'V': // CerealValue
			is, n = _IsCereal(buf)
			pos += n + 1

		case '#': // Repeat
			for i := 0; i < length; i++ {
				is, n = _IsCereal(buf[pos:])
				pos += n + 2 // 2 = 1) TypeFlag, 2)Semicolon after index
				if !is {
					goto end
				}
				is, n = _IsCereal(buf[pos:])
				pos += n + 2 // 2 = 1) TypeFlag, 2)Semicolon after index
				if !is {
					goto end
				}
			}

		case 'L': // Length(digits)
			for {
				if isColon(buf[pos]) {
					length, _ = strconv.Atoi(string(buf[start:pos]))
					break
				}
				if !isDigit(buf[pos]) {
					goto end
				}
				pos++
				if allConsumed(buf, pos) {
					break
				}
			}

		case 'D': // Digits
			for {
				pos++
				if allConsumed(buf, pos) {
					// Unexpected
					goto end
				}
				if isSemiColon(buf[pos]) {
					is = true
					pos++
					goto end
				}
				if isPeriod(buf[pos]) {
					// Floating point
					break
				}
				if !isDigit(buf[pos]) {
					// Unexpected
					goto end
				}
			}

		case buf[pos]: // Matched literals in pattern
			pos++

		default: // Mismatch!
			goto end
		}
	}
	is = true
end:
	if len(pat) < index {
		panic(fmt.Sprintf("Validation failed for '%s'", string(buf)))
	}
	return is, pos - start
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}
func isBool(b byte) bool {
	return b == '0' || b == '1'
}
func isSemiColon(b byte) bool {
	return b == ';'
}
func isPeriod(b byte) bool {
	return b == '.'
}
func isColon(b byte) bool {
	return b == ':'
}
func isNULL(b []byte) bool {
	return len(b) >= 2 && b[0] == 'N' && b[1] == ';'
}
func isDash(b byte) bool {
	return b == '-'
}
func isHash(b byte) bool {
	return b == '#'
}
func isBackSlash(b byte) bool {
	return b == '\\'
}
func isDoubleQuote(b byte) bool {
	return b == '"'
}
func isOpenBrace(b byte) bool {
	return b == '{'
}
func allConsumed(b []byte, i int) bool {
	return i == len(b)
}

//type validateFunc func(b []byte) bool
//
//var validateFuncs = map[TypeFlag]validateFunc{
//	CustomObjTypeFlag:  ValidateCustomObj,
//	NULLTypeFlag:       ValidateNULL,
//	ObjectTypeFlag:     ValidateObject,
//	VarRefTypeFlag:     ValidateVarRef,
//	PHP6StringTypeFlag: ValidatePHP6String,
//	ArrayTypeFlag:      ValidateArray,
//	BoolTypeFlag:       ValidateBool,
//	FloatTypeFlag:      ValidateFloat,
//	IntTypeFlag:        ValidateInt,
//	PHP3ObjTypeFlag:    ValidatePHP3Obj,
//	ObjRefTypeFlag:     ValidateObjRef,
//	StringTypeFlag:     ValidateString,
//}
//
//func ValidateCustomObj(b []byte) bool {
//	return true
//}
//
//func ValidateNULL(b []byte) bool {
//	return true
//}
//
//func ValidateObject(b []byte) bool {
//	return true
//}
//
//func ValidateVarRef(b []byte) bool {
//	return true
//}
//
//func ValidatePHP6String(b []byte) bool {
//	return true
//}
//
//func ValidateArray(b []byte) bool {
//	return true
//}
//
//func ValidateBool(b []byte) bool {
//	return true
//}
//
//func ValidateFloat(b []byte) bool {
//	return true
//}
//
//func ValidateInt(b []byte) bool {
//	return true
//}
//
//func ValidatePHP3Obj(b []byte) bool {
//	return true
//}
//
//func ValidateObjRef(b []byte) bool {
//	return true
//}
//
//func ValidateString(b []byte) bool {
//	return true
//}
