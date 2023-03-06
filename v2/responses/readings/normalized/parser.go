package normalized

import (
	"context"
	"encoding/json"
	"io"
)

// Item элемент архива нормализованных показаний точки учета
type Item struct {
	// E ошибка чтения записи архива показаний
	E error

	i interface{}
}

// IsError возвращает признак того, что элемент архива показаний содержит ошибку разбора записи архива
func (item *Item) IsError() bool {
	return item.E != nil
}

// ColdWaterReadings возвращает запись нормализованных показаний точки учета по холодному водоснабжению
func (item *Item) ColdWaterReadings() (ColdWaterReadings, bool) {
	if r, ok := item.i.(ColdWaterReadings); ok {
		return r, ok
	}

	return nil, false
}

// ElectricityReadings возвращает запись нормализованных показаний точки учета по электричеству
func (item *Item) ElectricityReadings() (ElectricityReadings, bool) {
	if r, ok := item.i.(ElectricityReadings); ok {
		return r, ok
	}

	return nil, false
}

// ElectricityCurrentReadings возвращает запись текущих нормализованных показаний точки учета по электричеству
func (item *Item) ElectricityCurrentReadings() (ElectricityCurrentReadings, bool) {
	if r, ok := item.i.(ElectricityCurrentReadings); ok {
		return r, ok
	}

	return nil, false
}

// GasReadings возвращает запись нормализованных показаний точки учета по газу
func (item *Item) GasReadings() (GasReadings, bool) {
	if r, ok := item.i.(GasReadings); ok {
		return r, ok
	}

	return nil, false
}

// HeatReadings возвращает запись нормализованных показаний точки учета по теплу
func (item *Item) HeatReadings() (HeatReadings, bool) {
	if r, ok := item.i.(HeatReadings); ok {
		return r, ok
	}

	return nil, false
}

// HotWaterReadings возвращает запись нормализованных показаний точки учета по горячему водоснабжению
func (item *Item) HotWaterReadings() (HotWaterReadings, bool) {
	if r, ok := item.i.(HotWaterReadings); ok {
		return r, ok
	}

	return nil, false
}

// WasteWaterReadings возвращает запись нормализованных показаний точки учета по сточным водам
func (item *Item) WasteWaterReadings() (WasteWaterReadings, bool) {
	if r, ok := item.i.(WasteWaterReadings); ok {
		return r, ok
	}

	return nil, false
}

// ParseColdWaterReadings выполняет разбор нормализованных показаний точки учета по холодному водоснабжению
//
// Пример разбора часового архива нормализованного показаний:
//
//	import (
//		"bytes"
//		"context"
//
//		"github.com/vitpelekhaty/go-eldis/v2"
//		"github.com/vitpelekhaty/go-eldis/v2/responses"
//		"github.com/vitpelekhaty/go-eldis/v2/responses/readings/normalized"
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
//		b, _ := conn.NormalizedReadings(ctx, pointID, HourArchive, from, to, Date)
//		sb := responses.Extract(responses.SectionNormalizedColdWater, bytes.NewBuffer(b))
//
//		items, err := normalized.ParseColdWaterReadings(context.Background(), bytes.NewReader(sb))
//
//		for item := range items {
//			if !item.IsError() {
//				...
//			}
//		}
//	}
func ParseColdWaterReadings(ctx context.Context, reader io.Reader) (<-chan *Item, error) {
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

				var readings coldWaterReadings

				if err := decoder.Decode(&readings); err == nil {
					out <- &Item{i: &readings}
				} else {
					out <- &Item{E: err}
				}
			}
		}
	}(decoder)

	return out, nil
}

// ParseElectricityReadings выполняет разбор нормализованных показаний точки учета по электричеству
//
// Пример разбора часового архива нормализованного показаний:
//
//	import (
//		"bytes"
//		"context"
//
//		"github.com/vitpelekhaty/go-eldis/v2"
//		"github.com/vitpelekhaty/go-eldis/v2/responses"
//		"github.com/vitpelekhaty/go-eldis/v2/responses/readings/normalized"
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
//		b, _ := conn.NormalizedReadings(ctx, pointID, HourArchive, from, to, Date)
//		sb := responses.Extract(responses.SectionNormalizedElectricity, bytes.NewBuffer(b))
//
//		items, err := normalized.ParseElectricityReadings(context.Background(), bytes.NewReader(sb))
//
//		for item := range items {
//			if !item.IsError() {
//				...
//			}
//		}
//	}
func ParseElectricityReadings(ctx context.Context, reader io.Reader) (<-chan *Item, error) {
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

				var readings electricityReadings

				if err := decoder.Decode(&readings); err == nil {
					out <- &Item{i: &readings}
				} else {
					out <- &Item{E: err}
				}
			}
		}
	}(decoder)

	return out, nil
}

