package normalized

import (
	"github.com/guregu/null"
)

var _ ColdWaterReadings = (*coldWaterReadings)(nil)

// coldWaterReadings нормализованные данные по ХВС
type coldWaterReadings struct {
	wasteWaterReadings

	// P_ давление, МПа
	P_ null.Float `json:"P,omitempty"`
}

// P возвращает давление, МПа
func (readings *coldWaterReadings) P() (float64, bool) {
	return readings.P_.Float64, readings.P_.Valid
}
