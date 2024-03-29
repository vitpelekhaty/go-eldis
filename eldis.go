package eldis

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/vitpelekhaty/go-eldis/archive"
	"github.com/vitpelekhaty/go-eldis/date"
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

// UnhandledErrorFunc обработчик появления ошибки внутри соединения с API ЭЛДИС
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
	return c.OpenWithContext(context.Background(), rawURL, withAuth)
}

// OpenWithContext открывает соединение с API ЭЛДИС
func (c *Connection) OpenWithContext(ctx context.Context, rawURL string, withAuth WithAuthOption) error {
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

	return c.login(ctx)
}

// Close закрывает соединение с API ЭЛДИС
func (c *Connection) Close() error {
	return c.CloseWithContext(context.Background())
}

// CloseWithContext закрывает соединение с API ЭЛДИС
func (c *Connection) CloseWithContext(ctx context.Context) error {
	if !c.Connected() {
		return nil
	}

	err := c.logout(ctx)

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
func (c *Connection) ListForDevelopment(flags ...Flag) ([]byte, error) {
	return c.ListForDevelopmentWithContext(context.Background(), flags...)
}

// ListForDevelopmentWithContext вызывает метод /api/v2/tv/listForDevelopment API для получения списка доступных точек
// учета
func (c *Connection) ListForDevelopmentWithContext(ctx context.Context, flags ...Flag) ([]byte, error) {
	if !c.Connected() {
		return nil, newMethodCallError(methodListForDevelopment, "GET", errors.New("no connection"))
	}

	methodRawURL, err := join(c.rawURL, methodListForDevelopment)

	if err != nil {
		return nil, newMethodCallError(methodListForDevelopment, "GET", err)
	}

	u, err := url.Parse(methodRawURL)

	if err != nil {
		return nil, newMethodCallError(methodListForDevelopment, "GET", err)
	}

	body, err := c.call(ctx, u, "GET", flags...)

	if err != nil {
		return nil, newMethodCallError(methodListForDevelopment, "GET", err)
	}

	message, err := c.responseStatus(body)

	if err != nil {
		return nil, newMethodCallError(methodListForDevelopment, "GET", err)
	}

	if message.StatusCode != http.StatusOK {
		return nil, newInternalError(methodListForDevelopment, "GET", message)
	}

	return body, nil
}

// UOMList вызывает метод /api/v2/uom/list API для получения списка единиц измерения
func (c *Connection) UOMList(flags ...Flag) ([]byte, error) {
	return c.UOMListWithContext(context.Background(), flags...)
}

// UOMListWithContext вызывает метод /api/v2/uom/list API для получения списка единиц измерения
func (c *Connection) UOMListWithContext(ctx context.Context, flags ...Flag) ([]byte, error) {
	if !c.Connected() {
		return nil, newMethodCallError(methodUOMList, "GET", errors.New("no connection"))
	}

	methodRawURL, err := join(c.rawURL, methodUOMList)

	if err != nil {
		return nil, newMethodCallError(methodUOMList, "GET", err)
	}

	u, err := url.Parse(methodRawURL)

	if err != nil {
		return nil, newMethodCallError(methodUOMList, "GET", err)
	}

	body, err := c.call(ctx, u, "GET", flags...)

	if err != nil {
		return nil, newMethodCallError(methodUOMList, "GET", err)
	}

	message, err := c.responseStatus(body)

	if err != nil {
		return nil, newMethodCallError(methodUOMList, "GET", err)
	}

	if message.StatusCode != http.StatusOK {
		return nil, newInternalError(methodUOMList, "GET", message)
	}

	return body, nil
}

// DataNormalized вызывает метод /api/v2/data/normalized для получения нормализованных (после достоверизации) показаний
// на точке учета
func (c *Connection) DataNormalized(regPointID string, archive archive.DataArchive, from, to RequestTime,
	dateType date.Type, flags ...Flag) ([]byte, error) {
	return c.DataNormalizedWithContext(context.Background(), regPointID, archive, from, to, dateType, flags...)
}

// DataNormalizedWithContext вызывает метод /api/v2/data/normalized для получения нормализованных (после достоверизации)
// показаний на точке учета
func (c *Connection) DataNormalizedWithContext(ctx context.Context, regPointID string, archive archive.DataArchive,
	from, to RequestTime, dateType date.Type, flags ...Flag) ([]byte, error) {
	if !c.Connected() {
		return nil, newMethodCallError(methodDataNormalized, "GET", errors.New("no connection"))
	}

	methodRawURL, err := join(c.rawURL, methodDataNormalized)

	if err != nil {
		return nil, newMethodCallError(methodDataNormalized, "GET", err)
	}

	u, err := url.Parse(methodRawURL)

	if err != nil {
		return nil, newMethodCallError(methodDataNormalized, "GET", err)
	}

	query := u.Query()

	query.Add("id", regPointID)
	query.Add("typeDataCode", strconv.Itoa(int(archive)))
	query.Add("startDate", from.String())
	query.Add("endDate", to.String())
	query.Add("dateType", dateType.String())

	u.RawQuery = query.Encode()

	body, err := c.call(ctx, u, "GET", flags...)

	if err != nil {
		return nil, newMethodCallError(methodDataNormalized, "GET", err)
	}

	message, err := c.responseStatus(body)

	if err != nil {
		return nil, newMethodCallError(methodDataNormalized, "GET", err)
	}

	if message.StatusCode != http.StatusOK {
		return nil, newInternalError(methodDataNormalized, "GET", message)
	}

	return body, nil
}

// RawData вызывает метод /api/v2/data/rawData для получения "сырых" показаний на точке учета
func (c *Connection) RawData(regPointID string, archive archive.DataArchive, from,
	to RequestTime, flags ...Flag) ([]byte, error) {
	return c.RawDataWithContext(context.Background(), regPointID, archive, from, to, flags...)
}

// RawDataWithContext вызывает метод /api/v2/data/rawData для получения "сырых" показаний на точке учета
func (c *Connection) RawDataWithContext(ctx context.Context, regPointID string, archive archive.DataArchive, from,
	to RequestTime, flags ...Flag) ([]byte, error) {
	if !c.Connected() {
		return nil, newMethodCallError(methodRawData, "GET", errors.New("no connection"))
	}

	methodRawURL, err := join(c.rawURL, methodRawData)

	if err != nil {
		return nil, newMethodCallError(methodRawData, "GET", err)
	}

	u, err := url.Parse(methodRawURL)

	if err != nil {
		return nil, newMethodCallError(methodRawData, "GET", err)
	}

	query := u.Query()

	query.Add("id", regPointID)
	query.Add("typeDataCode", strconv.Itoa(int(archive)))
	query.Add("startDate", from.String())
	query.Add("endDate", to.String())

	u.RawQuery = query.Encode()

	body, err := c.call(ctx, u, "GET", flags...)

	if err != nil {
		return nil, newMethodCallError(methodRawData, "GET", err)
	}

	message, err := c.responseStatus(body)

	if err != nil {
		return nil, newMethodCallError(methodRawData, "GET", err)
	}

	if message.StatusCode != http.StatusOK {
		return nil, newInternalError(methodRawData, "GET", message)
	}

	return body, nil
}

func (c *Connection) login(ctx context.Context) error {
	methodRawURL, err := join(c.rawURL, methodLogin)

	if err != nil {
		return err
	}

	form := url.Values{}

	form.Add("login", c.auth.username)
	form.Add("password", c.auth.password)

	req, err := http.NewRequestWithContext(ctx, "POST", methodRawURL, strings.NewReader(form.Encode()))

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

func (c *Connection) logout(ctx context.Context) error {
	methodRawURL, err := join(c.rawURL, methodLogout)

	if err != nil {
		return newMethodCallError(methodLogout, "GET", err)
	}

	u, err := url.Parse(methodRawURL)

	if err != nil {
		return newMethodCallError(methodLogout, "GET", err)
	}

	body, err := c.call(ctx, u, "GET")

	if err != nil {
		return newMethodCallError(methodLogout, "GET", err)
	}

	message, err := c.responseStatus(body)

	if err != nil {
		return newMethodCallError(methodLogout, "GET", err)
	}

	if message.StatusCode != http.StatusOK {
		return newInternalError(methodLogout, "GET", message)
	}

	return nil
}

func (c *Connection) call(ctx context.Context, u *url.URL, method string, flags ...Flag) ([]byte, error) {
	compressedResponse := flagExists(CompressedResponse, flags...)
	useCompressedResponseFlagInHeader := flagExists(UseCompressedResponseFlagInHeader, flags...)

	if compressedResponse && !useCompressedResponseFlagInHeader {
		q := u.Query()
		q.Set("compressed_response", "true")

		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Cookie", fmt.Sprintf("access_token=%s", c.token))
	req.Header.Set("key", c.auth.key)

	if compressedResponse && useCompressedResponseFlagInHeader {
		req.Header.Set("compressed-response", "true")
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			c.doUnhandledError(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
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

	if len(messages) == 0 {
		return &response.Message{StatusCode: http.StatusOK, Message: "Success"}, nil
	}

	return messages[0], nil
}

func (c *Connection) doUnhandledError(err error) {
	if c.OnUnhandledError != nil {
		c.OnUnhandledError(err)
	}
}
