package response

// MessageLevel уровень сообщения
type MessageLevel byte

const (
	// MessageLevelError ошибка
	MessageLevelError MessageLevel = iota
	// MessageLevelWarning предупреждение
	MessageLevelWarning
	// MessageLevelMessage сообщение
	MessageLevelMessage
)

// Message сообщение о результатах выполнения запроса в АИСКУТЭ ЭЛДИС
type Message struct {
	// StatusCode код HTTP ответа
	StatusCode int `json:"httpStatusCode"`
	// Code код сообщения
	Code int `json:"messageCode"`
	// Level уровень сообщения
	Level MessageLevel `json:"messageLevel"`
	// ID идентификатор записи, к которой относится сообщение
	ID string `json:"id"`
	// Name название записи
	Name string `json:"name"`
	// Message текст сообщения
	Message string `json:"message"`
}
