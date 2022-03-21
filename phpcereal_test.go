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
	s string             // Serialized String
	v string             // Go Value
	t string             // PHP Type
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
		n: "String:world",
		f: phpcereal.PHP6StringTypeFlag,
		s: `S:5:"world";`,
		v: `"world"`,
		t: "6string",
	},
	{
		n: "Array of Integers:[1,2,3]",
		f: phpcereal.ArrayTypeFlag,
		s: "a:3:{i:0;i:1;i:1;i:2;i:2;i:3;}",
		v: "[0=>1,1=>2,2=>3]",
		t: "array",
	},
	{
		n: "Integer:123",
		f: phpcereal.IntTypeFlag,
		s: "i:123;",
		v: "123",
		t: "int",
	},
	{
		n: "String:hello",
		f: phpcereal.StringTypeFlag,
		s: `s:5:"hello";`,
		v: `"hello"`,
		t: "string",
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
}

func TestParsing(t *testing.T) {
	for _, test := range testdata {
		t.Run(test.n, func(t *testing.T) {
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
			if assert.Equal(t, phpcereal.PHPType(test.t), root.GetType()) {
				assert.Equal(t, test.v, root.String())
			} else {
				return
			}
		})
	}
}
