package normalized

import "github.com/guregu/null"

var _ GasReadings = (*gasReadings)(nil)

type gasReadings struct {
	custom

	// T_ температура, °C
	T_ null.Float `json:"t,omitempty"`

	// P_ давление, МПа
	P_ null.Float `json:"P,omitempty"`

	// Vp_ рабочий объём, м3
	Vp_ null.Float `json:"Vp,omitempty"`

	// VpTotal_ накопленное значение рабочего объёма с момента сброса, м3
	VpTotal_ null.Float `json:"Vp_Total,omitempty"`

	// Vc стандартный объём, м3
	Vc_ null.Float `json:"Vc,omitempty"`

	// VcTotal_ накопленное значение стандартного объёма с момента сброса, м3
	VcTotal_ null.Float `json:"Vc_Total,omitempty"`
}

// T возвращает температуру, °C
func (readings *gasReadings) T() (float64, bool) {
	return readings.T_.Float64, readings.T_.Valid
}

// P возвращает давление, МПа
func (readings *gasReadings) P() (float64, bool) {
	return readings.P_.Float64, readings.P_.Valid
}

// Vp возвращает рабочий объём, м3
func (readings *gasReadings) Vp() (float64, bool) {
	return readings.Vp_.Float64, readings.Vp_.Valid
}

// VpTotal возвращает накопленное значение рабочего объёма с момента сброса, м3
func (readings *gasReadings) VpTotal() (float64, bool) {
	return readings.VpTotal_.Float64, readings.VpTotal_.Valid
}

// Vc возвращает стандартный объём, м3
func (readings *gasReadings) Vc() (float64, bool) {
	return readings.Vc_.Float64, readings.Vc_.Valid
}

// VcTotal возвращает накопленное значение стандартного объёма с момента сброса, м3
func (readings *gasReadings) VcTotal() (float64, bool) {
	return readings.VcTotal_.Float64, readings.VcTotal_.Valid
}
