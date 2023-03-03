package points

import (
	"time"

	"github.com/guregu/null"
)

// PointStatus состояние точки учета
type PointStatus byte

const (
	// PointStatusActive активная точка учета
	PointStatusActive PointStatus = iota + 1

	// PointStatusInactive неактивная точка учета
	PointStatusInactive

	// PointStatusMeasurementSchemeChanged на точке учета изменена схема измерения
	PointStatusMeasurementSchemeChanged

	// PointStatusWithoutAutoReading точка учета без автоопроса показаний
	PointStatusWithoutAutoReading
)

// IPoint интерфейс точки учета
type IPoint interface {
	// ID возвращает идентификатор точки учёта
	ID() string

	// DeviceID возвращает идентификатор прибора учёта
	DeviceID() string

	// ObjectID возвращает идентификатор объекта
	ObjectID() string

	// Status возвращает состояние точки учёта
	Status() PointStatus

	// CreatedOn возвращает дату создания точки учёта
	CreatedOn() time.Time

	// Identifier возвращает дополнительный идентификатор точки учёта
	Identifier() string

	// Identifier2 возвращает дополнительный идентификатор 2 точки учёта
	Identifier2() string

	// Address возвращает адрес объекта
	Address() string

	// DeviceName возвращает название прибора учёта
	DeviceName() string

	// SN возвращает серийный номер прибора учёта
	SN() string

	// ResourceName возвращает название ресурса
	ResourceName() string

	// MeasurePointNumber возвращает номер точки учёта
	MeasurePointNumber() string

	// MeasurePointName возвращает название точки учёта
	MeasurePointName() string

	// IsGVS возвращает признак, что точка учета горячего водоснабжения
	//
	// Не обязательный параметр, поэтому дополнительно функция возвращает признак наличия параметра
	IsGVS() (bool, bool)

	// IsHeat возвращает признак, что точка учета отопления
	//
	// Не обязательный параметр, поэтому дополнительно функция возвращает признак наличия параметра
	IsHeat() (bool, bool)

	// InputConfiguration возвращает конфигурацию ввода
	//
	// Не обязательный параметр, поэтому дополнительно функция возвращает признак наличия параметра
	InputConfiguration() (int, bool)
}

var _ IPoint = (*point)(nil)

type point struct {
	properties *properties
}

type properties struct {
	// ID идентификатор точки учёта
	ID string `json:"id"`

	// DeviceID идентификатор прибора учёта
	DeviceID string `json:"deviceID"`

	// ObjectID идентификатор объекта
	ObjectID string `json:"objectID"`

	// Status состояние точки учёта
	Status PointStatus `json:"status"`

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

// ID возвращает идентификатор точки учёта
func (p *point) ID() string {
	return p.properties.ID
}

// DeviceID возвращает идентификатор прибора учёта
func (p *point) DeviceID() string {
	return p.properties.DeviceID
}

// ObjectID возвращает идентификатор объекта
func (p *point) ObjectID() string {
	return p.properties.ObjectID
}

// Status возвращает состояние точки учёта
func (p *point) Status() PointStatus {
	return p.properties.Status
}

// CreatedOn возвращает дату создания точки учёта
func (p *point) CreatedOn() time.Time {
	return time.Unix(p.properties.CreatedOn, 0).UTC()
}

// Identifier возвращает дополнительный идентификатор точки учёта
func (p *point) Identifier() string {
	return p.properties.Identifier
}

// Identifier2 возвращает дополнительный идентификатор 2 точки учёта
func (p *point) Identifier2() string {
	return p.properties.Identifier2
}

// Address возвращает адрес объекта
func (p *point) Address() string {
	return p.properties.Address
}

// DeviceName возвращает название прибора учёта
func (p *point) DeviceName() string {
	return p.properties.DeviceName
}

// SN возвращает серийный номер прибора учёта
func (p *point) SN() string {
	return p.properties.SN
}

// ResourceName возвращает название ресурса
func (p *point) ResourceName() string {
	return p.properties.ResourceName
}

// MeasurePointNumber возвращает номер точки учёта
func (p *point) MeasurePointNumber() string {
	return p.properties.MeasurePointNumber
}

// MeasurePointName возвращает название точки учёта
func (p *point) MeasurePointName() string {
	return p.properties.MeasurePointName
}

// IsGVS возвращает признак, что точка учета горячего водоснабжения
//
// Не обязательный параметр, поэтому дополнительно функция возвращает признак наличия параметра
func (p *point) IsGVS() (bool, bool) {
	return p.properties.IsGVS.Bool, p.properties.IsGVS.Valid
}

// IsHeat возвращает признак, что точка учета отопления
//
// Не обязательный параметр, поэтому дополнительно функция возвращает признак наличия параметра
func (p *point) IsHeat() (bool, bool) {
	return p.properties.IsHeat.Bool, p.properties.IsHeat.Valid
}

// InputConfiguration возвращает конфигурацию ввода
//
// Не обязательный параметр, поэтому дополнительно функция возвращает признак наличия параметра
func (p *point) InputConfiguration() (int, bool) {
	return int(p.properties.InputConfiguration.Int64), p.properties.InputConfiguration.Valid
}
