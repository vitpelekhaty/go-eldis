package response

import (
	"encoding/json"
	"errors"
)

// LoginResponse ответ метода api/v2/users/login
type LoginResponse struct {
	Response struct {
		Users struct {
			Login struct {
				Result bool `json:"result"`
			} `json:"login"`
		} `json:"users"`

		Messages []Message `json:"messages"`
	} `json:"response"`
}

// Result результат авторизации
func (r *LoginResponse) Result() bool {
	return r.Response.Users.Login.Result
}

// Messages сообщения о результатах выполнения запроса в АИСКУТЭ ЭЛДИС
func (r *LoginResponse) Messages() []Message {
	return r.Response.Messages
}

// ParseLoginResponse возвращает результат разбора ответа метода /api/v2/users/login
func ParseLoginResponse(body []byte) (*LoginResponse, error) {
	if len(body) == 0 {
		return nil, errors.New("no data")
	}

	var resp LoginResponse

	err := json.Unmarshal(body, &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
