package points

import "time"

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

// Point точка учета
type Point map[string]interface{}

// ID возвращает идентификатор точки учёта
func (point Point) ID() string {
	return point["id"].(string)
}

// DeviceID возвращает идентификатор прибора учёта
func (point Point) DeviceID() string {
	return point["deviceID"].(string)
}

// ObjectID возвращает идентификатор объекта
func (point Point) ObjectID() string {
	return point["objectID"].(string)
}

// Status возвращает состояние точки учёта
func (point Point) Status() PointStatus {
	return PointStatus(point["status"].(int))
}

// CreatedOn возвращает дату создания точки учёта
func (point Point) CreatedOn() time.Time {
	createdOn := point["createdOn"].(int64)
	return time.Unix(createdOn, 0).UTC()
}

// Identifier возвращает дополнительный идентификатор точки учёта
func (point Point) Identifier() (string, bool) {
	if identifier, ok := point["identifier"]; ok {
		return identifier.(string), ok
	}

	return "", false
}

// Identifier2 возвращает дополнительный идентификатор 2 точки учёта
func (point Point) Identifier2() (string, bool) {
	if identifier, ok := point["identifier2"]; ok {
		return identifier.(string), ok
	}

	return "", false
}

// Address возвращает адрес объекта
func (point Point) Address() string {
	return point["address"].(string)
}

// DeviceName возвращает название прибора учёта
func (point Point) DeviceName() string {
	return point["deviceName"].(string)
}

// SN возвращает серийный номер прибора учёта
func (point Point) SN() string {
	return point["sn"].(string)
}

// ResourceName возвращает название ресурса
func (point Point) ResourceName() string {
	return point["resourceName"].(string)
}

// MeasurePointNumber возвращает номер точки учёта
func (point Point) MeasurePointNumber() string {
	return point["measurePointNumber"].(string)
}

// MeasurePointName возвращает название точки учёта
func (point Point) MeasurePointName() string {
	return point["measurePointName"].(string)
}

// IsGVS возвращает признак, что точка учета горячего водоснабжения
//
// Не обязательный параметр, поэтому дополнительно функция возвращает признак наличия параметра
func (point Point) IsGVS() (bool, bool) {
	if gvs, ok := point["isGVS"]; ok {
		return gvs.(bool), ok
	}

	return false, false
}

// IsHeat возвращает признак, что точка учета отопления
//
// Не обязательный параметр, поэтому дополнительно функция возвращает признак наличия параметра
func (point Point) IsHeat() (bool, bool) {
	if heat, ok := point["isHeat"]; ok {
		return heat.(bool), ok
	}

	return false, false
}

// InputConfiguration возвращает конфигурацию ввода
//
// Не обязательный параметр, поэтому дополнительно функция возвращает признак наличия параметра
func (point Point) InputConfiguration() (int, bool) {
	if input, ok := point["inputConfiguration"]; ok {
		return input.(int), ok
	}

	return 0, false
}
