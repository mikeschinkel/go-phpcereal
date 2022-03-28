package phpcereal

import (
	"fmt"
)

type ArrayElement struct {
	Key   CerealValue
	Value CerealValue
}

func (e ArrayElement) String() (s string) {
	return fmt.Sprintf("%s=>%s,", e.Key.String(), e.Value.String())
}
