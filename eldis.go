package eldis

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-eldis/archive"
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

// Connection соединение с API ЭЛДИС
type Connection struct {
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
	return nil
}

// Close закрывает соединение с API ЭЛДИС
func (c *Connection) Close() error {
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
