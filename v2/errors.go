package eldis

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNotConnected = errors.New("не установлено соединение")
	ErrConnected    = errors.New("соединение уже установлено")
)

// MethodCallError ошибка вызова метода API АИИС ЭЛДИС
type MethodCallError struct {
	// RawURL URL метода API
	RawURL string

	// Method метод
	Method string

	// Err ошибка, которая возникла во время вызова метода
	Err error
}

func (err *MethodCallError) Error() string {
	return fmt.Sprintf("%s %s: %s", err.Method, err.RawURL, err.Err)
}

// InternalServerError внутренняя ошибка сервера API АИИС ЭЛДИС
type InternalServerError struct {
	*MethodCallError
}

func (err *InternalServerError) Error() string {
	return fmt.Sprintf("%s %s: %s", err.Method, err.RawURL, err.Err)
}

// RemoteServerError ошибка сервера API АИИС ЭЛДИС
type RemoteServerError struct {
	// StatusCode код ошибки удаленного сервера
	StatusCode int
}

func (err *RemoteServerError) Error() string {
	return fmt.Sprintf("%d %s", err.StatusCode, http.StatusText(err.StatusCode))
}
