package phpcereal

import (
	"fmt"
)

type Visibility int8

const (
	Public Visibility = iota
	Protected
	Private
)

type ObjectProperty struct {
	escaped    bool
	name       string
	Visibility Visibility
	Value      CerealValue
}

func (prop ObjectProperty) GetName() string {
	return prop.name
}

func (prop ObjectProperty) GetNameLength() int {
	return stringLengthIgnoreNulls(prop.name)
}

func (prop ObjectProperty) GetQuotedName() string {
	if prop.escaped {
		return fmt.Sprintf(`\"%s\"`, prop.name)
	}
	return fmt.Sprintf(`"%s"`, prop.name)
}

func (prop ObjectProperty) GetEscaped() bool {
	return prop.escaped
}

func (prop *ObjectProperty) SetEscaped(e bool) {
	prop.escaped = e
}

// NonPublicName returns just the name and not the other stuff for private and protected vars
//
// 	Private:  \0className\0propName
// 	Protected:  \0*\0propName
//
func (prop ObjectProperty) NonPublicName() (s string) {
	s = prop.name
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] != '0' {
			continue
		}
		s = s[i+1:]
		break
	}
	return s
}

func (prop ObjectProperty) String() (s string) {
	var modifier string
	switch prop.Visibility {
	case Public:
		modifier = ""
	case Protected:
		modifier = "^"
	case Private:
		modifier = "~"
	}
	return fmt.Sprintf("%s%s:%s,", modifier, prop.NonPublicName(), prop.Value.String())
}

func (prop *ObjectProperty) Parse(p *Parser) {
	var pf ParseFunc
	var err error
	var cv CerealValue

	nameTypeFlag := p.EatTypeFlag()
	if nameTypeFlag != StringTypeFlag {
		msg := "invalid prefix character for object property name; expected %q, got %q"
		p.Err = fmt.Errorf(msg, StringTypeFlag, nameTypeFlag)
		goto end
	}
	pf, err = p.GetParseFunc(nameTypeFlag)
	if err != nil {
		p.Err = err
		goto end
	}
	cv = pf(p)
	if p.Err != nil {
		p.Err = fmt.Errorf("error parsing property name; %w", p.Err)
		goto end
	}
	prop.name = cv.GetValue().(string)
	if prop.name == "" {
		p.Err = fmt.Errorf("object property name is empty")
		goto end
	}
	if prop.name[0:2] == "\\0" {
		if len(prop.name) <= 3 {
			p.Err = fmt.Errorf("truncated object property name")
			goto end
		}
		if prop.name[2] == '*' {
			// Protected -seem-to- have the format "\0*\0propname"
			prop.Visibility = Protected
		} else {
			// Private -seem-to- have the format "\0classname\0propname"
			prop.Visibility = Private
		}
	}
	pf, err = p.GetParseFunc(p.EatTypeFlag())
	if err != nil {
		p.Err = err
		goto end
	}
	prop.Value = pf(p)
	if p.Err != nil {
		p.Err = fmt.Errorf("error parsing array value; %w", p.Err)
		goto end
	}
end:
}
