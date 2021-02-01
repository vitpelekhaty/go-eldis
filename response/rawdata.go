package response

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/guregu/null"
)

// RawValue "сырое" значение поля показаний прибора учета
type RawValue string

// AsInt значение типа int поля "сырых" показаний прибора учета. Если поле не содержит значения типа int, то в
// результате возвращается ошибка преобразования типов
func (v *RawValue) AsInt() (int, error) {
	val, err := strconv.Atoi(string(*v))

	if err != nil {
		return -1, err
	}

	return val, err
}

// String строковое значениеполя "сырых" показаний прибора учета
func (v *RawValue) String() (string, error) {
	return string(*v), nil
}

// AsTime значение типа time поля "сырых" показаний прибора учета. Если поле не содержит значения типа time, то в
// результате возвращается ошибка преобразования типов или ошибка определения layout
func (v *RawValue) AsTime() (time.Time, error) {
	return dateparse.ParseAny(string(*v), dateparse.PreferMonthFirst(true))
}

// AsFloat значение типа float64 поля "сырых" показаний прибора учета. Если поле не содержит значения типа float, то в
// результате возвращается ошибка преобразования типов
func (v *RawValue) AsFloat() (float64, error) {
	val := strings.ReplaceAll(string(*v), ",", ".")
	return strconv.ParseFloat(val, 64)
}

// RawRecord "сырые" данные из ответа метода api/v2/data/rawData
type RawRecord map[string]RawValue

// rawRecord "сырые" данных из ответа метода api/v2/data/rawData
type rawRecord map[string]null.String

const rawDataPath = "response.data.rawData"

// ParseRawDataWithContext возвращает канал, в который будет записывать "сырые" показания приборов учета в
// АИСКУТЭ ЭЛДИС, полученных в ответе API при вызове метода /api/v2/data/rawData
//
// body - ответ указанного метода API; ctx - контекст, с помощью которого можно при необходимости остановить
// перечисление точек учета
//
// Перечисление показаний приборов учета будет выполняться до полного чтения всех показаний из ответа или до вызова
// функции cancel отмены контекста
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseRawDataWithContext(ctx context.Context, body []byte) (<-chan struct {
	Record RawRecord
	Err    error
}, error) {
	raw, err := getBytes(body, rawDataPath)

	if err != nil {
		return nil, err
	}

	out := make(chan struct {
		Record RawRecord
		Err    error
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

				buffer := make(rawRecord)

				if err := decoder.Decode(&buffer); err == nil {
					record := make(RawRecord)

					for key, value := range buffer {
						if value.Valid {
							record[key] = RawValue(value.String)
						}
					}

					out <- struct {
						Record RawRecord
						Err    error
					}{Record: record, Err: nil}
				} else {
					out <- struct {
						Record RawRecord
						Err    error
					}{Record: nil, Err: err}
				}
			}
		}
	}(raw)

	return out, nil
}

// ParseRawData возвращает канал, в который будет записывать "сырые" показания приборов учета в
// АИСКУТЭ ЭЛДИС, полученных в ответе API при вызове метода /api/v2/data/rawData
//
// body - ответ указанного метода API
//
// Перечисление показаний приборов учета будет выполняться до полного чтения всех показаний из ответа или до вызова
// функции cancel отмены контекста
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseRawData(body []byte) (<-chan struct {
	Record RawRecord
	Err    error
}, error) {
	return ParseRawDataWithContext(context.TODO(), body)
}
