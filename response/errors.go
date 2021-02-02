package response

import (
	"errors"
	"fmt"
)

// PathError отсутствует путь в JSON
type PathError struct {
	Path string
}

func (e *PathError) Error() string {
	return fmt.Sprintf("no path %s in a response", e.Path)
}

// errEmptyBody пустое тело ответа метода API ЭЛДИС
var errEmptyBody = errors.New("empty body")

// unavailableForSection не доступна реализация для указанной секции
var errUnavailableForSection = errors.New("unavailable for this section")
