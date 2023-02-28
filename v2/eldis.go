package eldis

import (
	"context"
	"net/http"
	"time"
)

// Credentials параметры доступа к API АИИС ЭЛДИС
type Credentials struct {
	// Username имя пользователя
	Username string

	// Password пароль пользователя
	Password string

	// AccessToken токен доступа
	AccessToken string
}

// Connection интерфейс соединения с API АИИС ЭЛДИС
type Connection interface {
	// Open открывает соединение с API АИИС ЭЛДИС
	Open(ctx context.Context, url string, credentials Credentials) error

	// Close закрывает соединение с API АИИС ЭЛДИС
	Close(ctx context.Context) error

	// ListForDevelopment возвращает список доступных пользователю точек учета
	ListForDevelopment(ctx context.Context) ([]byte, error)

	// NormalizedReadings возвращает нормализованные показания точки учета, удовлетворяющие условиям
	NormalizedReadings(ctx context.Context, pointID string, archive Archive, from time.Time, to time.Time,
		dateType DateType) ([]byte, error)

	// RawReadings возвращает "сырые" показания точки учета, удовлетворяющие условиям
	RawReadings(ctx context.Context, pointID string, archive Archive, from time.Time, to time.Time) ([]byte, error)
}

// Connect возвращает соединение с API АИИС ЭЛДИС
func Connect(ctx context.Context, url string, credentials Credentials, options ...ConnectionOption) (Connection, error) {
	conn := &connection{client: http.DefaultClient}

	for _, option := range options {
		option(conn)
	}

	err := conn.Open(ctx, url, credentials)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

// ConnectionOption опция соединения с API АИИС ЭЛДИС
type ConnectionOption func(conn *connection)

// WithHTTPClient устанавливает пользовательский HTTP клиент
func WithHTTPClient(client *http.Client) ConnectionOption {
	return func(conn *connection) {
		if client != nil {
			conn.client = client
		} else {
			conn.client = http.DefaultClient
		}
	}
}
