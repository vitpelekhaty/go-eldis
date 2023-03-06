package normalized

import (
	"time"

	"github.com/guregu/null"
)

var _ ElectricityReadings = (*electricityReadings)(nil)
var _ ElectricityCurrentReadings = (*electricityCurrentReadings)(nil)

type electricityReadings struct {
	// Date_ дата архива, приборная
	Date_ int64 `json:"date"`

	// DateOnEndOfArchive_ дата архива на конец часа
	DateOnEndOfArchive_ int64 `json:"dateOnEndOfArchive"`

	// DateWithTimeBias_ дата архива на конец часа с учётом времени на приборе и часового пояса
	DateWithTimeBias_ int64 `json:"dateWithTimeBias"`

	// Ap1_ активная энергия прямого направления по тарифу 1, кВт*ч
	Ap1_ null.Float `json:"Ap1,omitempty"`

	// Ap1Total_ накопленное значение активной энергии прямого направления по тарифу 1 с момента сброса, кВт*ч
	Ap1Total_ null.Float `json:"Ap1_Total,omitempty"`

	// Ap2_ активная энергия прямого направления по тарифу 2, кВт*ч
	Ap2_ null.Float `json:"Ap2,omitempty"`

	// Ap2Total_ накопленное значение активной энергии прямого направления по тарифу 2 с момента сброса, кВт*ч
	Ap2Total_ null.Float `json:"Ap2_Total,omitempty"`

	// Ap3_ активная энергия прямого направления по тарифу 3, кВт*ч
	Ap3_ null.Float `json:"Ap3,omitempty"`

	// Ap3Total_ накопленное значение активной энергии прямого направления по тарифу 3 с момента сброса, кВт*ч
	Ap3Total_ null.Float `json:"Ap3_Total,omitempty"`

	// Ap4_ активная энергия прямого направления по тарифу 4, кВт*ч
	Ap4_ null.Float `json:"Ap4,omitempty"`

	// Ap4Total_ накопленное значение активной энергии прямого направления по тарифу 4 с момента сброса, кВт*ч
	Ap4Total_ null.Float `json:"Ap4_Total,omitempty"`

	// Am1_ активная энергия обратного направления по тарифу 1, кВт*ч
	Am1_ null.Float `json:"Am1,omitempty"`

	// Am1Total_ накопленное значение активной энергии обратного направления по тарифу 1 с момента сброса, кВт*ч
	Am1Total_ null.Float `json:"Am1_Total,omitempty"`

	// Am2_ активная энергия обратного направления по тарифу 2, кВт*ч
	Am2_ null.Float `json:"Am2,omitempty"`

	// Am2Total_ накопленное значение активной энергии обратного направления по тарифу 2 с момента сброса, кВт*ч
	Am2Total_ null.Float `json:"Am2_Total,omitempty"`

	// Am3_ активная энергия обратного направления по тарифу 3, кВт*ч
	Am3_ null.Float `json:"Am3,omitempty"`

	// Am3Total_ накопленное значение активной энергии обратного направления по тарифу 3 с момента сброса, кВт*ч
	Am3Total_ null.Float `json:"Am3_Total,omitempty"`

	// Am4_ активная энергия обратного направления по тарифу 4, кВт*ч
	Am4_ null.Float `json:"Am4,omitempty"`

	// Am4Total_ накопленное значение активной энергии обратного направления по тарифу 4 с момента сброса, кВт*ч
	Am4Total_ null.Float `json:"Am4_Total,omitempty"`

	// Rp1_ реактивная энергия прямого направления по тарифу 1, кВар*ч
	Rp1_ null.Float `json:"Rp1,omitempty"`

	// Rp1Total_ накопленное значение реактивной энергии прямого направления по тарифу 1 с момента сброса, кВар*ч
	Rp1Total_ null.Float `json:"Rp1_Total,omitempty"`

	// Rp2_ реактивная энергия прямого направления по тарифу 2, кВар*ч
	Rp2_ null.Float `json:"Rp2,omitempty"`

	// Rp2Total_ накопленное значение реактивной энергии прямого направления по тарифу 2 с момента сброса, кВар*ч
	Rp2Total_ null.Float `json:"Rp2_Total,omitempty"`

	// Rp3_ реактивная энергия прямого направления по тарифу 3, кВар*ч
	Rp3_ null.Float `json:"Rp3,omitempty"`

	// Rp3Total_ накопленное значение реактивной энергии прямого направления по тарифу 3 с момента сброса, кВар*ч
	Rp3Total_ null.Float `json:"Rp3_Total,omitempty"`

	// Rp4_ реактивная энергия прямого направления по тарифу 4, кВар*ч
	Rp4_ null.Float `json:"Rp4,omitempty"`

	// Rp4Total_ накопленное значение реактивной энергии прямого направления по тарифу 4 с момента сброса, кВар*ч
	Rp4Total_ null.Float `json:"Rp4_Total,omitempty"`

	// Rm1_ реактивная энергия обратного направления по тарифу 1, кВар*ч
	Rm1_ null.Float `json:"Rm1,omitempty"`

	// Rm1Total_ накопленное значение реактивной энергии обратного направления по тарифу 1 с момента сброса, кВар*ч
	Rm1Total_ null.Float `json:"Rm1_Total,omitempty"`

	// Rm2_ реактивная энергия обратного направления по тарифу 2, кВар*ч
	Rm2_ null.Float `json:"Rm2,omitempty"`

	// Rm2Total_ накопленное значение реактивной энергии обратного направления по тарифу 2 с момента сброса, кВар*ч
	Rm2Total_ null.Float `json:"Rm2_Total,omitempty"`

	// Rm3_ реактивная энергия обратного направления по тарифу 3, кВар*ч
	Rm3_ null.Float `json:"Rm3,omitempty"`

	// Rm3Total_ накопленное значение реактивной энергии обратного направления по тарифу 3 с момента сброса, кВар*ч
	Rm3Total_ null.Float `json:"Rm3_Total,omitempty"`

	// Rm4_ реактивная энергия обратного направления по тарифу 4, кВар*ч
	Rm4_ null.Float `json:"Rm4,omitempty"`

	// Rm4Total_ накопленное значение реактивной энергии обратного направления по тарифу 4 с момента сброса, кВар*ч
	Rm4Total_ null.Float `json:"Rm4_Total,omitempty"`

	// Ap_ сумма тарифов активной энергии прямого направления, кВт*ч
	Ap_ null.Float `json:"Ap,omitempty"`

	// ApTotal_ сумма тарифов накопленных значений активной энергии прямого направления с момента сброса, кВт*ч
	ApTotal_ null.Float `json:"Ap_Total,omitempty"`

	// Am_ сумма тарифов активной энергии обратного направления, кВар*ч
	Am_ null.Float `json:"Am,omitempty"`

	// AmTotal_ сумма тарифов накопленных значений активной энергии обратного направления с момента сброса, кВт*ч
	AmTotal_ null.Float `json:"Am_Total,omitempty"`

	// Rp_ сумма тарифов реактивной энергии прямого направления, кВар*ч
	Rp_ null.Float `json:"Rp,omitempty"`

	// RpTotal_ сумма тарифов накопленных значений реактивной энергии прямого направления с момента сброса, кВар*ч
	RpTotal_ null.Float `json:"Rp_Total,omitempty"`

	// Rm_ сумма тарифов реактивной энергии обратного направления, кВар*ч
	Rm_ null.Float `json:"Rm,omitempty"`

	// RmTotal_ сумма тарифов накопленных значений реактивной энергии обратного направления с момента сброса, кВар*ч
	RmTotal_ null.Float `json:"Rm_Total,omitempty"`

	// Pp_ активная мощность прямого направления, кВт
	Pp_ null.Float `json:"Pp,omitempty"`

	// Pm_ активная мощность обратного направления, кВт
	Pm_ null.Float `json:"Pm,omitempty"`

	// Qp_ реактивная мощность прямого направления, кВар
	Qp_ null.Float `json:"Qp,omitempty"`

	// Qm_ реактивная мощность обратного направления, кВар
	Qm_ null.Float `json:"Qm,omitempty"`

	// NS_ наличие нештатной ситуации
	NS_ bool `json:"ns"`
}

