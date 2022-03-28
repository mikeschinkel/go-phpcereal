package phpcereal

type CerealValue interface {
	GetValue() interface{}
	GetType() PHPType
	GetTypeFlag() TypeFlag
	String() string
	Serialized() string
	SerializedLen() int
	Parse(*Parser) CerealValue
}
