package phpcereal

type TypeFlag byte

const (
	CustomObjTypeFlag  TypeFlag = 'C'
	NULLTypeFlag       TypeFlag = 'N'
	ObjectTypeFlag     TypeFlag = 'O'
	VarRefTypeFlag     TypeFlag = 'R'
	PHP6StringTypeFlag TypeFlag = 'S'
	ArrayTypeFlag      TypeFlag = 'a'
	BoolTypeFlag       TypeFlag = 'b'
	FloatTypeFlag      TypeFlag = 'd'
	IntTypeFlag        TypeFlag = 'i'
	PHP3ObjTypeFlag    TypeFlag = 'o'
	ObjRefTypeFlag     TypeFlag = 'r'
	StringTypeFlag     TypeFlag = 's'
)

type PHPType string
type TypeFlagSetter interface {
	SetTypeFlag(TypeFlag)
}

type StringReplacer interface {
	ReplaceString(from, to string, times int)
}

type Chars interface {
	string | []byte
}
