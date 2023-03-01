package responses

import (
	"bytes"

	"github.com/tidwall/gjson"
)

// Section раздел ответа метода API ЭЛДИС
type Section string

const (
	// SectionListForDevelopment раздел доступных точек учета
	SectionListForDevelopment Section = "response.tv.listForDevelopment"

	// SectionNormalizedWasteWater раздел нормализованных показаний по сточным водам
	SectionNormalizedWasteWater Section = "response.data.normalized.#(wasteWater).wasteWater"

	// SectionNormalizedColdWater раздел нормализованных показаний по холодной воде
	SectionNormalizedColdWater Section = "response.data.normalized.#(coldWater).coldWater"

	// SectionNormalizedHotWater раздел нормализованных показаний по горячей воде
	SectionNormalizedHotWater Section = "response.data.normalized.#(hotWater).hotWater"

	// SectionNormalizedHeat раздел нормализованных показаний по теплу
	SectionNormalizedHeat Section = "response.data.normalized.#(heat).heat"

	// SectionNormalizedElectricity раздел нормализованных показаний по электричеству
	SectionNormalizedElectricity Section = "response.data.normalized.#(electricity).electricity"

	// SectionNormalizedElectricityCurrent раздел нормализованных текущих показаний по электричеству
	SectionNormalizedElectricityCurrent Section = "response.data.normalized.#(electricityCurrent).electricityCurrent"

	// SectionNormalizedGas раздел нормализованных показаний по газу
	SectionNormalizedGas Section = "response.data.normalized.#(gas).gas"

	// SectionRaw раздел "сырых" показаний прибора учета
	SectionRaw Section = "response.data.rawData"
)

// Extract возвращает извлеченное содержимое указанной секции из буфера
func Extract(section Section, buff *bytes.Buffer) []byte {
	if buff == nil || buff.Len() == 0 {
		return make([]byte, 0)
	}

	data := buff.Bytes()

	result := gjson.GetBytes(data, string(section))

	if result.Index > 0 {
		return data[result.Index : result.Index+len(result.Raw)]
	}

	return []byte(result.Raw)
}