type electricityCurrentReadings struct {
	// Date_ дата архива, приборная
	Date_ int64 `json:"date"`

	// DateOnEndOfArchive_ дата архива на конец часа
	DateOnEndOfArchive_ int64 `json:"dateOnEndOfArchive"`

	// DateWithTimeBias_ дата архива на конец часа с учётом времени на приборе и часового пояса
	DateWithTimeBias_ int64 `json:"dateWithTimeBias"`

	// S1_ полная мощность по фазе 1, кВА
	S1_ null.Float `json:"S1,omitempty"`

	// S2_ полная мощность по фазе 2, кВА
	S2_ null.Float `json:"S2,omitempty"`

	// S3_ полная мощность по фазе 3, кВА
	S3_ null.Float `json:"S3,omitempty"`

	// P1_ активная мощность по фазе 1, кВт
	P1_ null.Float `json:"P1,omitempty"`

	// P2_ активная мощность по фазе 2, кВт
	P2_ null.Float `json:"P2,omitempty"`

	// P3_ активная мощность по фазе 3, кВт
	P3_ null.Float `json:"P3,omitempty"`

	// Q1_ реактивная мощность по фазе 1, кВар
	Q1_ null.Float `json:"Q1,omitempty"`

	// Q2_ реактивная мощность по фазе 2, кВар
	Q2_ null.Float `json:"Q2,omitempty"`

	// Q3_ реактивная мощность по фазе 3, кВар
	Q3_ null.Float `json:"Q3,omitempty"`

	// U1_ напряжение по фазе 1, В
	U1_ null.Float `json:"U1,omitempty"`

	// U2_ напряжение по фазе 2, В
	U2_ null.Float `json:"U2,omitempty"`

	// U3_ напряжение по фазе 3, В
	U3_ null.Float `json:"U3,omitempty"`

	// I1_ сила тока по фазе 1, А
	I1_ null.Float `json:"I1,omitempty"`

	// I2_ сила тока по фазе 2, А
	I2_ null.Float `json:"I2,omitempty"`

	// I3_ сила тока по фазе 3, А
	I3_ null.Float `json:"I3,omitempty"`

	// S_ сумма фаз полной мощности, кВА
	S_ null.Float `json:"S,omitempty"`

	// P_ сумма фаз активной мощности, кВт
	P_ null.Float `json:"P,omitempty"`

	// Q_ сумма фаз реактивной мощности, кВар
	Q_ null.Float `json:"Q,omitempty"`

	// U_ сумма фаз напряжения, В
	U_ null.Float `json:"U,omitempty"`

	// I_ сумма фаз силы тока, А
	I_ null.Float `json:"I,omitempty"`

	// F_ частота тока, Гц
	F_ null.Float `json:"F,omitempty"`

	// CosPhi1_ коэффициент мощности по фазе 1
	CosPhi1_ null.Float `json:"CosPhi1,omitempty"`

	// CosPhi2_ коэффициент мощности по фазе 2
	CosPhi2_ null.Float `json:"CosPhi2,omitempty"`

	// CosPhi3_ коэффициент мощности по фазе 3
	CosPhi3_ null.Float `json:"CosPhi3,omitempty"`

	// PhiU12_ угол между фазными напряжениями фазы 1 и 2
	PhiU12_ null.Float `json:"PhiU12,omitempty"`

	// PhiU13_ угол между фазными напряжениями фазы 1 и 3
	PhiU13_ null.Float `json:"PhiU13,omitempty"`

	// PhiU23_ угол между фазными напряжениями фазы 2 и 3
	PhiU23_ null.Float `json:"PhiU23,omitempty"`
}

