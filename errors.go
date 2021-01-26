package eldis

import (
	"fmt"
	"strings"

	"github.com/vitpelekhaty/go-eldis/response"
)

// MethodCallError ошибка вызова метода API
type MethodCallError struct {
	// RawURL адрес метода API
	RawURL string
	// Method HTTP-метод
	Method string
	// Err ошибка вызова метода
	Err error
}

func (e *MethodCallError) Error() string {
	var b strings.Builder

	if strings.TrimSpace(e.Method) != "" {
		b.WriteString(e.Method)
	}

	if strings.TrimSpace(e.RawURL) != "" {
		if b.Len() > 0 {
			b.WriteRune(' ')
		}

		b.WriteString(e.RawURL)
	}

	if e.Err != nil {
		if b.Len() > 0 {
			b.WriteString(": ")
		}

		b.WriteString(e.Err.Error())
	}

	return b.String()
}

// InternalError внутренняя ошибка АИСКУТЭ ЭЛДИС при обработке запроса
type InternalError struct {
	// RawURL адрес метода API
	RawURL string
	// Method HTTP метод
	Method string
	// Status состояние ошибки
	Status *response.Message
	// Err внутренняя ошибка API
	Err error
}

func (e *InternalError) Error() string {
	var b strings.Builder

	if strings.TrimSpace(e.Method) != "" {
		b.WriteString(e.Method)
	}

	if strings.TrimSpace(e.RawURL) != "" {
		if b.Len() > 0 {
			b.WriteRune(' ')
		}

		b.WriteString(e.RawURL)
	}

	if e.Err != nil {
		if b.Len() > 0 {
			b.WriteString(": ")
		}

		b.WriteString(e.Err.Error())
	}

	return b.String()
}

func newMethodCallError(rawURL, method string, err error) *MethodCallError {
	return &MethodCallError{
		RawURL: rawURL,
		Method: method,
		Err:    err,
	}
}

func newInternalError(rawURL, method string, status *response.Message) *InternalError {
	return &InternalError{
		RawURL: rawURL,
		Method: method,
		Status: status,
		Err:    fmt.Errorf("%d %s", status.Code, status.Message),
	}
}
