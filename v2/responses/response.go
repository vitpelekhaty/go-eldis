package responses

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Response базовый тип ответ АИИС ЭЛДИС
type Response struct {
	// Body данные ответа АИСКУТЭ
	Body struct {
		// Messages сообщения о результатах выполнения запроса к API АИСКУТЭ
		Messages []*Message `json:"messages"`
	} `json:"response"`
}

// Messages возвращает сообщения о результатах выполнения запроса в АИИС ЭЛДИС
func (response *Response) Messages() []*Message {
	return response.Body.Messages
}

// Parse возвращает результат разбора базового ответа АИСКУТЭ ЭЛДИС
func Parse(body *bytes.Buffer) (*Response, error) {
	if body == nil || body.Len() == 0 {
		return nil, errors.New("пустой ответ API")
	}

	var response Response

	err := json.Unmarshal(body.Bytes(), &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
