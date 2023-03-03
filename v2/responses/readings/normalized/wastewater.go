package normalized

import (
	"github.com/guregu/null"
)

var _ WasteWaterReadings = (*wasteWaterReadings)(nil)

type wasteWaterReadings struct {
	custom

	// V_ объём, м3
	V_ null.Float `json:"V,omitempty"`

	// VTotal_ накопленное значение объёма с момента сброса, м3
	VTotal_ null.Float `json:"V_Total,omitempty"`

	// M_ масса, т
	M_ null.Float `json:"M,omitempty"`

	// MTotal_ накопленное значение массы с момента сброса, т
	MTotal_ null.Float `json:"M_Total,omitempty"`
}

// V возвращает объём, м3
func (readings *wasteWaterReadings) V() (float64, bool) {
	return readings.V_.Float64, readings.V_.Valid
}

// VTotal возвращает накопленное значение объёма с момента сброса, м3
func (readings *wasteWaterReadings) VTotal() (float64, bool) {
	return readings.VTotal_.Float64, readings.VTotal_.Valid
}

// M возвращает массу, т
func (readings *wasteWaterReadings) M() (float64, bool) {
	return readings.M_.Float64, readings.M_.Valid
}

// MTotal возвращает накопленное значение массы с момента сброса, т
func (readings *wasteWaterReadings) MTotal() (float64, bool) {
	return readings.MTotal_.Float64, readings.MTotal_.Valid
}
