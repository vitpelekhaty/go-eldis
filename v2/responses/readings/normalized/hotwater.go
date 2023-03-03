package normalized

import (
	"github.com/guregu/null"
)

var _ HotWaterReadings = (*hotWaterReadings)(nil)

type hotWaterReadings struct {
	custom

	// Q1_ количество тепловой энергии на подающем трубопроводе, Гкал
	Q1_ null.Float `json:"Q1,omitempty"`

	// Q1Total_ накопленное значение потребления тепловой энергии на подающем трубопроводе с момента сброса, Гкал
	Q1Total_ null.Float `json:"Q1_Total,omitempty"`

	// Q2_ количество тепловой энергии на обратном трубопроводе, Гкал
	Q2_ null.Float `json:"Q2,omitempty"`

	// Q2Total_ накопленное значение потребления тепловой энергии на обратном трубопроводе с момента сброса, Гкал
	Q2Total_ null.Float `json:"Q2_Total,omitempty"`

	// T1_ температура на подающем трубопроводе, °C
	T1_ null.Float `json:"t1,omitempty"`

	// T2_ температура на обратном трубопроводе, °C
	T2_ null.Float `json:"t2,omitempty"`

	// V1_ объём на подающем трубопроводе, м3
	V1_ null.Float `json:"V1,omitempty"`

	// V1Total_ накопленное значение объёма на подающем трубопроводе с момента сброса, м3
	V1Total_ null.Float `json:"V1_Total,omitempty"`

	// V2_ объём на обратном трубопроводе, м3
	V2_ null.Float `json:"V2,omitempty"`

	// V2Total_ накопленное значение объёма на обратном трубопроводе с момента сброса, м3
	V2Total_ null.Float `json:"V2_Total,omitempty"`

	// M1_ масса на подающем трубопроводе, т
	M1_ null.Float `json:"M1,omitempty"`

	// M1Total_ накопленное значение массы на подающем трубопроводе с момента сброса, т
	M1Total_ null.Float `json:"M1_Total,omitempty"`

	// M2_ масса на обратном трубопроводе, т
	M2_ null.Float `json:"M2,omitempty"`

	// M2Total_ накопленное значение массы на обратном трубопроводе с момента сброса, т
	M2Total_ null.Float `json:"M2_Total,omitempty"`

	// P1_ давление на подающем трубопроводе, МПа
	P1_ null.Float `json:"P1,omitempty"`

	// P2_ давление на обратном трубопроводе, МПа
	P2_ null.Float `json:"P2,omitempty"`

	// Q_ количество тепловой энергии по всей системе, Гкал
	Q_ null.Float `json:"Q,omitempty"`

	// QTotal_ накопленное значение потребления тепловой энергии по всей системе с момента сброса, Гкал
	QTotal_ null.Float `json:"Q_Total,omitempty"`

	// DT_ разность температур, °C
	DT_ null.Float `json:"dt,omitempty"`

	// DV_ разность объёмов, м3
	DV_ null.Float `json:"dV,omitempty"`

	// DM_ разность масс, т
	DM_ null.Float `json:"dM,omitempty"`

	// Tcw_ температура холодной воды, °C
	Tcw_ null.Float `json:"tcw,omitempty"`

	// Pcw_ давление холодной воды, МПа
	Pcw_ null.Float `json:"Pcw,omitempty"`

	// Tdt_ время, в течение которого разность температур на подающем и обратном трубопроводах была меньше допустимой
	// нормированной разности температур для теплосчетчика, ч
	Tdt_ null.Float `json:"Tdt,omitempty"`
}

// Q1 возвращает количество тепловой энергии на подающем трубопроводе, Гкал
func (readings *hotWaterReadings) Q1() (float64, bool) {
	return readings.Q1_.Float64, readings.Q1_.Valid
}

// Q1Total возвращает накопленное значение потребления тепловой энергии на подающем трубопроводе с момента
// сброса, Гкал
func (readings *hotWaterReadings) Q1Total() (float64, bool) {
	return readings.Q1Total_.Float64, readings.Q1Total_.Valid
}

// Q2 возвращает количество тепловой энергии на обратном трубопроводе, Гкал
func (readings *hotWaterReadings) Q2() (float64, bool) {
	return readings.Q2_.Float64, readings.Q2_.Valid
}

