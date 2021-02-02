package archive

import (
	"fmt"
)

// ErrorUnknownArchive неизвестный тип архива показаний
type ErrorUnknownArchive struct {
	Value string
}

func (e *ErrorUnknownArchive) Error() string {
	return fmt.Sprintf(`unknown archive "%s"`, e.Value)
}
