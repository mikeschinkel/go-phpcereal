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
		n: "Array: Complex Escaped",
		f: phpcereal.ArrayTypeFlag,
		t: "array",
		e: true,
		s: `a:24:{` +
			`i:0;s:34:\"advanced-custom-fields-pro/acf.php\";` +
			`i:1;s:55:\"advanced-post-types-order/advanced-post-types-order.php\";` +
			`i:2;s:39:\"aryo-activity-log/aryo-activity-log.php\";` +
			`i:3;s:28:\"category-posts/cat-posts.php\";` +
			`i:4;s:33:\"classic-editor/classic-editor.php\";` +
			`i:5;s:49:\"constant-contact-forms/constant-contact-forms.php\";` +
			`i:6;s:35:\"disable-xml-rpc/disable-xml-rpc.php\";` +
			`i:7;s:61:\"divi_extended_column_layouts/divi_extended_column_layouts.php\";` +
			`i:8;s:32:\"duplicate-page/duplicatepage.php\";` +
			`i:9;s:49:\"elegant-themes-updater/elegant-themes-updater.php\";` +
			`i:10;s:45:\"enable-media-replace/enable-media-replace.php\";` +
			`i:11;s:53:\"enhanced-media-library-pro/enhanced-media-library.php\";` +
			`i:12;s:21:\"include-me/plugin.php\";` +
			`i:13;s:30:\"interactive-world-maps/map.php\";` +
			`i:14;s:63:\"limit-login-attempts-reloaded/limit-login-attempts-reloaded.php\";` +
			`i:15;s:25:\"menu-image/menu-image.php\";i:17;s:19:\"monarch/monarch.php\";` +
			`i:16;s:33:\"posts-to-posts/posts-to-posts.php\";i:19;s:27:\"redirection/redirection.php\";` +
			`i:28;s:47:\"regenerate-thumbnails/regenerate-thumbnails.php\";` +
			`i:20;s:23:\"revslider/revslider.php\";` +
			`i:21;s:45:\"taxonomy-terms-order/taxonomy-terms-order.php\";` +
			`i:22;s:24:\"wordpress-seo/wp-seo.php\";` +
			`i:23;s:29:\"wp-shopify-pro/wp-shopify.php\";` +
			`}`,
		v: `[` +
			`0=>"advanced-custom-fields-pro/acf.php",` +
			`1=>"advanced-post-types-order/advanced-post-types-order.php",` +
			`2=>"aryo-activity-log/aryo-activity-log.php",` +
			`3=>"category-posts/cat-posts.php",` +
			`4=>"classic-editor/classic-editor.php",` +
			`5=>"constant-contact-forms/constant-contact-forms.php",` +
			`6=>"disable-xml-rpc/disable-xml-rpc.php",` +
			`7=>"divi_extended_column_layouts/divi_extended_column_layouts.php",` +
			`8=>"duplicate-page/duplicatepage.php",` +
			`9=>"elegant-themes-updater/elegant-themes-updater.php",` +
			`10=>"enable-media-replace/enable-media-replace.php",` +
			`11=>"enhanced-media-library-pro/enhanced-media-library.php",` +
			`12=>"include-me/plugin.php",` +
			`13=>"interactive-world-maps/map.php",` +
			`14=>"limit-login-attempts-reloaded/limit-login-attempts-reloaded.php",` +
			`15=>"menu-image/menu-image.php",` +
			`17=>"monarch/monarch.php",` +
			`16=>"posts-to-posts/posts-to-posts.php",` +
			`19=>"redirection/redirection.php",` +
			`28=>"regenerate-thumbnails/regenerate-thumbnails.php",` +
			`20=>"revslider/revslider.php",` +
			`21=>"taxonomy-terms-order/taxonomy-terms-order.php",` +
			`22=>"wordpress-seo/wp-seo.php",23=>"wp-shopify-pro/wp-shopify.php",` +
			`]`,
	},
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
	var s string

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
			if test.e {
				s = phpcereal.MaybeGetSQLSerialized(root)
			} else {
				s = root.Serialized()
			}
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
