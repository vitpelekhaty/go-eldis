package date

import (
	"fmt"
)

type ErrorUnknownType struct {
	Value string
}

func (e *ErrorUnknownType) Error() string {
	return fmt.Sprintf(`unknown type of date "%s"`, e.Value)
}
