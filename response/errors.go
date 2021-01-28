package response

import (
	"fmt"
)

// PathError отсутствует путь в JSON
type PathError struct {
	Path string
}

func (e *PathError) Error() string {
	return fmt.Sprintf("no path %s in a response", e.Path)
}