// Q2Total возвращает накопленное значение потребления тепловой энергии на обратном трубопроводе с момента
// сброса, Гкал
func (readings *hotWaterReadings) Q2Total() (float64, bool) {
	return readings.Q2Total_.Float64, readings.Q2Total_.Valid
}

// T1 возвращает температуру на подающем трубопроводе, °C
func (readings *hotWaterReadings) T1() (float64, bool) {
	return readings.T1_.Float64, readings.T1_.Valid
}

// T2 возвращает температуру на обратном трубопроводе, °C
func (readings *hotWaterReadings) T2() (float64, bool) {
	return readings.T2_.Float64, readings.T2_.Valid
}

// V1 возвращает объём на подающем трубопроводе, м3
func (readings *hotWaterReadings) V1() (float64, bool) {
	return readings.V1_.Float64, readings.V1_.Valid
}

// V1Total возвращает накопленное значение объёма на подающем трубопроводе с момента сброса, м3
func (readings *hotWaterReadings) V1Total() (float64, bool) {
	return readings.V1Total_.Float64, readings.V1Total_.Valid
}

// V2 возвращает объём на обратном трубопроводе, м3
func (readings *hotWaterReadings) V2() (float64, bool) {
	return readings.V2_.Float64, readings.V2_.Valid
}

// V2Total возвращает накопленное значение объёма на обратном трубопроводе с момента сброса, м3
func (readings *hotWaterReadings) V2Total() (float64, bool) {
	return readings.V2Total_.Float64, readings.V2Total_.Valid
}

// M1 возвращает массу на подающем трубопроводе, т
func (readings *hotWaterReadings) M1() (float64, bool) {
	return readings.M1_.Float64, readings.M1_.Valid
}

// M1Total возвращает накопленное значение массы на подающем трубопроводе с момента сброса, т
func (readings *hotWaterReadings) M1Total() (float64, bool) {
	return readings.M1Total_.Float64, readings.M1Total_.Valid
}

// M2 возвращает массу на обратном трубопроводе, т
func (readings *hotWaterReadings) M2() (float64, bool) {
	return readings.M2_.Float64, readings.M2_.Valid
}

// M2Total возвращает накопленное значение массы на обратном трубопроводе с момента сброса, т
func (readings *hotWaterReadings) M2Total() (float64, bool) {
	return readings.M2Total_.Float64, readings.M2Total_.Valid
}

// P1 возвращает давление на подающем трубопроводе, МПа
func (readings *hotWaterReadings) P1() (float64, bool) {
	return readings.P1_.Float64, readings.P1_.Valid
}

// P2 возвращает давление на обратном трубопроводе, МПа
func (readings *hotWaterReadings) P2() (float64, bool) {
	return readings.P2_.Float64, readings.P2_.Valid
}

// Q возвращает количество тепловой энергии по всей системе, Гкал
func (readings *hotWaterReadings) Q() (float64, bool) {
	return readings.Q_.Float64, readings.Q_.Valid
}

// QTotal возвращает накопленное значение потребления тепловой энергии по всей системе с момента сброса, Гкал
func (readings *hotWaterReadings) QTotal() (float64, bool) {
	return readings.QTotal_.Float64, readings.QTotal_.Valid
}

// DT возвращает разность температур, °C
func (readings *hotWaterReadings) DT() (float64, bool) {
	return readings.DT_.Float64, readings.DT_.Valid
}

// DV возвращает разность объёмов, м3
func (readings *hotWaterReadings) DV() (float64, bool) {
	return readings.DV_.Float64, readings.DV_.Valid
}

// DM возвращает разность масс, т
func (readings *hotWaterReadings) DM() (float64, bool) {
	return readings.DM_.Float64, readings.DM_.Valid
}

// Tcw возвращает температуру холодной воды, °C
func (readings *hotWaterReadings) Tcw() (float64, bool) {
	return readings.Tcw_.Float64, readings.Tcw_.Valid
}

// Pcw возвращает давление холодной воды, МПа
func (readings *hotWaterReadings) Pcw() (float64, bool) {
	return readings.Pcw_.Float64, readings.Pcw_.Valid
}

// Tdt возвращает время, в течение которого разность температур на подающем и обратном трубопроводах была меньше
// допустимой нормированной разности температур для теплосчетчика, ч
func (readings *hotWaterReadings) Tdt() (float64, bool) {
	return readings.Tdt_.Float64, readings.Tdt_.Valid
}
