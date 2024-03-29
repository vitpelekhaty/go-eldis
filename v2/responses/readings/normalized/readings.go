package normalized

import (
	"time"
)

// Readings интерфейс нормализованных показаний точки учета
type Readings interface {
	// Date возвращает дату архива, приборная
	Date() time.Time

	// DateOnEndOfArchive возвращает дату архива на конец часа
	DateOnEndOfArchive() time.Time

	// DateWithTimeBias возвращает дату архива на конец часа с учётом времени на приборе и часового пояса
	DateWithTimeBias() time.Time

	// TGMax возвращает время, в течение которого фактический массовый расход был выше максимального нормированного
	// значения для средства измерения, ч
	TGMax() (float64, bool)

	// TGMin возвращает время, в течение которого фактический массовый расход был меньше допустимого минимального
	// нормированного значения для средства измерения, ч
	TGMin() (float64, bool)

	// TFault возвращает время функционального отказа, ч
	TFault() (float64, bool)

	// Toff возвращает время отключения питания, ч
	Toff() (float64, bool)

	// TOtherNS возвращает суммарное время других нештатных ситуаций, ч
	TOtherNS() (float64, bool)

	// QntHIP возвращает время нормальной работы, ч
	QntHIP() (float64, bool)

	// QntHIPTotal возвращает накопленное значение времени нормальной работы с момента сброса, ч
	QntHIPTotal() (float64, bool)

	// QntP возвращает время отсутствия счёта, ч
	QntP() (float64, bool)

	// QntPTotal возвращает накопленное значение времени отсутствия счёта с момента сброса, ч
	QntPTotal() (float64, bool)

	// NS возвращает признак наличия нештатной ситуации
	NS() bool

	// Empty возвращает признак пустого архива
	Empty() bool
}

// WasteWaterReadings нормализованные данные по сточным водам
type WasteWaterReadings interface {
	Readings

	// V возвращает объём, м3
	V() (float64, bool)

	// VTotal возвращает накопленное значение объёма с момента сброса, м3
	VTotal() (float64, bool)

	// M возвращает массу, т
	M() (float64, bool)

	// MTotal возвращает накопленное значение массы с момента сброса, т
	MTotal() (float64, bool)
}

// ColdWaterReadings нормализованные данные по холодному водоснабжению
type ColdWaterReadings interface {
	WasteWaterReadings

	// P возвращает давление, МПа
	P() (float64, bool)
}

// HotWaterReadings нормализованные данные по горячему водоснабжению
type HotWaterReadings interface {
	Readings

	// Q1 возвращает количество тепловой энергии на подающем трубопроводе, Гкал
	Q1() (float64, bool)

	// Q1Total возвращает накопленное значение потребления тепловой энергии на подающем трубопроводе с момента
	// сброса, Гкал
	Q1Total() (float64, bool)

	// Q2 возвращает количество тепловой энергии на обратном трубопроводе, Гкал
	Q2() (float64, bool)

	// Q2Total возвращает накопленное значение потребления тепловой энергии на обратном трубопроводе с момента
	// сброса, Гкал
	Q2Total() (float64, bool)

	// T1 возвращает температуру на подающем трубопроводе, °C
	T1() (float64, bool)

	// T2 возвращает температуру на обратном трубопроводе, °C
	T2() (float64, bool)

	// V1 возвращает объём на подающем трубопроводе, м3
	V1() (float64, bool)

	// V1Total возвращает накопленное значение объёма на подающем трубопроводе с момента сброса, м3
	V1Total() (float64, bool)

	// V2 возвращает объём на обратном трубопроводе, м3
	V2() (float64, bool)

	// V2Total возвращает накопленное значение объёма на обратном трубопроводе с момента сброса, м3
	V2Total() (float64, bool)

	// M1 возвращает массу на подающем трубопроводе, т
	M1() (float64, bool)

	// M1Total возвращает накопленное значение массы на подающем трубопроводе с момента сброса, т
	M1Total() (float64, bool)

	// M2 возвращает массу на обратном трубопроводе, т
	M2() (float64, bool)

	// M2Total возвращает накопленное значение массы на обратном трубопроводе с момента сброса, т
	M2Total() (float64, bool)

	// P1 возвращает давление на подающем трубопроводе, МПа
	P1() (float64, bool)

	// P2 возвращает давление на обратном трубопроводе, МПа
	P2() (float64, bool)

	// Q возвращает количество тепловой энергии по всей системе, Гкал
	Q() (float64, bool)

	// QTotal возвращает накопленное значение потребления тепловой энергии по всей системе с момента сброса, Гкал
	QTotal() (float64, bool)

	// DT возвращает разность температур, °C
	DT() (float64, bool)

	// DV возвращает разность объёмов, м3
	DV() (float64, bool)

	// DM возвращает разность масс, т
	DM() (float64, bool)

	// Tcw возвращает температуру холодной воды, °C
	Tcw() (float64, bool)

	// Pcw возвращает давление холодной воды, МПа
	Pcw() (float64, bool)

	// Tdt возвращает время, в течение которого разность температур на подающем и обратном трубопроводах была меньше
	// допустимой нормированной разности температур для теплосчетчика, ч
	Tdt() (float64, bool)
}

