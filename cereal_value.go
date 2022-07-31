package phpcereal

type CerealValue interface {
	GetValue() interface{}
	GetType() PHPType
	GetTypeFlag() TypeFlag
	GetEscaped() bool
	String() string
	Serialized() string
	SerializedLen() int
	Parse(*Parser) CerealValue
	GetOpts() CerealOpts
	SetOpts(CerealOpts)
}

func ReplaceString(cv CerealValue, find, replace string) (err error) {
	sr, ok := cv.(StringReplacer)
	if ok {
		sr.ReplaceString(find, replace, -1)
	}
	return err
}
