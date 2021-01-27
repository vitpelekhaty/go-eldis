package response

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/guregu/null"
	"github.com/tidwall/gjson"
)

// RegPointStatus состояние точки учета
type RegPointStatus byte

const (
	// RegPointStatusActive активная точка учета
	RegPointStatusActive RegPointStatus = iota
	// RegPointStatusInactive неактивная точка учета
	RegPointStatusInactive
)

// RegPoint описание точки учета в АИСКУТЭ ЭЛДИС
type RegPoint struct {
	// ID идентификатор точки учёта
	ID string `json:"id"`
	// DeviceID идентификатор прибора учёта
	DeviceID string `json:"deviceID"`
	// ObjectID идентификатор объекта
	ObjectID string `json:"objectID"`
	// Status состояние точки учёта
	Status RegPointStatus `json:"status"`
	// CreatedOn дата создания точки учёта в unix timestamp
	CreatedOn int64 `json:"createdOn"`
	// Identifier2 дополнительный идентификатор точки учёта
	Identifier2 string `json:"identifier2"`
	// Address адрес объекта
	Address string `json:"address"`
	// Identifier идентификатор точки учёта
	Identifier string `json:"identifier"`
	// DeviceName название прибора учёта
	DeviceName string `json:"deviceName"`
	// SN серийный номер прибора учёта
	SN string `json:"sn"`
	// ResourceName название ресурса
	ResourceName string `json:"resourceName"`
	// MeasurePointNumber номер точки учёта
	MeasurePointNumber string `json:"measurePointNumber"`
	// MeasurePointName название точки учёта
	MeasurePointName string `json:"measurePointName"`
	// IsGVS ГВС
	IsGVS null.Bool `json:"isGVS,omitempty"`
	// InputConfiguration конфигурация ввода
	InputConfiguration null.Int `json:"inputConfiguration,omitempty"`
	// IsHeat отопление
	IsHeat null.Bool `json:"isHeat,omitempty"`
}

// Created время создания точки учета в АИСКУТЭ ЭЛДИС
func (rp *RegPoint) Created() time.Time {
	return time.Unix(rp.CreatedOn, 0).UTC()
}

// ParseRegPointsWithContext возвращает канал, в который будет записывать описания точек учета АИСКУТЭ ЭЛДИС,
// полученных в ответе API при вызове метода /api/v2/tv/listForDevelopment.
//
// body - ответ указанного метода API; ctx - контекст, с помощью которого можно при необходимости остановить
// перечисление точек учета
//
// Перечисление точек учета будет выполняться до полного чтения всех точек учета из ответа; до вызова функции cancel
// отмены контекста; до первой ошибки разбора ответа метода API
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseRegPointsWithContext(ctx context.Context, body []byte) <-chan struct {
	RegPoint *RegPoint
	Err      error
} {
	out := make(chan struct {
		RegPoint *RegPoint
		Err      error
	})

	go func() {
		defer close(out)

		listForDevelopment := gjson.GetBytes(body, "response.tv.listForDevelopment")

		var raw []byte

		if listForDevelopment.Index > 0 {
			raw = body[listForDevelopment.Index : listForDevelopment.Index+len(listForDevelopment.Raw)]
		} else {
			raw = []byte(listForDevelopment.Raw)
		}

		decoder := json.NewDecoder(bytes.NewReader(raw))

		_, err := decoder.Token()

		if err != nil {
			out <- struct {
				RegPoint *RegPoint
				Err      error
			}{RegPoint: nil, Err: err}

			return
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !decoder.More() {
					return
				}

				var regPoint RegPoint

				if err := decoder.Decode(&regPoint); err == nil {
					out <- struct {
						RegPoint *RegPoint
						Err      error
					}{RegPoint: &regPoint, Err: nil}
				} else {
					out <- struct {
						RegPoint *RegPoint
						Err      error
					}{RegPoint: nil, Err: err}

					return
				}
			}
		}
	}()

	return out
}

// ParseRegPoints возвращает канал, в который будет записывать описания точек учета АИСКУТЭ ЭЛДИС, полученных в ответе
// API при вызове метода /api/v2/tv/listForDevelopment.
//
// body - ответ указанного метода API
//
// Перечисление точек учета будет выполняться до полного чтения всех точек учета из ответа; до первой ошибки разбора
// ответа метода API
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseRegPoints(body []byte) <-chan struct {
	RegPoint *RegPoint
	Err      error
} {
	return ParseRegPointsWithContext(context.TODO(), body)
}
