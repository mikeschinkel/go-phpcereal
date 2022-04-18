package phpcereal

type SQLSerializedGetter interface {
	SQLSerialized() string
}

// MaybeGetSQLSerialized returns the serialized string for
// a CerealValue (cv) by first checking to see if the cv
// implements SQLSerializedGetter; if yes it calls
// .SQLSerialized() otherwise it calls .Serialized().
func MaybeGetSQLSerialized(cv CerealValue) (s string) {
	getter, ok := cv.(SQLSerializedGetter)
	if !ok {
		s = cv.Serialized()
		goto end
	}
	s = getter.SQLSerialized()
end:
	return s
}
