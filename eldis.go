package eldis

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/vitpelekhaty/go-eldis/archive"
	"github.com/vitpelekhaty/go-eldis/response"
)

// ConnectionOption параметр соединения с API ЭЛДИС
type ConnectionOption func(c *Connection)

// WithAuthOption тип параметра соединения для авторизации клиента к API ЭЛДИС
type WithAuthOption ConnectionOption

// WithAuth обязательный параметр соединения для авторизации клиента к API ЭЛДИС
func WithAuth(username, password, key string) WithAuthOption {
	return func(c *Connection) {
		c.auth = &auth{
			username: username,
			password: password,
			key:      key,
		}
	}
}

type UnhandledErrorFunc func(err error)

// Connection соединение с API ЭЛДИС
type Connection struct {
	// OnUnhandledError событие появления ошибки внутри соединения, которая не передается вызывающей стороне, но может
	// записана в лог вызывающей стороной
	OnUnhandledError UnhandledErrorFunc

	client *http.Client

	rawURL string

	token string
	auth  *auth
}

// NewConnection возвращает новое соединение с API ЭЛДИС
func NewConnection(client *http.Client) (*Connection, error) {
	if client == nil {
		return nil, errors.New("no HTTP client")
	}

	return &Connection{
		client: client,
	}, nil
}

// Open открывает соединение с API ЭЛДИС
func (c *Connection) Open(rawURL string, withAuth WithAuthOption) error {
	if c.Connected() {
		return errors.New("connection has already been established")
	}

	_, err := url.Parse(rawURL)

	if err != nil {
		return err
	}

	c.rawURL = rawURL

	if withAuth == nil {
		return errors.New("no auth info")
	}

	withAuth(c)

	return c.login()
}

// Close закрывает соединение с API ЭЛДИС
func (c *Connection) Close() error {
	if !c.Connected() {
		return nil
	}

	err := c.logout()

	if err != nil {
		return err
	}

	c.token = ""
	c.auth = nil

	return nil
}

// Connected возвращает статус соединения с API ЭЛДИС
func (c *Connection) Connected() bool {
	return strings.TrimSpace(c.token) != ""
}

// ListForDevelopment вызывает метод /api/v2/tv/listForDevelopment API для получения списка доступных точек учета
func (c *Connection) ListForDevelopment() ([]byte, error) {
	return nil, nil
}

// UOMList вызывает метод /api/v2/uom/list API для получения списка единиц измерения
func (c *Connection) UOMList() ([]byte, error) {
	return nil, nil
}

func (c *Connection) DataNormalized(regPointID string, archive archive.DataArchive, from, to RequestTime,
	dateType DateType) ([]byte, error) {
	return nil, nil
}

func (c *Connection) RawData(regPointID string, archive archive.DataArchive, from,
	to RequestTime) ([]byte, error) {
	return nil, nil
}

func (c *Connection) login() error {
	methodRawURL, err := join(c.rawURL, methodLogin)

	if err != nil {
		return err
	}

	form := url.Values{}

	form.Add("login", c.auth.username)
	form.Add("password", c.auth.password)

	req, err := http.NewRequest("POST", methodRawURL, strings.NewReader(form.Encode()))

	if err != nil {
		return newMethodCallError(methodLogin, "POST", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("key", c.auth.key)

	resp, err := c.client.Do(req)

	if err != nil {
		return newMethodCallError(methodLogin, "POST", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			c.doUnhandledError(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return newMethodCallError(methodLogin, "POST", fmt.Errorf("%d %s", resp.StatusCode, resp.Status))
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return newMethodCallError(methodLogin, "POST", err)
	}

	message, err := c.responseStatus(body)

	if err != nil {
		return newMethodCallError(methodLogin, "POST", err)
	}

	if message.StatusCode != http.StatusOK {
		return newInternalError(methodLogin, "POST", message)
	}

	token, err := c.cookie(resp.Cookies(), "access_token")

	if err != nil {
		return newMethodCallError(methodLogin, "POST", err)
	}

	c.token = token.Value

	return nil
}

func (c *Connection) logout() error {
	return nil
}

func (c *Connection) call(u url.URL, method string) ([]byte, error) {
	return nil, nil
}

func (c *Connection) checkConnection() error {
	return nil
}

func (c *Connection) cookie(cookies []*http.Cookie, name string) (*http.Cookie, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("invalid cookie name")
	}

	if len(cookies) == 0 {
		return nil, fmt.Errorf("no %s", name)
	}

	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie, nil
		}
	}

	return nil, fmt.Errorf("no %s", name)
}

func (c *Connection) responseStatus(body []byte) (*response.Message, error) {
	resp, err := response.Parse(body)

	if err != nil {
		return nil, err
	}

	messages := resp.Messages()

	if len(messages) == 1 {
		return &messages[0], nil
	}

	return nil, errors.New("invalid response format")
}

func (c *Connection) doUnhandledError(err error) {
	if c.OnUnhandledError != nil {
		c.OnUnhandledError(err)
	}
}
