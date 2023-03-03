package points

import (
	"context"
	"encoding/json"
	"io"
)

// Item элемент списка точек учета
type Item struct {
	// P точка учета
	P Point

	// E ошибка чтения свойств точки учета
	E error
}

// IsError возвращает признак того, что элемент списка точек учета содержит ошибку чтения свойств точки учета
func (item *Item) IsError() bool {
	return item.E != nil
}

// Parse выполняет разбор списка точек учета
//
// Пример кода:
//
//	import (
//		"bytes"
//		"context"
//
//		"github.com/vitpelekhaty/go-eldis/v2"
//		"github.com/vitpelekhaty/go-eldis/v2/responses"
//		"github.com/vitpelekhaty/go-eldis/v2/responses/points"
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
//		b, _ := conn.ListForDevelopment(ctx)
//		sb := responses.Extract(responses.SectionListForDevelopment, bytes.NewBuffer(b))
//
//		items, _ := points.Parse(context.Background(), bytes.NewReader(sb))
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

				var props properties

				if err := decoder.Decode(&props); err == nil {
					out <- &Item{P: &point{properties: &props}}
				} else {
					out <- &Item{E: err}
				}
			}
		}
	}(decoder)

	return out, nil
}
