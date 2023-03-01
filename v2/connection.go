package eldis

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"

	"github.com/vitpelekhaty/go-eldis/v2/responses"
)

var _ Connection = (*connection)(nil)

type connection struct {
	client *http.Client
	token  string
	rawURL string

	credentials *Credentials
}

// Open открывает соединение с API АИИС ЭЛДИС
func (conn *connection) Open(ctx context.Context, rawURL string, credentials Credentials) error {
	if conn.connected() {
		return ErrConnected
	}

	if _, err := url.Parse(rawURL); err != nil {
		return err
	}

	conn.rawURL = rawURL
	conn.credentials = &credentials

	token, err := conn.reconnect(ctx, conn.credentials)

	if err != nil {
		return err
	}

	conn.token = token

	return nil
}

const pathLogout = "/api/v2/users/logout"

// Close закрывает соединение с API АИИС ЭЛДИС
func (conn *connection) Close(ctx context.Context) error {
	if !conn.connected() {
		return nil
	}

	rawURL, err := join(conn.rawURL, nil, pathLogout)

	if err != nil {
		return err
	}

	headers := map[string]string{
		"Cookie": fmt.Sprintf("access_token=%s", conn.token),
		"key":    conn.credentials.AccessToken,
	}

	response, err := conn.call(ctx, http.MethodGet, rawURL, headers, nil)

	if err != nil {
		return err
	}

	_, err = body(response)

	if err != nil {
		return err
	}

	conn.token = ""

	return nil
}

const pathListForDevelopment = "/api/v2/tv/listForDevelopment"

// ListForDevelopment возвращает список доступных пользователю точек учета
func (conn *connection) ListForDevelopment(ctx context.Context) ([]byte, error) {
	if !conn.connected() {
		return nil, ErrNotConnected
	}

	rawURL, err := join(conn.rawURL, nil, pathListForDevelopment)

	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Cookie": fmt.Sprintf("access_token=%s", conn.token),
		"key":    conn.credentials.AccessToken,
	}

	response, err := conn.call(ctx, http.MethodGet, rawURL, headers, nil)

	if err != nil {
		if rse, ok := err.(*RemoteServerError); ok {
			if rse.StatusCode == http.StatusUnauthorized {
				token, err := conn.reconnect(ctx, conn.credentials)

				if err != nil {
					return nil, err
				}

				conn.token = token
				headers["Cookie"] = fmt.Sprintf("access_token=%s", conn.token)

				response, err = conn.call(ctx, http.MethodGet, rawURL, headers, nil)

				if err != nil {
					return nil, err
				}
			}
		} else {
			return nil, err
		}
	}

	b, err := body(response)

	if err != nil {
		return nil, err
	}

	return b, nil
}

const pathDataNormalized = "/api/v2/data/normalized"
const layoutRequestTime = "02.01.2006 15:04:05"

// NormalizedReadings возвращает нормализованные показания точки учета, удовлетворяющие условиям
func (conn *connection) NormalizedReadings(ctx context.Context, pointID string, archive Archive, from, to time.Time,
	dateType DateType) ([]byte, error) {
	if !conn.connected() {
		return nil, ErrNotConnected
	}

	query := map[string]string{
		"id":           pointID,
		"typeDataCode": codeArchive[archive],
		"startDate":    from.Format(layoutRequestTime),
		"endDate":      to.Format(layoutRequestTime),
		"dateType":     string(dateType),
	}

	rawURL, err := join(conn.rawURL, query, pathDataNormalized)

	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Cookie": fmt.Sprintf("access_token=%s", conn.token),
		"key":    conn.credentials.AccessToken,
	}

	response, err := conn.call(ctx, http.MethodGet, rawURL, headers, nil)

	if err != nil {
		if rse, ok := err.(*RemoteServerError); ok {
			if rse.StatusCode == http.StatusUnauthorized {
				token, err := conn.reconnect(ctx, conn.credentials)

				if err != nil {
					return nil, err
				}

				conn.token = token
				headers["Cookie"] = fmt.Sprintf("access_token=%s", conn.token)

				response, err = conn.call(ctx, http.MethodGet, rawURL, headers, nil)

				if err != nil {
					return nil, err
				}
			}
		} else {
			return nil, err
		}
	}

	b, err := body(response)

	if err != nil {
		return nil, err
	}

	return b, nil
}

