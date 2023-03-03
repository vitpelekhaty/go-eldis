package raw

import (
	"context"
	"encoding/json"
	"io"
)

// ReadingsRecord запись архива показаний непосредственно с приборов учета
//
// Набор полей записи архива показаний зависит от типа прибора учета, поэтому, чтобы абстрагироваться от типа прибора,
// используется map
type ReadingsRecord map[string]interface{}

// Item элемент архива показаний непосредственно с приборов учета
type Item struct {
	// Record запись архива показаний непосредственно с приборов учета
	Record ReadingsRecord

	// E ошибка чтения записи архива показаний
	E error
}

// IsError возвращает признак того, что элемент архива показаний содержит ошибку разбора записи архива
func (item *Item) IsError() bool {
	return item.E != nil
}

// Parse выполняет разбор необработанных в АИИС ЭЛДИС показаний на точке учета
//
// Пример кода:
//
//	import (
//		"bytes"
//		"context"
//
//		"github.com/vitpelekhaty/go-eldis/v2"
//		"github.com/vitpelekhaty/go-eldis/v2/responses"
//		"github.com/vitpelekhaty/go-eldis/v2/responses/readings/raw"
//	)
//
//	func main() {
//		...
//		conn, err := eldis.Connect(ctx, rawURL, Credentials{Username: username, Password: password, AccessToken: accessToken})
//
//		defer func() {
//			_ = conn.Close(ctx)
//		}()
//
//		b, _ := conn.RawReadings(ctx, pointID, HourArchive, from, to)
//		sb := responses.Extract(responses.SectionRaw, bytes.NewBuffer(b))
//
//		items, _ := raw.Parse(context.Background(), bytes.NewReader(sb))
//
//		for item := range items {
//			if !item.IsError() {
//				...
//			}
//		}
//	}
func Parse(ctx context.Context, reader io.Reader) (<-chan *Item, error) {
	decoder := json.NewDecoder(reader)

	_, err := decoder.Token()

	if err != nil {
		return nil, err
	}

	out := make(chan *Item)

	go func(decoder *json.Decoder) {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !decoder.More() {
					return
				}

				var record ReadingsRecord

				if err := decoder.Decode(&record); err == nil {
					out <- &Item{Record: record}
				} else {
					out <- &Item{E: err}
				}
			}
		}
	}(decoder)

	return out, nil
}
