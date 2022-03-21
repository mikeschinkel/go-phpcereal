package phpcereal

import (
	"fmt"
)

type ArrayElement struct {
	Key   ValueAccessor
	Value ValueAccessor
}

func (e ArrayElement) String() (s string) {
	return fmt.Sprintf("%s=>%s,", e.Key.String(), e.Value.String())
}