// HeatReadings нормализованные данные по теплу
type HeatReadings interface {
	HotWaterReadings

	// T3 возвращает температуру на подпитке, °C
	T3() (float64, bool)

	// V3 возвращает объём на подпитке, м3
	V3() (float64, bool)

	// V3Total возвращает накопленное значение объёма на подпитке с момента сброса, м3
	V3Total() (float64, bool)

	// M3 возвращает массу на подпитке, т
	M3() (float64, bool)

	// M3Total возвращает накопленное значение массы на подпитке с момента сброса, т
	M3Total() (float64, bool)

	// P3 возвращает давление на подпитке, МПа
	P3() (float64, bool)

	// Ta возвращает температуру наружного воздуха, °C
	Ta() (float64, bool)
}

// ElectricityReadings нормализованные данные по электричеству
type ElectricityReadings interface {
	// Date возвращает дату архива, приборная
	Date() time.Time

	// DateOnEndOfArchive возвращает дату архива на конец часа
	DateOnEndOfArchive() time.Time

	// DateWithTimeBias возвращает дату архива на конец часа с учётом времени на приборе и часового пояса
	DateWithTimeBias() time.Time

	// Ap1 возвращает активную энергию прямого направления по тарифу 1, кВт*ч
	Ap1() (float64, bool)

	// Ap1Total возвращает накопленное значение активной энергии прямого направления по тарифу 1 с момента
	// сброса, кВт*ч
	Ap1Total() (float64, bool)

	// Ap2 возвращает активную энергию прямого направления по тарифу 2, кВт*ч
	Ap2() (float64, bool)

	// Ap2Total возвращает накопленное значение активной энергии прямого направления по тарифу 2 с момента
	// сброса, кВт*ч
	Ap2Total() (float64, bool)

	// Ap3 возвращает активную энергию прямого направления по тарифу 3, кВт*ч
	Ap3() (float64, bool)

	// Ap3Total возвращает накопленное значение активной энергии прямого направления по тарифу 3 с момента
	// сброса, кВт*ч
	Ap3Total() (float64, bool)

	// Ap4 возвращает активную энергию прямого направления по тарифу 4, кВт*ч
	Ap4() (float64, bool)

	// Ap4Total возвращает накопленное значение активной энергии прямого направления по тарифу 4 с момента
	// сброса, кВт*ч
	Ap4Total() (float64, bool)

	// Am1 возвращает активную энергию обратного направления по тарифу 1, кВт*ч
	Am1() (float64, bool)

	// Am1Total возвращает накопленное значение активной энергии обратного направления по тарифу 1 с момента
	// сброса, кВт*ч
	Am1Total() (float64, bool)

	// Am2 возвращает активную энергию обратного направления по тарифу 2, кВт*ч
	Am2() (float64, bool)

	// Am2Total возвращает накопленное значение активной энергии обратного направления по тарифу 2 с момента сброса, кВт*ч
	Am2Total() (float64, bool)

	// Am3 возвращает активную энергию обратного направления по тарифу 3, кВт*ч
	Am3() (float64, bool)

	// Am3Total возвращает накопленное значение активной энергии обратного направления по тарифу 3 с момента
	// сброса, кВт*ч
	Am3Total() (float64, bool)

	// Am4 возвращает активную энергию обратного направления по тарифу 4, кВт*ч
	Am4() (float64, bool)

	// Am4Total возвращает накопленное значение активной энергии обратного направления по тарифу 4 с момента
	// сброса, кВт*ч
	Am4Total() (float64, bool)

	// Rp1 возвращает реактивную энергию прямого направления по тарифу 1, кВар*ч
	Rp1() (float64, bool)

	// Rp1Total возвращает накопленное значение реактивной энергии прямого направления по тарифу 1 с момента
	// сброса, кВар*ч
	Rp1Total() (float64, bool)

	// Rp2 возвращает реактивную энергию прямого направления по тарифу 2, кВар*ч
	Rp2() (float64, bool)

	// Rp2Total возвращает накопленное значение реактивной энергии прямого направления по тарифу 2 с момента сброса, кВар*ч
	Rp2Total() (float64, bool)

	// Rp3 возвращает реактивную энергию прямого направления по тарифу 3, кВар*ч
	Rp3() (float64, bool)

	// Rp3Total возвращает накопленное значение реактивной энергии прямого направления по тарифу 3 с момента
	// сброса, кВар*ч
	Rp3Total() (float64, bool)

	// Rp4 возвращает реактивную энергию прямого направления по тарифу 4, кВар*ч
	Rp4() (float64, bool)

	// Rp4Total возвращает накопленное значение реактивной энергии прямого направления по тарифу 4 с момента
	// сброса, кВар*ч
	Rp4Total() (float64, bool)

	// Rm1 возвращает реактивную энергию обратного направления по тарифу 1, кВар*ч
	Rm1() (float64, bool)

	// Rm1Total возвращает накопленное значение реактивной энергии обратного направления по тарифу 1 с момента
	// сброса, кВар*ч
	Rm1Total() (float64, bool)

	// Rm2 возвращает реактивную энергию обратного направления по тарифу 2, кВар*ч
	Rm2() (float64, bool)

	// Rm2Total возвращает накопленное значение реактивной энергии обратного направления по тарифу 2 с момента
	// сброса, кВар*ч
	Rm2Total() (float64, bool)

	// Rm3 возвращает реактивную энергию обратного направления по тарифу 3, кВар*ч
	Rm3() (float64, bool)

	// Rm3Total возвращает накопленное значение реактивной энергии обратного направления по тарифу 3 с момента
	// сброса, кВар*ч
	Rm3Total() (float64, bool)

	// Rm4 возвращает реактивную энергию обратного направления по тарифу 4, кВар*ч
	Rm4() (float64, bool)

	// Rm4Total возвращает накопленное значение реактивной энергии обратного направления по тарифу 4 с момента
	// сброса, кВар*ч
	Rm4Total() (float64, bool)

	// Ap возвращает сумму тарифов активной энергии прямого направления, кВт*ч
	Ap() (float64, bool)

	// ApTotal возвращает сумму тарифов накопленных значений активной энергии прямого направления с момента
	// сброса, кВт*ч
	ApTotal() (float64, bool)

	// Am возвращает сумму тарифов активной энергии обратного направления, кВар*ч
	Am() (float64, bool)

	// AmTotal возвращает сумму тарифов накопленных значений активной энергии обратного направления с момента
	// сброса, кВт*ч
	AmTotal() (float64, bool)

	// Rp возвращает сумму тарифов реактивной энергии прямого направления, кВар*ч
	Rp() (float64, bool)

	// RpTotal возвращает сумму тарифов накопленных значений реактивной энергии прямого направления с момента
	// сброса, кВар*ч
	RpTotal() (float64, bool)

	// Rm возвращает сумму тарифов реактивной энергии обратного направления, кВар*ч
	Rm() (float64, bool)

	// RmTotal возвращает сумму тарифов накопленных значений реактивной энергии обратного направления с момента
	// сброса, кВар*ч
	RmTotal() (float64, bool)

	// Pp возвращает активную мощность прямого направления, кВт
	Pp() (float64, bool)

	// Pm возвращает активную мощность обратного направления, кВт
	Pm() (float64, bool)

	// Qp возвращает реактивную мощность прямого направления, кВар
	Qp() (float64, bool)

	// Qm возвращает реактивную мощность обратного направления, кВар
	Qm() (float64, bool)

	// NS возвращает признак наличия нештатной ситуации
	NS() bool
}