// Date возвращает дату архива, приборная
func (readings *electricityReadings) Date() time.Time {
	return time.Unix(readings.Date_, 0).UTC()
}

// DateOnEndOfArchive возвращает дату архива на конец часа
func (readings *electricityReadings) DateOnEndOfArchive() time.Time {
	return time.Unix(readings.DateOnEndOfArchive_, 0).UTC()
}

// DateWithTimeBias возвращает дату архива на конец часа с учётом времени на приборе и часового пояса
func (readings *electricityReadings) DateWithTimeBias() time.Time {
	return time.Unix(readings.DateWithTimeBias_, 0).UTC()
}

// Ap1 возвращает активную энергию прямого направления по тарифу 1, кВт*ч
func (readings *electricityReadings) Ap1() (float64, bool) {
	return readings.Ap1_.Float64, readings.Ap1_.Valid
}

// Ap1Total возвращает накопленное значение активной энергии прямого направления по тарифу 1 с момента
// сброса, кВт*ч
func (readings *electricityReadings) Ap1Total() (float64, bool) {
	return readings.Ap1Total_.Float64, readings.Ap1Total_.Valid
}

// Ap2 возвращает активную энергию прямого направления по тарифу 2, кВт*ч
func (readings *electricityReadings) Ap2() (float64, bool) {
	return readings.Ap2_.Float64, readings.Ap2_.Valid
}

