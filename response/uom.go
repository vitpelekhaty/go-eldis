package response

import (
	"bytes"
	"context"
	"encoding/json"
)

// UoM единица измерений (Unit of Measurement)
type UoM struct {
	// ID идентификатор единицы измерения
	ID string `json:"id"`
	// Code код единицы измерения
	Code int `json:"code"`
	// Name название единицы измерения
	Name string `json:"name"`
}

// UoMGroup группа единиц измерения
type UoMGroup struct {
	// Name название группы
	Name string `json:"name"`
	// UOM список единиц измерений
	UOM []*UoM `json:"uom"`
}

const uomPath = "response.uom.list"

// ParseUoMGroupsWithContext возвращает канал, в который будет записывать группы справочников единиц измерения в
// АИСКУТЭ ЭЛДИС, полученных в ответе API при вызове метода /api/v2/uom/list
//
// body - ответ указанного метода API; ctx - контекст, с помощью которого можно при необходимости остановить
// перечисление точек учета
//
// Перечисление групп справочников будет выполняться до полного чтения всех групп из ответа или до вызова функции
// cancel отмены контекста
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseUoMGroupsWithContext(ctx context.Context, body []byte) (<-chan struct {
	RegPoint *UoMGroup
	Err      error
}, error) {
	raw, err := getBytes(body, uomPath)

	if err != nil {
		return nil, err
	}

	out := make(chan struct {
		RegPoint *UoMGroup
		Err      error
	})

	go func(b []byte) {
		defer close(out)

		decoder := json.NewDecoder(bytes.NewReader(b))

		_, err := decoder.Token()

		if err != nil {
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

				var group UoMGroup

				if err := decoder.Decode(&group); err == nil {
					out <- struct {
						RegPoint *UoMGroup
						Err      error
					}{RegPoint: &group, Err: nil}
				} else {
					out <- struct {
						RegPoint *UoMGroup
						Err      error
					}{RegPoint: nil, Err: err}
				}
			}
		}
	}(raw)

	return out, nil
}

// ParseUoMGroups возвращает канал, в который будет записывать группы справочников единиц измерения в
// АИСКУТЭ ЭЛДИС, полученных в ответе API при вызове метода /api/v2/uom/list
//
// body - ответ указанного метода API
//
// Перечисление групп справочников будет выполняться до полного чтения всех групп из ответа
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseUoMGroups(body []byte) (<-chan struct {
	RegPoint *UoMGroup
	Err      error
}, error) {
	return ParseUoMGroupsWithContext(context.TODO(), body)
}
