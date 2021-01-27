package response

import (
	"encoding/json"
	"errors"
)

// CustomResponse базовый ответ АИСКУТЭ ЭЛДИС
type CustomResponse struct {
	// Response данные ответа АИСКУТЭ
	Response struct {
		// Messages сообщения о результатах выполнения запроса к API АИСКУТЭ
		Messages []*Message `json:"messages"`
	} `json:"response"`
}

// Messages возвращает сообщения о результатах выполнения запроса в АИСКУТЭ ЭЛДИС
func (r *CustomResponse) Messages() []*Message {
	return r.Response.Messages
}

// Parse возвращает результат разбора базового ответа АИСКУТЭ ЭЛДИС
func Parse(body []byte) (*CustomResponse, error) {
	if len(body) == 0 {
		return nil, errors.New("no data")
	}

	var resp CustomResponse

	err := json.Unmarshal(body, &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
