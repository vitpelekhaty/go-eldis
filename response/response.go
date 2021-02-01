package response

import (
	"errors"

	"github.com/tidwall/gjson"
)

// Section раздел ответа метода API ЭЛДИС
type Section byte

const (
	// SectionListForDevelopment раздел доступных точек учета
	SectionListForDevelopment Section = iota
	// SectionNormalizedWasteWaterData раздел нормализованных показаний по сточным водам
	SectionNormalizedWasteWaterData
	// SectionNormalizedColdWaterData раздел нормализованных показаний по холодной воде
	SectionNormalizedColdWaterData
	// SectionNormalizedHotWaterData раздел нормализованных показаний по горячей воде
	SectionNormalizedHotWaterData
	// SectionNormalizedHeatData раздел нормализованных показаний по теплу
	SectionNormalizedHeatData
	// SectionNormalizedElectricityData раздел нормализованных показаний по электричеству
	SectionNormalizedElectricityData
	// SectionNormalizedElectricityCurrentData раздел нормализованных текущих показаний по электричеству
	SectionNormalizedElectricityCurrentData
	// SectionNormalizedGasData раздел нормализованных показаний по газу
	SectionNormalizedGasData
	// SectionUoM раздел групп справочников единиц измерений
	SectionUoM
	// SectionRawData раздел "сырых" показаний прибора учета
	SectionRawData
)

// Get извлекает из тела body ответа метода API ЭЛДИС содержимое указанного раздела section. Если ответ не содержит
// указанный раздел, то возвращается ошибка типа PathError
func Get(section Section, body []byte) ([]byte, error) {
	switch section {
	case SectionListForDevelopment:
		return getBytes(body, listForDevelopmentPath)
	case SectionNormalizedWasteWaterData:
		return getBytes(body, normalizedWasteWaterPath)
	case SectionNormalizedColdWaterData:
		return getBytes(body, normalizedColdWaterPath)
	case SectionNormalizedHotWaterData:
		return getBytes(body, normalizedHotWaterPath)
	case SectionNormalizedHeatData:
		return getBytes(body, normalizedHeatPath)
	case SectionNormalizedElectricityData:
		return getBytes(body, normalizedElectricityPath)
	case SectionNormalizedElectricityCurrentData:
		return getBytes(body, normalizedElectricityCurrentPath)
	case SectionNormalizedGasData:
		return getBytes(body, normalizedGasPath)
	case SectionUoM:
		return getBytes(body, uomPath)
	case SectionRawData:
		return getBytes(body, rawDataPath)
	default:
		return nil, errors.New("unavailable for this section")
	}
}

func getBytes(body []byte, path string) ([]byte, error) {
	var raw []byte

	section := gjson.GetBytes(body, path)

	if section.Index > 0 {
		raw = body[section.Index : section.Index+len(section.Raw)]
	} else {
		raw = []byte(section.Raw)
	}

	if len(raw) == 0 {
		return nil, &PathError{Path: path}
	}

	return raw, nil
}