// Ap2Total возвращает накопленное значение активной энергии прямого направления по тарифу 2 с момента
// сброса, кВт*ч
func (readings *electricityReadings) Ap2Total() (float64, bool) {
	return readings.Ap2Total_.Float64, readings.Ap2Total_.Valid
}

// Ap3 возвращает активную энергию прямого направления по тарифу 3, кВт*ч
func (readings *electricityReadings) Ap3() (float64, bool) {
	return readings.Ap3_.Float64, readings.Ap3_.Valid
}

// Ap3Total возвращает накопленное значение активной энергии прямого направления по тарифу 3 с момента
// сброса, кВт*ч
func (readings *electricityReadings) Ap3Total() (float64, bool) {
	return readings.Ap3Total_.Float64, readings.Ap3Total_.Valid
}

// Ap4 возвращает активную энергию прямого направления по тарифу 4, кВт*ч
func (readings *electricityReadings) Ap4() (float64, bool) {
	return readings.Ap4_.Float64, readings.Ap4_.Valid
}

// Ap4Total возвращает накопленное значение активной энергии прямого направления по тарифу 4 с момента
// сброса, кВт*ч
func (readings *electricityReadings) Ap4Total() (float64, bool) {
	return readings.Ap4Total_.Float64, readings.Ap4Total_.Valid
}

// Am1 возвращает активную энергию обратного направления по тарифу 1, кВт*ч
func (readings *electricityReadings) Am1() (float64, bool) {
	return readings.Am1_.Float64, readings.Am1_.Valid
}

// Am1Total возвращает накопленное значение активной энергии обратного направления по тарифу 1 с момента
// сброса, кВт*ч
func (readings *electricityReadings) Am1Total() (float64, bool) {
	return readings.Am1Total_.Float64, readings.Am1Total_.Valid
}

// Am2 возвращает активную энергию обратного направления по тарифу 2, кВт*ч
func (readings *electricityReadings) Am2() (float64, bool) {
	return readings.Am2_.Float64, readings.Am2_.Valid
}

// Am2Total возвращает накопленное значение активной энергии обратного направления по тарифу 2 с момента сброса, кВт*ч
func (readings *electricityReadings) Am2Total() (float64, bool) {
	return readings.Am2Total_.Float64, readings.Am2Total_.Valid
}

// Am3 возвращает активную энергию обратного направления по тарифу 3, кВт*ч
func (readings *electricityReadings) Am3() (float64, bool) {
	return readings.Am3_.Float64, readings.Am3_.Valid
}

// Am3Total возвращает накопленное значение активной энергии обратного направления по тарифу 3 с момента
// сброса, кВт*ч
func (readings *electricityReadings) Am3Total() (float64, bool) {
	return readings.Am3Total_.Float64, readings.Am3Total_.Valid
}

// Am4 возвращает активную энергию обратного направления по тарифу 4, кВт*ч
func (readings *electricityReadings) Am4() (float64, bool) {
	return readings.Am4_.Float64, readings.Am4_.Valid
}

// Am4Total возвращает накопленное значение активной энергии обратного направления по тарифу 4 с момента
// сброса, кВт*ч
func (readings *electricityReadings) Am4Total() (float64, bool) {
	return readings.Am4Total_.Float64, readings.Am4Total_.Valid
}

// Rp1 возвращает реактивную энергию прямого направления по тарифу 1, кВар*ч
func (readings *electricityReadings) Rp1() (float64, bool) {
	return readings.Rp1_.Float64, readings.Rp1_.Valid
}

// Rp1Total возвращает накопленное значение реактивной энергии прямого направления по тарифу 1 с момента
// сброса, кВар*ч
func (readings *electricityReadings) Rp1Total() (float64, bool) {
	return readings.Rp1Total_.Float64, readings.Rp1Total_.Valid
}

