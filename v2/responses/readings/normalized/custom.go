package normalized

import (
	"time"

	"github.com/guregu/null"
)

var _ Readings = (*custom)(nil)

type custom struct {
	// Date_ дата архива, приборная
	Date_ int64 `json:"date"`

	// DateOnEndOfArchive_ дата архива на конец часа
	DateOnEndOfArchive_ int64 `json:"dateOnEndOfArchive"`

	// DateWithTimeBias_ дата архива на конец часа с учётом времени на приборе и часового пояса
	DateWithTimeBias_ int64 `json:"dateWithTimeBias"`

	// TGMax_ время, в течение которого фактический массовый расход был выше максимального нормированного значения для
	// средства измерения, ч
	TGMax_ null.Float `json:"TGmax,omitempty"`

	// TGMin_ время, в течение которого фактический массовый расход был меньше допустимого минимального нормированного
	// значения для средства измерения, ч
	TGMin_ null.Float `json:"TGmin,omitempty"`

	// TFault_ время функционального отказа, ч
	TFault_ null.Float `json:"TFault,omitempty"`

	// Toff_ время отключения питания, ч
	Toff_ null.Float `json:"Toff,omitempty"`

	// TOtherNS_ суммарное время других нештатных ситуаций, ч
	TOtherNS_ null.Float `json:"TOtherNS,omitempty"`

	// QntHIP_ время нормальной работы, ч
	QntHIP_ null.Float `json:"QntHIP,omitempty"`

	// QntHIPTotal_ накопленное значение времени нормальной работы с момента сброса, ч
	QntHIPTotal_ null.Float `json:"QntHIP_Total,omitempty"`

	// QntP_ время отсутствия счёта, ч
	QntP_ null.Float `json:"QntP,omitempty"`

	// QntPTotal_ накопленное значение времени отсутствия счёта с момента сброса, ч
	QntPTotal_ null.Float `json:"QntP_Total,omitempty"`

	// NS_ наличие нештатной ситуации
	NS_ bool `json:"ns"`

	// Empty_ признак пустого архива
	Empty_ bool `json:"empty"`
}

// Date возвращает дату архива, приборная
func (readings *custom) Date() time.Time {
	return time.Unix(readings.Date_, 0).UTC()
}

// DateOnEndOfArchive возвращает дату архива на конец часа
func (readings *custom) DateOnEndOfArchive() time.Time {
	return time.Unix(readings.DateOnEndOfArchive_, 0).UTC()
}

// DateWithTimeBias возвращает дату архива на конец часа с учётом времени на приборе и часового пояса
func (readings *custom) DateWithTimeBias() time.Time {
	return time.Unix(readings.DateWithTimeBias_, 0).UTC()
}

// TGMax возвращает время, в течение которого фактический массовый расход был выше максимального нормированного
// значения для средства измерения, ч
func (readings *custom) TGMax() (float64, bool) {
	return readings.TGMax_.Float64, readings.TGMax_.Valid
}

// TGMin возвращает время, в течение которого фактический массовый расход был меньше допустимого минимального
// нормированного значения для средства измерения, ч
func (readings *custom) TGMin() (float64, bool) {
	return readings.TGMin_.Float64, readings.TGMin_.Valid
}

// TFault возвращает время функционального отказа, ч
func (readings *custom) TFault() (float64, bool) {
	return readings.TFault_.Float64, readings.TFault_.Valid
}

// Toff возвращает время отключения питания, ч
func (readings *custom) Toff() (float64, bool) {
	return readings.Toff_.Float64, readings.Toff_.Valid
}

// TOtherNS возвращает суммарное время других нештатных ситуаций, ч
func (readings *custom) TOtherNS() (float64, bool) {
	return readings.TOtherNS_.Float64, readings.TOtherNS_.Valid
}

// QntHIP возвращает время нормальной работы, ч
func (readings *custom) QntHIP() (float64, bool) {
	return readings.QntHIP_.Float64, readings.QntHIP_.Valid
}

// QntHIPTotal возвращает накопленное значение времени нормальной работы с момента сброса, ч
func (readings *custom) QntHIPTotal() (float64, bool) {
	return readings.QntHIPTotal_.Float64, readings.QntHIPTotal_.Valid
}

// QntP возвращает время отсутствия счёта, ч
func (readings *custom) QntP() (float64, bool) {
	return readings.QntP_.Float64, readings.QntP_.Valid
}

// QntPTotal возвращает накопленное значение времени отсутствия счёта с момента сброса, ч
func (readings *custom) QntPTotal() (float64, bool) {
	return readings.QntPTotal_.Float64, readings.QntPTotal_.Valid
}

// NS возвращает признак наличия нештатной ситуации
func (readings *custom) NS() bool {
	return readings.NS_
}

// Empty возвращает признак пустого архива
func (readings *custom) Empty() bool {
	return readings.Empty_
}
