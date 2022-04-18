package phpcereal_test

import (
	"testing"

	"github.com/mikeschinkel/go-phpcereal"
	"github.com/stretchr/testify/assert"
)

//const (
//	sAll      = `a:12:{s:4:"Null";N;s:6:"String";s:5:"hello";s:7:"Php6Str";s:5:"world";s:3:"Int";i:123;s:4:"Bool";b:1;s:5:"False";b:0;s:5:"Float";d:12.34;s:6:"Object";O:3:"Foo":2:{s:3:"foo";s:3:"abc";s:8:"\0Foo\0bar";i:13;}s:5:"Array";a:3:{i:0;i:1;i:1;i:2;i:2;i:3;}s:9:"ObjectRef";R:9;s:8:"CSObject";O:3:"Bar":1:{s:3:"foo";s:54:"O:3:"Foo":2:{s:3:"foo";s:3:"abc";s:8:"\0Foo\0bar";i:13;}";}s:6:"VarRef";R:3;}`
//	sCSObject = `O:3:"Bar":1:{s:3:"foo";s:54:"O:3:"Foo":2:{s:3:"foo";s:3:"abc";s:8:"\0Foo\0bar";i:13;}";}`
//	sVarRef   = `s:5:"hello";`
//)

type TestData struct {
	n string             // Test Name
	f phpcereal.TypeFlag // Type Flag
	e bool               // Escaped
	s string             // Serialized String
	v string             // Go Value
	t string             // PHP Type
	r []string           // Find/Replace strings
}

func (test TestData) IsCereal() bool {
	if test.e {
		return phpcereal.IsEscapedCereal(test.s)
	}
	return phpcereal.IsCereal(test.s)
}

var testdata = []TestData{
	{
		n: "Object: Foo",
		f: phpcereal.ObjectTypeFlag,
		s: `O:3:"Foo":3:{s:3:"foo";s:3:"abc";s:8:"\0Foo\0bar";i:13;s:6:"\0*\0baz";b:1;}`,
		v: `Foo{foo:"abc",~bar:13,^baz:true}`, // Legend: ~private, ^protected
		t: "Foo",
	},
	{
		n: "Float:12.34",
		f: phpcereal.FloatTypeFlag,
		s: "d:12.34;",
		v: "12.34",
		t: "float",
	},
	{
		n: "Array of URLs",
		f: phpcereal.ArrayTypeFlag,
		s: `a:3:{i:0;s:40:"https://en.wiktionary.org/wiki/enquoting";i:1;s:41:"https://en.wiktionary.org/wiki/whiff#Verb";i:2;s:42:"https://en.wiktionary.org/wiki/tea#Spanish";}`,
		v: `[0=>"https://en.wikipedia.org/wiki/enquoting",1=>"https://en.wikipedia.org/wiki/whiff#Verb",2=>"https://en.wikipedia.org/wiki/tea#Spanish",]`,
		t: "array",
		r: []string{"wiktionary.org", "wikipedia.org"},
	},
	{
		n: "Array of Integers:[1,2,3]",
		f: phpcereal.ArrayTypeFlag,
		s: "a:3:{i:0;i:1;i:1;i:2;i:2;i:3;}",
		v: "[0=>1,1=>2,2=>3,]",
		t: "array",
	},
	{
		n: "String:hello",
		f: phpcereal.StringTypeFlag,
		s: `s:5:"hello";`,
		v: `"herro"`,
		t: "string",
		r: []string{"ll", "rr"},
	},
	{
		n: "String:world",
		f: phpcereal.PHP6StringTypeFlag,
		s: `S:5:"world";`,
		v: `"world"`,
		t: "6string",
	},
	{
		n: "NULL",
		f: phpcereal.NULLTypeFlag,
		s: `N;`,
		v: "NULL",
		t: "NULL",
	},
	{
		n: "Bool:true",
		f: phpcereal.BoolTypeFlag,
		s: "b:1;",
		v: "true",
		t: "bool",
	},
	{
		n: "Bool:false",
		f: phpcereal.BoolTypeFlag,
		s: "b:0;",
		v: "false",
		t: "bool",
	},
	{
		n: "Integer:123",
		f: phpcereal.IntTypeFlag,
		s: "i:123;",
		v: "123",
		t: "int",
	},
}

func TestParsing(t *testing.T) {
	for _, test := range testdata {
		t.Run(test.n, func(t *testing.T) {
			if !test.IsCereal() {
				t.Errorf("failed to validate: %s", test.s)
			}
			sp := phpcereal.NewParser(test.s)
			root, err := sp.Parse()
			if err != nil {
				t.Error(err.Error())
				return
			}
			s := root.Serialized()
			assert.Equal(t, test.s, s)
			if assert.NotEmpty(t, s) {
				assert.Equal(t, test.f, phpcereal.TypeFlag(s[0]))
			} else {
				return
			}
			if len(test.r) == 2 {
				if sr, ok := root.(phpcereal.StringReplacer); ok {
					sr.ReplaceString(test.r[0], test.r[1], -1)
				}
			}
			if assert.Equal(t, phpcereal.PHPType(test.t), root.GetType()) {
				assert.Equal(t, test.v, root.String())
			} else {
				t.Errorf("mismatch in types: %s <> %s", phpcereal.PHPType(test.t), root.GetType())
			}
		})
	}
}