// Rp2 возвращает реактивную энергию прямого направления по тарифу 2, кВар*ч
func (readings *electricityReadings) Rp2() (float64, bool) {
	return readings.Rp2_.Float64, readings.Rp2_.Valid
}

// Rp2Total возвращает накопленное значение реактивной энергии прямого направления по тарифу 2 с момента сброса, кВар*ч
func (readings *electricityReadings) Rp2Total() (float64, bool) {
	return readings.Rp2Total_.Float64, readings.Rp2Total_.Valid
}

// Rp3 возвращает реактивную энергию прямого направления по тарифу 3, кВар*ч
func (readings *electricityReadings) Rp3() (float64, bool) {
	return readings.Rp3_.Float64, readings.Rp3_.Valid
}

// Rp3Total возвращает накопленное значение реактивной энергии прямого направления по тарифу 3 с момента
// сброса, кВар*ч
func (readings *electricityReadings) Rp3Total() (float64, bool) {
	return readings.Rp3Total_.Float64, readings.Rp3Total_.Valid
}

// Rp4 возвращает реактивную энергию прямого направления по тарифу 4, кВар*ч
func (readings *electricityReadings) Rp4() (float64, bool) {
	return readings.Rp4_.Float64, readings.Rp4_.Valid
}

// Rp4Total возвращает накопленное значение реактивной энергии прямого направления по тарифу 4 с момента
// сброса, кВар*ч
func (readings *electricityReadings) Rp4Total() (float64, bool) {
	return readings.Rp4Total_.Float64, readings.Rp4Total_.Valid
}

// Rm1 возвращает реактивную энергию обратного направления по тарифу 1, кВар*ч
func (readings *electricityReadings) Rm1() (float64, bool) {
	return readings.Rm1_.Float64, readings.Rm1_.Valid
}

// Rm1Total возвращает накопленное значение реактивной энергии обратного направления по тарифу 1 с момента
// сброса, кВар*ч
func (readings *electricityReadings) Rm1Total() (float64, bool) {
	return readings.Rm1Total_.Float64, readings.Rm1Total_.Valid
}

// Rm2 возвращает реактивную энергию обратного направления по тарифу 2, кВар*ч
func (readings *electricityReadings) Rm2() (float64, bool) {
	return readings.Rm2_.Float64, readings.Rm2_.Valid
}

// Rm2Total возвращает накопленное значение реактивной энергии обратного направления по тарифу 2 с момента
// сброса, кВар*ч
func (readings *electricityReadings) Rm2Total() (float64, bool) {
	return readings.Rm2Total_.Float64, readings.Rm2Total_.Valid
}

// Rm3 возвращает реактивную энергию обратного направления по тарифу 3, кВар*ч
func (readings *electricityReadings) Rm3() (float64, bool) {
	return readings.Rm3_.Float64, readings.Rm3_.Valid
}

// Rm3Total возвращает накопленное значение реактивной энергии обратного направления по тарифу 3 с момента
// сброса, кВар*ч
func (readings *electricityReadings) Rm3Total() (float64, bool) {
	return readings.Rm3Total_.Float64, readings.Rm3Total_.Valid
}

// Rm4 возвращает реактивную энергию обратного направления по тарифу 4, кВар*ч
func (readings *electricityReadings) Rm4() (float64, bool) {
	return readings.Rm4_.Float64, readings.Rm4_.Valid
}

// Rm4Total возвращает накопленное значение реактивной энергии обратного направления по тарифу 4 с момента
// сброса, кВар*ч
func (readings *electricityReadings) Rm4Total() (float64, bool) {
	return readings.Rm4Total_.Float64, readings.Rm4Total_.Valid
}

// Ap возвращает сумму тарифов активной энергии прямого направления, кВт*ч
func (readings *electricityReadings) Ap() (float64, bool) {
	return readings.Ap_.Float64, readings.Ap_.Valid
}

// ApTotal возвращает сумму тарифов накопленных значений активной энергии прямого направления с момента
// сброса, кВт*ч
func (readings *electricityReadings) ApTotal() (float64, bool) {
	return readings.ApTotal_.Float64, readings.ApTotal_.Valid
}