// ElectricityCurrentReadings текущие нормализованные данные по электричеству
type ElectricityCurrentReadings interface {
	// Date возвращает дату архива, приборная
	Date() time.Time

	// DateOnEndOfArchive возвращает дату архива на конец часа
	DateOnEndOfArchive() time.Time

	// DateWithTimeBias возвращает дату архива на конец часа с учётом времени на приборе и часового пояса
	DateWithTimeBias() time.Time

	// S1 возвращает полную мощность по фазе 1, кВА
	S1() (float64, bool)

	// S2 возвращает полную мощность по фазе 2, кВА
	S2() (float64, bool)

	// S3 возвращает полную мощность по фазе 3, кВА
	S3() (float64, bool)

	// P1 возвращает активную мощность по фазе 1, кВт
	P1() (float64, bool)

	// P2 возвращает активную мощность по фазе 2, кВт
	P2() (float64, bool)

	// P3 возвращает активную мощность по фазе 3, кВт
	P3() (float64, bool)

	// Q1 возвращает реактивную мощность по фазе 1, кВар
	Q1() (float64, bool)

	// Q2 возвращает реактивную мощность по фазе 2, кВар
	Q2() (float64, bool)

	// Q3 возвращает реактивную мощность по фазе 3, кВар
	Q3() (float64, bool)

	// U1 возвращает напряжение по фазе 1, В
	U1() (float64, bool)

	// U2 возвращает напряжение по фазе 2, В
	U2() (float64, bool)

	// U3 возвращает напряжение по фазе 3, В
	U3() (float64, bool)

	// I1 возвращает силу тока по фазе 1, А
	I1() (float64, bool)

	// I2 возвращает силу тока по фазе 2, А
	I2() (float64, bool)

	// I3 возвращает силу тока по фазе 3, А
	I3() (float64, bool)

	// S возвращает сумму фаз полной мощности, кВА
	S() (float64, bool)

	// P возвращает сумму фаз активной мощности, кВт
	P() (float64, bool)

	// Q возвращает сумму фаз реактивной мощности, кВар
	Q() (float64, bool)

	// U возвращает сумму фаз напряжения, В
	U() (float64, bool)

	// I возвращает сумму фаз силы тока, А
	I() (float64, bool)

	// F возвращает частоту тока, Гц
	F() (float64, bool)

	// CosPhi1 возвращает коэффициент мощности по фазе 1
	CosPhi1() (float64, bool)

	// CosPhi2 возвращает коэффициент мощности по фазе 2
	CosPhi2() (float64, bool)

	// CosPhi3 возвращает коэффициент мощности по фазе 3
	CosPhi3() (float64, bool)

	// PhiU12 возвращает угол между фазными напряжениями фазы 1 и 2
	PhiU12() (float64, bool)

	// PhiU13 возвращает угол между фазными напряжениями фазы 1 и 3
	PhiU13() (float64, bool)

	// PhiU23 возвращает угол между фазными напряжениями фазы 2 и 3
	PhiU23() (float64, bool)
}

// GasReadings нормализованные данные по газу
type GasReadings interface {
	Readings

	// T возвращает температуру, °C
	T() (float64, bool)

	// P возвращает давление, МПа
	P() (float64, bool)

	// Vp возвращает рабочий объём, м3
	Vp() (float64, bool)

	// VpTotal возвращает накопленное значение рабочего объёма с момента сброса, м3
	VpTotal() (float64, bool)

	// Vc возвращает стандартный объём, м3
	Vc() (float64, bool)

	// VcTotal возвращает накопленное значение стандартного объёма с момента сброса, м3
	VcTotal() (float64, bool)
}
