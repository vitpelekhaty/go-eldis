package response

import (
	"encoding/json"
)

// LogoutResponse ответ метода api/v2/users/logout
type LogoutResponse struct {
	Response struct {
		Users struct {
			Logout struct {
				Result bool `json:"result"`
			} `json:"logout"`
		} `json:"users"`

		Messages []Message `json:"messages"`
	} `json:"response"`
}

// Result результат операции
func (r *LogoutResponse) Result() bool {
	return r.Response.Users.Logout.Result
}

// Messages сообщения о результатах выполнения запроса в АИСКУТЭ ЭЛДИС
func (r *LogoutResponse) Messages() []Message {
	return r.Response.Messages
}

// ParseLogoutResponse возвращает результат разбора ответа метода /api/v2/users/logout
func ParseLogoutResponse(body []byte) (*LogoutResponse, error) {
	if len(body) == 0 {
		return nil, errEmptyBody
	}

	var resp LogoutResponse

	err := json.Unmarshal(body, &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
