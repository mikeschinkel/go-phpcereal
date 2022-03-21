<?php

// see https://formats.kaitai.io/php_serialized_value/index.html

use JetBrains\PhpStorm\ArrayShape;

class Foo
{
    public string $foo;
    private int   $bar;

    function __construct(string $foo, int $bar)
    {
        $this->foo = $foo;
        $this->bar = $bar;
    }
    function __ToString():string {
        return sprintf("new Foo('%s',%d)",$this->foo,$this->bar);
    }
}

class Bar implements Serializable
{
    public Foo $foo;

    function __construct(Foo $foo)
    {
        $this->foo = $foo;
    }
    function __ToString():string {
        return sprintf("new Bar('%s')",str_replace( "'", "\'", (string)$this->foo));
    }

    public function serialize(): ?string
    {
        return serialize($this->foo);
    }

    public function unserialize(string $data)
    {
        return unserialize($data);
    }

    #[ArrayShape(["foo" => "string"])]
    public function __serialize(): array
    {
        return ["foo" => serialize($this->foo)];
    }

    public function __unserialize(array $data): void
    {
        $this->foo = unserialize($data[ "foo" ]);
    }
}
function printArray() {

}

$foo = new Foo("abc", 13);

$a             = array(
    "Null"      /*  - N -  */ => null,
    "String"    /*  - s -  */ => "hello",
    "Php6Str"   /*  - S -  */ => "world",
    "Int"       /*  - i -  */ => 123,
    "Bool"      /*  - b -  */ => true,
    "False"     /*  - b -  */ => false,
    "Float"     /*  - d -  */ => 12.34,
    "Object"    /*  - O -  */ => $foo,
    "Array"     /*  - a -  */ => [1, 2, 3],
    "ObjectRef" /*  - R -  */ => &$foo,
    "CSObject"  /*  - C -  */ => new Bar($foo),

//    "o - Php3Object"             => "",
);
$a[ "VarRef" ] =& $a[ "String" ];


function to_table($value,$type,$indent=""):string {
    $string = "";
    switch (GetValueType($value)) {
        case "array":
            $array = [];
            foreach ($value as $t => $v) {
                $array[] = to_table($v,$t,$indent."\t");
            }
            $string .= sprintf("%s[\n%s%s]\n",$indent,implode(",",$array),$indent);
        default:
            $string .= sprintf("%s%-12s: %-20s\n", $indent,$type, serialize($value));
    }
    return $string;
}

echo to_table($a,"Array");

//echo serialize($a);

//$a2  = unserialize(serialize($a));
//printf("%s\n", $a2[ "VarRef" ]);
//$a2[ "String" ] = "goodbye";
//printf("%s\n", $a2[ "VarRef" ]);