// Am возвращает сумму тарифов активной энергии обратного направления, кВар*ч
func (readings *electricityReadings) Am() (float64, bool) {
	return readings.Am_.Float64, readings.Am_.Valid
}

// AmTotal возвращает сумму тарифов накопленных значений активной энергии обратного направления с момента
// сброса, кВт*ч
func (readings *electricityReadings) AmTotal() (float64, bool) {
	return readings.AmTotal_.Float64, readings.AmTotal_.Valid
}

// Rp возвращает сумму тарифов реактивной энергии прямого направления, кВар*ч
func (readings *electricityReadings) Rp() (float64, bool) {
	return readings.Rp_.Float64, readings.Rp_.Valid
}

// RpTotal возвращает сумму тарифов накопленных значений реактивной энергии прямого направления с момента
// сброса, кВар*ч
func (readings *electricityReadings) RpTotal() (float64, bool) {
	return readings.RpTotal_.Float64, readings.RpTotal_.Valid
}

// Rm возвращает сумму тарифов реактивной энергии обратного направления, кВар*ч
func (readings *electricityReadings) Rm() (float64, bool) {
	return readings.Rm_.Float64, readings.Rm_.Valid
}

// RmTotal возвращает сумму тарифов накопленных значений реактивной энергии обратного направления с момента
// сброса, кВар*ч
func (readings *electricityReadings) RmTotal() (float64, bool) {
	return readings.RmTotal_.Float64, readings.RmTotal_.Valid
}

// Pp возвращает активную мощность прямого направления, кВт
func (readings *electricityReadings) Pp() (float64, bool) {
	return readings.Pp_.Float64, readings.Pp_.Valid
}

// Pm возвращает активную мощность обратного направления, кВт
func (readings *electricityReadings) Pm() (float64, bool) {
	return readings.Pm_.Float64, readings.Pm_.Valid
}

// Qp возвращает реактивную мощность прямого направления, кВар
func (readings *electricityReadings) Qp() (float64, bool) {
	return readings.Qp_.Float64, readings.Qp_.Valid
}

// Qm возвращает реактивную мощность обратного направления, кВар
func (readings *electricityReadings) Qm() (float64, bool) {
	return readings.Qm_.Float64, readings.Qm_.Valid
}

// NS возвращает признак наличия нештатной ситуации
func (readings *electricityReadings) NS() bool {
	return readings.NS_
}

// Date возвращает дату архива, приборная
func (readings *electricityCurrentReadings) Date() time.Time {
	return time.Unix(readings.Date_, 0).UTC()
}

// DateOnEndOfArchive возвращает дату архива на конец часа
func (readings *electricityCurrentReadings) DateOnEndOfArchive() time.Time {
	return time.Unix(readings.DateOnEndOfArchive_, 0).UTC()
}

// DateWithTimeBias возвращает дату архива на конец часа с учётом времени на приборе и часового пояса
func (readings *electricityCurrentReadings) DateWithTimeBias() time.Time {
	return time.Unix(readings.DateWithTimeBias_, 0).UTC()
}

// S1 возвращает полную мощность по фазе 1, кВА
func (readings *electricityCurrentReadings) S1() (float64, bool) {
	return readings.S1_.Float64, readings.S1_.Valid
}

// S2 возвращает полную мощность по фазе 2, кВА
func (readings *electricityCurrentReadings) S2() (float64, bool) {
	return readings.S2_.Float64, readings.S2_.Valid
}

// S3 возвращает полную мощность по фазе 3, кВА
func (readings *electricityCurrentReadings) S3() (float64, bool) {
	return readings.S3_.Float64, readings.S3_.Valid
}

// P1 возвращает активную мощность по фазе 1, кВт
func (readings *electricityCurrentReadings) P1() (float64, bool) {
	return readings.P1_.Float64, readings.P1_.Valid
}

// P2 возвращает активную мощность по фазе 2, кВт
func (readings *electricityCurrentReadings) P2() (float64, bool) {
	return readings.P2_.Float64, readings.P2_.Valid
}

// P3 возвращает активную мощность по фазе 3, кВт
func (readings *electricityCurrentReadings) P3() (float64, bool) {
	return readings.P3_.Float64, readings.P3_.Valid
}

