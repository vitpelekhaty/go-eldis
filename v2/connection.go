package eldis

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
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

var _ Connection = (*connection)(nil)

type connection struct {
	client *http.Client
	token  string
	rawURL string
	user   *UserOptions
}

// Open открывает соединение с API АИИС ЭЛДИС
func (conn *connection) Open(ctx context.Context, rawURL string, user UserOptions) error {
	if conn.connected() {
		return ErrConnected
	}

	if _, err := url.Parse(rawURL); err != nil {
		return err
	}

	conn.rawURL = rawURL

	token, err := conn.reconnect(ctx, user.Username, user.Password, user.AccessToken)

	if err != nil {
		return err
	}

	conn.token = token
	conn.user = &user

	return nil
}

// Close закрывает соединение с API АИИС ЭЛДИС
func (conn *connection) Close(ctx context.Context) error {
	if !conn.connected() {
		return nil
	}

	return nil
}

// ListForDevelopment возвращает список доступных пользователю точек учета
func (conn *connection) ListForDevelopment(ctx context.Context) ([]byte, error) {
	if !conn.connected() {
		return nil, ErrNotConnected
	}

	return nil, nil
}

// NormalizedReadings возвращает нормализованные показания точки учета, удовлетворяющие условиям
func (conn *connection) NormalizedReadings(ctx context.Context, pointID string, archive Archive, from, to time.Time,
	dateType DateType) ([]byte, error) {
	if !conn.connected() {
		return nil, ErrNotConnected
	}

	return nil, nil
}

// RawReadings возвращает "сырые" показания точки учета, удовлетворяющие условиям
func (conn *connection) RawReadings(ctx context.Context, pointID string, archive Archive, from, to time.Time) ([]byte, error) {
	if !conn.connected() {
		return nil, ErrNotConnected
	}

	return nil, nil
}

func (conn *connection) connected() bool {
	return strings.TrimSpace(conn.token) != ""
}

const pathLogin = "/api/v2/users/login"

func (conn *connection) reconnect(ctx context.Context, username, password, accessToken string) (token string, err error) {
	rawURL, err := join(conn.rawURL, nil, pathLogin)

	if err != nil {
		return
	}

	form := url.Values{}

	form.Set("login", username)
	form.Set("password", password)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
		"key":          accessToken,
	}

	_, err = conn.call(ctx, http.MethodPost, rawURL, headers, strings.NewReader(form.Encode()))

	if err != nil {
		return
	}

	return "", nil
}

func (conn *connection) call(ctx context.Context, method, rawURL string, headers map[string]string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, method, rawURL, body)

	if err != nil {
		return nil, &MethodCallError{Method: method, RawURL: rawURL, Err: err}
	}

	for header, value := range headers {

		request.Header.Set(header, value)
	}

	response, err := conn.client.Do(request)

	if err != nil {
		return nil, &MethodCallError{Method: method, RawURL: rawURL, Err: err}
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return nil, &MethodCallError{
			Method: method,
			RawURL: rawURL,
			Err:    fmt.Errorf("%d %s", response.StatusCode, http.StatusText(response.StatusCode)),
		}
	}

	return io.ReadAll(response.Body)
}

func join(addr string, query map[string]string, paths ...string) (string, error) {
	u, err := url.Parse(addr)

	if err != nil {
		return "", err
	}

	if len(paths) > 0 {
		elem := []string{u.Path}
		elem = append(elem, paths...)

		p := path.Join(elem...)
		u.Path = path.Clean(p)
	}

	if len(query) > 0 {
		q := u.Query()

		for key, value := range query {
			q.Set(key, value)
		}

		u.RawQuery = q.Encode()
	}

	return u.String(), nil
}