// ParseElectricityCurrentReadings выполняет разбор текущих нормализованных показаний точки учета по электричеству
//
// Пример разбора часового архива нормализованного показаний:
//
//	import (
//		"bytes"
//		"context"
//
//		"github.com/vitpelekhaty/go-eldis/v2"
//		"github.com/vitpelekhaty/go-eldis/v2/responses"
//		"github.com/vitpelekhaty/go-eldis/v2/responses/readings/normalized"
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
//		b, _ := conn.NormalizedReadings(ctx, pointID, HourArchive, from, to, Date)
//		sb := responses.Extract(responses.SectionNormalizedElectricityCurrent, bytes.NewBuffer(b))
//
//		items, err := normalized.ParseElectricityCurrentReadings(context.Background(), bytes.NewReader(sb))
//
//		for item := range items {
//			if !item.IsError() {
//				...
//			}
//		}
//	}
func ParseElectricityCurrentReadings(ctx context.Context, reader io.Reader) (<-chan *Item, error) {
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

				var readings electricityCurrentReadings

				if err := decoder.Decode(&readings); err == nil {
					out <- &Item{i: &readings}
				} else {
					out <- &Item{E: err}
				}
			}
		}
	}(decoder)

	return out, nil
}

// ParseGasReadings выполняет разбор нормализованных показаний точки учета по газу
//
// Пример разбора часового архива нормализованного показаний:
//
//	import (
//		"bytes"
//		"context"
//
//		"github.com/vitpelekhaty/go-eldis/v2"
//		"github.com/vitpelekhaty/go-eldis/v2/responses"
//		"github.com/vitpelekhaty/go-eldis/v2/responses/readings/normalized"
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
//		b, _ := conn.NormalizedReadings(ctx, pointID, HourArchive, from, to, Date)
//		sb := responses.Extract(responses.SectionNormalizedGas, bytes.NewBuffer(b))
//
//		items, err := normalized.ParseGasReadings(context.Background(), bytes.NewReader(sb))
//
//		for item := range items {
//			if !item.IsError() {
//				...
//			}
//		}
//	}
func ParseGasReadings(ctx context.Context, reader io.Reader) (<-chan *Item, error) {
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

				var readings gasReadings

				if err := decoder.Decode(&readings); err == nil {
					out <- &Item{i: &readings}
				} else {
					out <- &Item{E: err}
				}
			}
		}
	}(decoder)

	return out, nil
}

// ParseHeatReadings выполняет разбор нормализованных показаний точки учета по теплу
//
// Пример разбора часового архива нормализованного показаний:
//
//	import (
//		"bytes"
//		"context"
//
//		"github.com/vitpelekhaty/go-eldis/v2"
//		"github.com/vitpelekhaty/go-eldis/v2/responses"
//		"github.com/vitpelekhaty/go-eldis/v2/responses/readings/normalized"
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
//		b, _ := conn.NormalizedReadings(ctx, pointID, HourArchive, from, to, Date)
//		sb := responses.Extract(responses.SectionNormalizedHeat, bytes.NewBuffer(b))
//
//		items, err := normalized.ParseHeatReadings(context.Background(), bytes.NewReader(sb))
//
//		for item := range items {
//			if !item.IsError() {
//				...
//			}
//		}
//	}
func ParseHeatReadings(ctx context.Context, reader io.Reader) (<-chan *Item, error) {
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

				var readings heatReadings

				if err := decoder.Decode(&readings); err == nil {
					out <- &Item{i: &readings}
				} else {
					out <- &Item{E: err}
				}
			}
		}
	}(decoder)

	return out, nil
}

// ParseHotWaterReadings выполняет разбор нормализованных показаний точки учета по горячему водоснабжению
//
// Пример разбора часового архива нормализованного показаний:
//
//	import (
//		"bytes"
//		"context"
//
//		"github.com/vitpelekhaty/go-eldis/v2"
//		"github.com/vitpelekhaty/go-eldis/v2/responses"
//		"github.com/vitpelekhaty/go-eldis/v2/responses/readings/normalized"
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
//		b, _ := conn.NormalizedReadings(ctx, pointID, HourArchive, from, to, Date)
//		sb := responses.Extract(responses.SectionNormalizedHotWater, bytes.NewBuffer(b))
//
//		items, err := normalized.ParseHotWaterReadings(context.Background(), bytes.NewReader(sb))
//
//		for item := range items {
//			if !item.IsError() {
//				...
//			}
//		}
//	}
func ParseHotWaterReadings(ctx context.Context, reader io.Reader) (<-chan *Item, error) {
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

				var readings hotWaterReadings

				if err := decoder.Decode(&readings); err == nil {
					out <- &Item{i: &readings}
				} else {
					out <- &Item{E: err}
				}
			}
		}
	}(decoder)

	return out, nil
}

// ParseWasteWaterReadings выполняет разбор нормализованных показаний точки учета по сточным водам
//
// Пример разбора часового архива нормализованного показаний:
//
//	import (
//		"bytes"
//		"context"
//
//		"github.com/vitpelekhaty/go-eldis/v2"
//		"github.com/vitpelekhaty/go-eldis/v2/responses"
//		"github.com/vitpelekhaty/go-eldis/v2/responses/readings/normalized"
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
//		b, _ := conn.NormalizedReadings(ctx, pointID, HourArchive, from, to, Date)
//		sb := responses.Extract(responses.SectionNormalizedWasteWater, bytes.NewBuffer(b))
//
//		items, err := normalized.ParseWasteWaterReadings(context.Background(), bytes.NewReader(sb))
//
//		for item := range items {
//			if !item.IsError() {
//				...
//			}
//		}
//	}
func ParseWasteWaterReadings(ctx context.Context, reader io.Reader) (<-chan *Item, error) {
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

				var readings wasteWaterReadings

				if err := decoder.Decode(&readings); err == nil {
					out <- &Item{i: &readings}
				} else {
					out <- &Item{E: err}
				}
			}
		}
	}(decoder)

	return out, nil
}
