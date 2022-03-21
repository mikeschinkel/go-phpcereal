package phpcereal

var _ ValueAccessor = (*NullValue)(nil)

type NullValue struct{}

func (v NullValue) GetValue() interface{} {
	return nil
}

func (v NullValue) GetType() PHPType {
	return "NULL"
}

func (v NullValue) GetTypeFlag() TypeFlag {
	return NULLTypeFlag
}

func (v NullValue) String() string {
	return "NULL"
}

func (v NullValue) Serialized() string {
	return "N;"
}

func (v NullValue) SerializedLen() int {
	return 2
}

func (v NullValue) Parse(_ *Parser) (_ ValueAccessor) {
	return v
}