// Q1 возвращает реактивную мощность по фазе 1, кВар
func (readings *electricityCurrentReadings) Q1() (float64, bool) {
	return readings.Q1_.Float64, readings.Q1_.Valid
}

// Q2 возвращает реактивную мощность по фазе 2, кВар
func (readings *electricityCurrentReadings) Q2() (float64, bool) {
	return readings.Q2_.Float64, readings.Q2_.Valid
}

// Q3 возвращает реактивную мощность по фазе 3, кВар
func (readings *electricityCurrentReadings) Q3() (float64, bool) {
	return readings.Q3_.Float64, readings.Q3_.Valid
}

// U1 возвращает напряжение по фазе 1, В
func (readings *electricityCurrentReadings) U1() (float64, bool) {
	return readings.U1_.Float64, readings.U1_.Valid
}

// U2 возвращает напряжение по фазе 2, В
func (readings *electricityCurrentReadings) U2() (float64, bool) {
	return readings.U2_.Float64, readings.U2_.Valid
}

// U3 возвращает напряжение по фазе 3, В
func (readings *electricityCurrentReadings) U3() (float64, bool) {
	return readings.U3_.Float64, readings.U3_.Valid
}

// I1 возвращает силу тока по фазе 1, А
func (readings *electricityCurrentReadings) I1() (float64, bool) {
	return readings.I1_.Float64, readings.I1_.Valid
}

// I2 возвращает силу тока по фазе 2, А
func (readings *electricityCurrentReadings) I2() (float64, bool) {
	return readings.I2_.Float64, readings.I2_.Valid
}

// I3 возвращает силу тока по фазе 3, А
func (readings *electricityCurrentReadings) I3() (float64, bool) {
	return readings.I3_.Float64, readings.I3_.Valid
}

// S возвращает сумму фаз полной мощности, кВА
func (readings *electricityCurrentReadings) S() (float64, bool) {
	return readings.S_.Float64, readings.S_.Valid
}

// P возвращает сумму фаз активной мощности, кВт
func (readings *electricityCurrentReadings) P() (float64, bool) {
	return readings.P_.Float64, readings.P_.Valid
}

// Q возвращает сумму фаз реактивной мощности, кВар
func (readings *electricityCurrentReadings) Q() (float64, bool) {
	return readings.Q_.Float64, readings.Q_.Valid
}

// U возвращает сумму фаз напряжения, В
func (readings *electricityCurrentReadings) U() (float64, bool) {
	return readings.U_.Float64, readings.U_.Valid
}

// I возвращает сумму фаз силы тока, А
func (readings *electricityCurrentReadings) I() (float64, bool) {
	return readings.I_.Float64, readings.I_.Valid
}

// F возвращает частоту тока, Гц
func (readings *electricityCurrentReadings) F() (float64, bool) {
	return readings.F_.Float64, readings.F_.Valid
}

// CosPhi1 возвращает коэффициент мощности по фазе 1
func (readings *electricityCurrentReadings) CosPhi1() (float64, bool) {
	return readings.CosPhi1_.Float64, readings.CosPhi1_.Valid
}

// CosPhi2 возвращает коэффициент мощности по фазе 2
func (readings *electricityCurrentReadings) CosPhi2() (float64, bool) {
	return readings.CosPhi2_.Float64, readings.CosPhi2_.Valid
}

// CosPhi3 возвращает коэффициент мощности по фазе 3
func (readings *electricityCurrentReadings) CosPhi3() (float64, bool) {
	return readings.CosPhi3_.Float64, readings.CosPhi3_.Valid
}

// PhiU12 возвращает угол между фазными напряжениями фазы 1 и 2
func (readings *electricityCurrentReadings) PhiU12() (float64, bool) {
	return readings.PhiU12_.Float64, readings.PhiU12_.Valid
}

// PhiU13 возвращает угол между фазными напряжениями фазы 1 и 3
func (readings *electricityCurrentReadings) PhiU13() (float64, bool) {
	return readings.PhiU13_.Float64, readings.PhiU13_.Valid
}

// PhiU23 возвращает угол между фазными напряжениями фазы 2 и 3
func (readings *electricityCurrentReadings) PhiU23() (float64, bool) {
	return readings.PhiU23_.Float64, readings.PhiU23_.Valid
}