const pathDataRaw = "/api/v2/data/rawData"

// RawReadings возвращает "сырые" показания точки учета, удовлетворяющие условиям
func (conn *connection) RawReadings(ctx context.Context, pointID string, archive Archive, from, to time.Time) ([]byte, error) {
	if !conn.connected() {
		return nil, ErrNotConnected
	}

	query := map[string]string{
		"id":           pointID,
		"typeDataCode": codeArchive[archive],
		"startDate":    from.Format(layoutRequestTime),
		"endDate":      to.Format(layoutRequestTime),
	}

	rawURL, err := join(conn.rawURL, query, pathDataRaw)

	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Cookie": fmt.Sprintf("access_token=%s", conn.token),
		"key":    conn.credentials.AccessToken,
	}

	response, err := conn.call(ctx, http.MethodGet, rawURL, headers, nil)

	if err != nil {
		if rse, ok := err.(*RemoteServerError); ok {
			if rse.StatusCode == http.StatusUnauthorized {
				token, err := conn.reconnect(ctx, conn.credentials)

				if err != nil {
					return nil, err
				}

				conn.token = token
				headers["Cookie"] = fmt.Sprintf("access_token=%s", conn.token)

				response, err = conn.call(ctx, http.MethodGet, rawURL, headers, nil)

				if err != nil {
					return nil, err
				}
			}
		} else {
			return nil, err
		}
	}

	b, err := body(response)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (conn *connection) connected() bool {
	return strings.TrimSpace(conn.token) != ""
}

const pathLogin = "/api/v2/users/login"

func (conn *connection) reconnect(ctx context.Context, credentials *Credentials) (token string, err error) {
	rawURL, err := join(conn.rawURL, nil, pathLogin)

	if err != nil {
		return
	}

	form := url.Values{}

	form.Set("login", credentials.Username)
	form.Set("password", credentials.Password)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
		"key":          credentials.AccessToken,
	}

	response, err := conn.call(ctx, http.MethodPost, rawURL, headers, strings.NewReader(form.Encode()))

	if err != nil {
		return
	}

	_, err = body(response)

	if err != nil {
		return
	}

	c, ok := cookie(response, "access_token")

	if !ok || c.Value == "" {
		return "", &MethodCallError{Method: http.MethodPost, RawURL: rawURL, Err: ErrNotConnected}
	}

	token = c.Value

	return
}

func (conn *connection) call(ctx context.Context, method, rawURL string, headers map[string]string, body io.Reader) (*http.Response, error) {
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

	if response.StatusCode != http.StatusOK {
		return nil, &RemoteServerError{StatusCode: response.StatusCode}
	}

	return response, nil
}

func body(response *http.Response) ([]byte, error) {
	defer func() {
		_ = response.Body.Close()
	}()

	b, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &MethodCallError{Method: response.Request.Method, RawURL: response.Request.URL.String(), Err: err}
	}

	resp, err := responses.Parse(bytes.NewBuffer(b))

	if err != nil {
		return nil, &MethodCallError{Method: response.Request.Method, RawURL: response.Request.URL.String(), Err: err}
	}

	var ve error

	for _, message := range resp.Messages() {
		if message.Level == responses.MessageLevelError && message.StatusCode != http.StatusOK {
			ve = multierror.Append(ve, fmt.Errorf("%d %s", message.Code, message.Message))
		}
	}

	if ve != nil {
		return nil, &InternalServerError{
			&MethodCallError{
				Method: response.Request.Method,
				RawURL: response.Request.URL.String(),
				Err:    ve,
			},
		}
	}

	return b, nil
}

func cookie(response *http.Response, name string) (*http.Cookie, bool) {
	for _, cookie := range response.Cookies() {
		if strings.EqualFold(cookie.Name, name) {
			return cookie, true
		}
	}

	return nil, false
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
