package normalized

import (
	"github.com/guregu/null"
)

var _ HeatReadings = (*heatReadings)(nil)

type heatReadings struct {
	hotWaterReadings

	// T3_ температура на подпитке, °C
	T3_ null.Float `json:"t3,omitempty"`

	// V3_ объём на подпитке, м3
	V3_ null.Float `json:"V3,omitempty"`

	// V3Total_ накопленное значение объёма на подпитке с момента сброса, м3
	V3Total_ null.Float `json:"V3_Total,omitempty"`

	// M3_ масса на подпитке, т
	M3_ null.Float `json:"M3,omitempty"`

	// M3Total_ накопленное значение массы на подпитке с момента сброса, т
	M3Total_ null.Float `json:"M3_Total,omitempty"`

	// P3_ давление на подпитке, МПа
	P3_ null.Float `json:"P3,omitempty"`

	// Ta_ температура наружного воздуха, °C
	Ta_ null.Float `json:"ta,omitempty"`
}

// T3 возвращает температуру на подпитке, °C
func (readings *heatReadings) T3() (float64, bool) {
	return readings.T3_.Float64, readings.T3_.Valid
}

// V3 возвращает объём на подпитке, м3
func (readings *heatReadings) V3() (float64, bool) {
	return readings.V3_.Float64, readings.V3_.Valid
}

// V3Total возвращает накопленное значение объёма на подпитке с момента сброса, м3
func (readings *heatReadings) V3Total() (float64, bool) {
	return readings.V3Total_.Float64, readings.V3Total_.Valid
}

// M3 возвращает массу на подпитке, т
func (readings *heatReadings) M3() (float64, bool) {
	return readings.M3_.Float64, readings.M3_.Valid
}

// M3Total возвращает накопленное значение массы на подпитке с момента сброса, т
func (readings *heatReadings) M3Total() (float64, bool) {
	return readings.M3Total_.Float64, readings.M3Total_.Valid
}

// P3 возвращает давление на подпитке, МПа
func (readings *heatReadings) P3() (float64, bool) {
	return readings.P3_.Float64, readings.P3_.Valid
}

// Ta возвращает температуру наружного воздуха, °C
func (readings *heatReadings) Ta() (float64, bool) {
	return readings.Ta_.Float64, readings.Ta_.Valid
}
