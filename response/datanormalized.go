package response

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/guregu/null"
)

// CustomNormalized базовая структура нормализованных данных
type CustomNormalized struct {
	// Date дата архива, приборная
	Date int64 `json:"date"`
	// DateOnEndOfArchive дата архива на конец часа
	DateOnEndOfArchive int64 `json:"dateOnEndOfArchive"`
	// DateWithTimeBias дата архива на конец часа с учётом времени на приборе и часового пояса
	DateWithTimeBias int64 `json:"dateWithTimeBias"`
	// TGMax время, в течение которого фактический массовый расход был выше максимального нормированного значения для
	// средства измерения, ч
	TGMax null.Float `json:"TGmax,omitempty"`
	// TGMin время, в течение которого фактический массовый расход был меньше допустимого минимального нормированного
	// значения для средства измерения, ч
	TGMin null.Float `json:"TGmin,omitempty"`
	// TFault время функционального отказа, ч
	TFault null.Float `json:"TFault,omitempty"`
	// Toff время отключения питания, ч
	Toff null.Float `json:"Toff,omitempty"`
	// TOtherNS суммарное время других нештатных ситуаций, ч
	TOtherNS null.Float `json:"TOtherNS,omitempty"`
	// QntHIP время нормальной работы, ч
	QntHIP null.Float `json:"QntHIP,omitempty"`
	// QntHIPTotal накопленное значение времени нормальной работы с момента сброса, ч
	QntHIPTotal null.Float `json:"QntHIP_Total,omitempty"`
	// QntP время отсутствия счёта, ч
	QntP null.Float `json:"QntP,omitempty"`
	// QntPTotal накопленное значение времени отсутствия счёта с момента сброса, ч
	QntPTotal null.Float `json:"QntP_Total,omitempty"`
	// NS наличие нештатной ситуации
	NS bool `json:"ns"`
	// Empty признак пустого архива
	Empty bool `json:"empty"`
}

// WasteWaterNormalized нормализованные данные по сточным водам
type WasteWaterNormalized struct {
	CustomNormalized

	// V объём, м3
	V null.Float `json:"V,omitempty"`
	// VTotal накопленное значение объёма с момента сброса, м3
	VTotal null.Float `json:"V_Total,omitempty"`
	// M масса, т
	M null.Float `json:"M,omitempty"`
	// MTotal накопленное значение массы с момента сброса, т
	MTotal null.Float `json:"M_Total,omitempty"`
}

// ColdWaterNormalized нормализованные данные по ХВС
type ColdWaterNormalized struct {
	WasteWaterNormalized

	// P давление, МПа
	P null.Float `json:"P,omitempty"`
}

// HotWaterNormalized нормализованные данные по ГВС
type HotWaterNormalized struct {
	CustomNormalized

	// Q1 количество тепловой энергии на подающем трубопроводе, Гкал
	Q1 null.Float `json:"Q1,omitempty"`
	// Q1Total накопленное значение потребления тепловой энергии на подающем трубопроводе с момента сброса, Гкал
	Q1Total null.Float `json:"Q1_Total,omitempty"`
	// Q2 количество тепловой энергии на обратном трубопроводе, Гкал
	Q2 null.Float `json:"Q2,omitempty"`
	// Q2Total накопленное значение потребления тепловой энергии на обратном трубопроводе с момента сброса, Гкал
	Q2Total null.Float `json:"Q2_Total,omitempty"`
	// T1 температура на подающем трубопроводе, °C
	T1 null.Float `json:"t1,omitempty"`
	// T2 температура на обратном трубопроводе, °C
	T2 null.Float `json:"t2,omitempty"`
	// V1 объём на подающем трубопроводе, м3
	V1 null.Float `json:"V1,omitempty"`
	// V1Total накопленное значение объёма на подающем трубопроводе с момента сброса, м3
	V1Total null.Float `json:"V1_Total,omitempty"`
	// V2 объём на обратном трубопроводе, м3
	V2 null.Float `json:"V2,omitempty"`
	// V2Total накопленное значение объёма на обратном трубопроводе с момента сброса, м3
	V2Total null.Float `json:"V2_Total,omitempty"`
	// M1 масса на подающем трубопроводе, т
	M1 null.Float `json:"M1,omitempty"`
	// M1Total накопленное значение массы на подающем трубопроводе с момента сброса, т
	M1Total null.Float `json:"M1_Total,omitempty"`
	// M2 масса на обратном трубопроводе, т
	M2 null.Float `json:"M2,omitempty"`
	// M2Total накопленное значение массы на обратном трубопроводе с момента сброса, т
	M2Total null.Float `json:"M2_Total,omitempty"`
	// P1 давление на подающем трубопроводе, МПа
	P1 null.Float `json:"P1,omitempty"`
	// P2 давление на обратном трубопроводе, МПа
	P2 null.Float `json:"P2,omitempty"`
	// Q количество тепловой энергии по всей системе, Гкал
	Q null.Float `json:"Q,omitempty"`
	// QTotal накопленное значение потребления тепловой энергии по всей системе с момента сброса, Гкал
	QTotal null.Float `json:"Q_Total,omitempty"`
	// DT разность температур, °C
	DT null.Float `json:"dt,omitempty"`
	// DV разность объёмов, м3
	DV null.Float `json:"dV,omitempty"`
	// DM разность масс, т
	DM null.Float `json:"dM,omitempty"`
	// Tcw температура холодной воды, °C
	Tcw null.Float `json:"tcw,omitempty"`
	// Pcw давление холодной воды, МПа
	Pcw null.Float `json:"Pcw,omitempty"`
	// Tdt время, в течение которого разность температур на подающем и обратном трубопроводах была меньше допустимой
	// нормированной разности температур для теплосчетчика, ч
	Tdt null.Float `json:"Tdt,omitempty"`
}

// HeatNormalized нормализованные данные по теплу
type HeatNormalized struct {
	HotWaterNormalized

	// T3 температура на подпитке, °C
	T3 null.Float `json:"t3,omitempty"`
	// V3 объём на подпитке, м3
	V3 null.Float `json:"V3,omitempty"`
	// V3Total накопленное значение объёма на подпитке с момента сброса, м3
	V3Total null.Float `json:"V3_Total,omitempty"`
	// M3 масса на подпитке, т
	M3 null.Float `json:"M3,omitempty"`
	// M3Total накопленное значение массы на подпитке с момента сброса, т
	M3Total null.Float `json:"M3_Total,omitempty"`
	// P3 давление на подпитке, МПа
	P3 null.Float `json:"P3,omitempty"`
	// Ta температура наружного воздуха, °C
	Ta null.Float `json:"ta,omitempty"`
}

// ElectricityNormalized нормализованные данные по электричеству
type ElectricityNormalized struct {
	// Date дата архива, приборная
	Date int64 `json:"date"`
	// DateOnEndOfArchive дата архива на конец часа
	DateOnEndOfArchive int64 `json:"dateOnEndOfArchive"`
	// DateWithTimeBias дата архива на конец часа с учётом времени на приборе и часового пояса
	DateWithTimeBias int64 `json:"dateWithTimeBias"`
	// Ap1 активная энергия прямого направления по тарифу 1, кВт*ч
	Ap1 null.Float `json:"Ap1,omitempty"`
	// Ap1Total накопленное значение активной энергии прямого направления по тарифу 1 с момента сброса, кВт*ч
	Ap1Total null.Float `json:"Ap1_Total,omitempty"`
	// Ap2 активная энергия прямого направления по тарифу 2, кВт*ч
	Ap2 null.Float `json:"Ap2,omitempty"`
	// Ap2Total накопленное значение активной энергии прямого направления по тарифу 2 с момента сброса, кВт*ч
	Ap2Total null.Float `json:"Ap2_Total,omitempty"`
	// Ap3 активная энергия прямого направления по тарифу 3, кВт*ч
	Ap3 null.Float `json:"Ap3,omitempty"`
	// Ap3Total накопленное значение активной энергии прямого направления по тарифу 3 с момента сброса, кВт*ч
	Ap3Total null.Float `json:"Ap3_Total,omitempty"`
	// Ap4 активная энергия прямого направления по тарифу 4, кВт*ч
	Ap4 null.Float `json:"Ap4,omitempty"`
	// Ap4Total накопленное значение активной энергии прямого направления по тарифу 4 с момента сброса, кВт*ч
	Ap4Total null.Float `json:"Ap4_Total,omitempty"`
	// Am1 активная энергия обратного направления по тарифу 1, кВт*ч
	Am1 null.Float `json:"Am1,omitempty"`
	// Am1Total накопленное значение активной энергии обратного направления по тарифу 1 с момента сброса, кВт*ч
	Am1Total null.Float `json:"Am1_Total,omitempty"`
	// Am2 активная энергия обратного направления по тарифу 2, кВт*ч
	Am2 null.Float `json:"Am2,omitempty"`
	// Am2Total накопленное значение активной энергии обратного направления по тарифу 2 с момента сброса, кВт*ч
	Am2Total null.Float `json:"Am2_Total,omitempty"`
	// Am3 активная энергия обратного направления по тарифу 3, кВт*ч
	Am3 null.Float `json:"Am3,omitempty"`
	// Am3Total накопленное значение активной энергии обратного направления по тарифу 3 с момента сброса, кВт*ч
	Am3Total null.Float `json:"Am3_Total,omitempty"`
	// Am4 активная энергия обратного направления по тарифу 4, кВт*ч
	Am4 null.Float `json:"Am4,omitempty"`
	// Am4Total накопленное значение активной энергии обратного направления по тарифу 4 с момента сброса, кВт*ч
	Am4Total null.Float `json:"Am4_Total,omitempty"`
	// Rp1 реактивная энергия прямого направления по тарифу 1, кВар*ч
	Rp1 null.Float `json:"Rp1,omitempty"`
	// Rp1Total накопленное значение реактивной энергии прямого направления по тарифу 1 с момента сброса, кВар*ч
	Rp1Total null.Float `json:"Rp1_Total,omitempty"`
	// Rp2 реактивная энергия прямого направления по тарифу 2, кВар*ч
	Rp2 null.Float `json:"Rp2,omitempty"`
	// Rp2Total накопленное значение реактивной энергии прямого направления по тарифу 2 с момента сброса, кВар*ч
	Rp2Total null.Float `json:"Rp2_Total,omitempty"`
	// Rp3 реактивная энергия прямого направления по тарифу 3, кВар*ч
	Rp3 null.Float `json:"Rp3,omitempty"`
	// Rp3Total накопленное значение реактивной энергии прямого направления по тарифу 3 с момента сброса, кВар*ч
	Rp3Total null.Float `json:"Rp3_Total,omitempty"`
	// Rp4 реактивная энергия прямого направления по тарифу 4, кВар*ч
	Rp4 null.Float `json:"Rp4,omitempty"`
	// Rp4Total накопленное значение реактивной энергии прямого направления по тарифу 4 с момента сброса, кВар*ч
	Rp4Total null.Float `json:"Rp4_Total,omitempty"`
	// Rm1 реактивная энергия обратного направления по тарифу 1, кВар*ч
	Rm1 null.Float `json:"Rm1,omitempty"`
	// Rm1Total накопленное значение реактивной энергии обратного направления по тарифу 1 с момента сброса, кВар*ч
	Rm1Total null.Float `json:"Rm1_Total,omitempty"`
	// Rm2 реактивная энергия обратного направления по тарифу 2, кВар*ч
	Rm2 null.Float `json:"Rm2,omitempty"`
	// Rm2Total накопленное значение реактивной энергии обратного направления по тарифу 2 с момента сброса, кВар*ч
	Rm2Total null.Float `json:"Rm2_Total,omitempty"`
	// Rm3 реактивная энергия обратного направления по тарифу 3, кВар*ч
	Rm3 null.Float `json:"Rm3,omitempty"`
	// Rm3Total накопленное значение реактивной энергии обратного направления по тарифу 3 с момента сброса, кВар*ч
	Rm3Total null.Float `json:"Rm3_Total,omitempty"`
	// Rm4 реактивная энергия обратного направления по тарифу 4, кВар*ч
	Rm4 null.Float `json:"Rm4,omitempty"`
	// Rm4Total накопленное значение реактивной энергии обратного направления по тарифу 4 с момента сброса, кВар*ч
	Rm4Total null.Float `json:"Rm4_Total,omitempty"`
	// Ap сумма тарифов активной энергии прямого направления, кВт*ч
	Ap null.Float `json:"Ap,omitempty"`
	// ApTotal сумма тарифов накопленных значений активной энергии прямого направления с момента сброса, кВт*ч
	ApTotal null.Float `json:"Ap_Total,omitempty"`
	// Am сумма тарифов активной энергии обратного направления, кВар*ч
	Am null.Float `json:"Am,omitempty"`
	// AmTotal сумма тарифов накопленных значений активной энергии обратного направления с момента сброса, кВт*ч
	AmTotal null.Float `json:"Am_Total,omitempty"`
	// Rp сумма тарифов реактивной энергии прямого направления, кВар*ч
	Rp null.Float `json:"Rp,omitempty"`
	// RpTotal сумма тарифов накопленных значений реактивной энергии прямого направления с момента сброса, кВар*ч
	RpTotal null.Float `json:"Rp_Total,omitempty"`
	// Rm сумма тарифов реактивной энергии обратного направления, кВар*ч
	Rm null.Float `json:"Rm,omitempty"`
	// RmTotal сумма тарифов накопленных значений реактивной энергии обратного направления с момента сброса, кВар*ч
	RmTotal null.Float `json:"Rm_Total,omitempty"`
	// Pp активная мощность прямого направления, кВт
	Pp null.Float `json:"Pp,omitempty"`
	// Pm активная мощность обратного направления, кВт
	Pm null.Float `json:"Pm,omitempty"`
	// Qp реактивная мощность прямого направления, кВар
	Qp null.Float `json:"Qp,omitempty"`
	// Qm реактивная мощность обратного направления, кВар
	Qm null.Float `json:"Qm,omitempty"`
	// NS наличие нештатной ситуации
	NS bool `json:"ns"`
}

// ElectricityCurrentNormalized текущие нормализованные данные по электричеству
type ElectricityCurrentNormalized struct {
	// Date дата архива, приборная
	Date int64 `json:"date"`
	// DateOnEndOfArchive дата архива на конец часа
	DateOnEndOfArchive int64 `json:"dateOnEndOfArchive"`
	// DateWithTimeBias дата архива на конец часа с учётом времени на приборе и часового пояса
	DateWithTimeBias int64 `json:"dateWithTimeBias"`
	// S1 полная мощность по фазе 1, кВА
	S1 null.Float `json:"S1,omitempty"`
	// S2 полная мощность по фазе 2, кВА
	S2 null.Float `json:"S2,omitempty"`
	// S3 полная мощность по фазе 3, кВА
	S3 null.Float `json:"S3,omitempty"`
	// P1 активная мощность по фазе 1, кВт
	P1 null.Float `json:"P1,omitempty"`
	// P2 активная мощность по фазе 2, кВт
	P2 null.Float `json:"P2,omitempty"`
	// P3 активная мощность по фазе 3, кВт
	P3 null.Float `json:"P3,omitempty"`
	// Q1 реактивная мощность по фазе 1, кВар
	Q1 null.Float `json:"Q1,omitempty"`
	// Q2 реактивная мощность по фазе 2, кВар
	Q2 null.Float `json:"Q2,omitempty"`
	// Q3 реактивная мощность по фазе 3, кВар
	Q3 null.Float `json:"Q3,omitempty"`
	// U1 напряжение по фазе 1, В
	U1 null.Float `json:"U1,omitempty"`
	// U2 напряжение по фазе 2, В
	U2 null.Float `json:"U2,omitempty"`
	// U3 напряжение по фазе 3, В
	U3 null.Float `json:"U3,omitempty"`
	// I1 сила тока по фазе 1, А
	I1 null.Float `json:"I1,omitempty"`
	// I2 сила тока по фазе 2, А
	I2 null.Float `json:"I2,omitempty"`
	// I3 сила тока по фазе 3, А
	I3 null.Float `json:"I3,omitempty"`
	// S сумма фаз полной мощности, кВА
	S null.Float `json:"S,omitempty"`
	// P сумма фаз активной мощности, кВт
	P null.Float `json:"P,omitempty"`
	// Q сумма фаз реактивной мощности, кВар
	Q null.Float `json:"Q,omitempty"`
	// U сумма фаз напряжения, В
	U null.Float `json:"U,omitempty"`
	// I сумма фаз силы тока, А
	I null.Float `json:"I,omitempty"`
	// F частота тока, Гц
	F null.Float `json:"F,omitempty"`
	// CosPhi1 коэффициент мощности по фазе 1
	CosPhi1 null.Float `json:"CosPhi1,omitempty"`
	// CosPhi2 коэффициент мощности по фазе 2
	CosPhi2 null.Float `json:"CosPhi2,omitempty"`
	// CosPhi3 коэффициент мощности по фазе 3
	CosPhi3 null.Float `json:"CosPhi3,omitempty"`
	// PhiU12 угол между фазными напряжениями фазы 1 и 2
	PhiU12 null.Float `json:"PhiU12,omitempty"`
	// PhiU13 угол между фазными напряжениями фазы 1 и 3
	PhiU13 null.Float `json:"PhiU13,omitempty"`
	// PhiU23 угол между фазными напряжениями фазы 2 и 3
	PhiU23 null.Float `json:"PhiU23,omitempty"`
}

// GasNormalized нормализованные данные по газу
type GasNormalized struct {
	CustomNormalized

	// T температура, °C
	T null.Float `json:"t,omitempty"`
	// P давление, МПа
	P null.Float `json:"P,omitempty"`
	// Vp рабочий объём, м3
	Vp null.Float `json:"Vp,omitempty"`
	// VpTotal накопленное значение рабочего объёма с момента сброса, м3
	VpTotal null.Float `json:"Vp_Total,omitempty"`
	// Vc стандартный объём, м3
	Vc null.Float `json:"Vc,omitempty"`
	// VcTotal накопленное значение стандартного объёма с момента сброса, м3
	VcTotal null.Float `json:"Vc_Total,omitempty"`
}

const (
	normalizedWasteWaterPath         = "response.data.normalized.#(wasteWater).wasteWater"
	normalizedColdWaterPath          = "response.data.normalized.#(coldWater).coldWater"
	normalizedHotWaterPath           = "response.data.normalized.#(hotWater).hotWater"
	normalizedHeatPath               = "response.data.normalized.#(heat).heat"
	normalizedElectricityPath        = "response.data.normalized.#(electricity).electricity"
	normalizedElectricityCurrentPath = "response.data.normalized.#(electricityCurrent).electricityCurrent"
	normalizedGasPath                = "response.data.normalized.#(gas).gas"
)

// ParseWasteWaterNormalizedWithContext возвращает канал, в который будет записывать нормализованные данные по сточным
// водам, полученные в ответе API при вызове метода /api/v2/data/normalized.
//
// body - ответ указанного метода API; ctx - контекст, с помощью которого можно при необходимости остановить чтение
// массива нормализованных данных
//
// Чтение массива данных будет выполняться до полного чтения элементов массива или до вызова функции cancel отмены
// контекста
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseWasteWaterNormalizedWithContext(ctx context.Context, body []byte) (<-chan struct {
	Data *WasteWaterNormalized
	Err  error
}, error) {
	raw, err := getBytes(body, normalizedWasteWaterPath)

	if err != nil {
		return nil, err
	}

	out := make(chan struct {
		Data *WasteWaterNormalized
		Err  error
	})

	go func(b []byte) {
		defer close(out)

		decoder := json.NewDecoder(bytes.NewReader(b))

		_, err := decoder.Token()

		if err != nil {
			return
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !decoder.More() {
					return
				}

				var data WasteWaterNormalized

				if err := decoder.Decode(&data); err == nil {
					out <- struct {
						Data *WasteWaterNormalized
						Err  error
					}{Data: &data, Err: nil}
				} else {
					out <- struct {
						Data *WasteWaterNormalized
						Err  error
					}{Data: nil, Err: err}
				}
			}
		}
	}(raw)

	return out, nil
}

// ParseWasteWaterNormalized возвращает канал, в который будет записывать нормализованные данные по сточным водам,
// полученные в ответе API при вызове метода /api/v2/data/normalized.
//
// body - ответ указанного метода API
//
// Чтение массива данных будет выполняться до полного чтения элементов массива
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseWasteWaterNormalized(body []byte) (<-chan struct {
	Data *WasteWaterNormalized
	Err  error
}, error) {
	return ParseWasteWaterNormalizedWithContext(context.TODO(), body)
}

// ParseColdWaterNormalizedWithContext возвращает канал, в который будет записывать нормализованные данные по холодной
// воде, полученные в ответе API при вызове метода /api/v2/data/normalized.
//
// body - ответ указанного метода API; ctx - контекст, с помощью которого можно при необходимости остановить чтение
// массива нормализованных данных
//
// Чтение массива данных будет выполняться до полного чтения элементов массива или до вызова функции cancel отмены
// контекста
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseColdWaterNormalizedWithContext(ctx context.Context, body []byte) (<-chan struct {
	Data *ColdWaterNormalized
	Err  error
}, error) {
	raw, err := getBytes(body, normalizedColdWaterPath)

	if err != nil {
		return nil, err
	}

	out := make(chan struct {
		Data *ColdWaterNormalized
		Err  error
	})

	go func(b []byte) {
		defer close(out)

		decoder := json.NewDecoder(bytes.NewReader(b))

		_, err := decoder.Token()

		if err != nil {
			return
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !decoder.More() {
					return
				}

				var data ColdWaterNormalized

				if err := decoder.Decode(&data); err == nil {
					out <- struct {
						Data *ColdWaterNormalized
						Err  error
					}{Data: &data, Err: nil}
				} else {
					out <- struct {
						Data *ColdWaterNormalized
						Err  error
					}{Data: nil, Err: err}
				}
			}
		}
	}(raw)

	return out, nil
}

// ParseColdWaterNormalized возвращает канал, в который будет записывать нормализованные данные по холодной воде,
// полученные в ответе API при вызове метода /api/v2/data/normalized.
//
// body - ответ указанного метода API
//
// Чтение массива данных будет выполняться до полного чтения элементов массива
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseColdWaterNormalized(body []byte) (<-chan struct {
	Data *ColdWaterNormalized
	Err  error
}, error) {
	return ParseColdWaterNormalizedWithContext(context.TODO(), body)
}

// ParseHotWaterNormalizedWithContext возвращает канал, в который будет записывать нормализованные данные по горячей
// воде, полученные в ответе API при вызове метода /api/v2/data/normalized.
//
// body - ответ указанного метода API; ctx - контекст, с помощью которого можно при необходимости остановить чтение
// массива нормализованных данных
//
// Чтение массива данных будет выполняться до полного чтения элементов массива или до вызова функции cancel отмены
// контекста
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseHotWaterNormalizedWithContext(ctx context.Context, body []byte) (<-chan struct {
	Data *HotWaterNormalized
	Err  error
}, error) {
	raw, err := getBytes(body, normalizedHotWaterPath)

	if err != nil {
		return nil, err
	}

	out := make(chan struct {
		Data *HotWaterNormalized
		Err  error
	})

	go func(b []byte) {
		defer close(out)

		decoder := json.NewDecoder(bytes.NewReader(b))

		_, err := decoder.Token()

		if err != nil {
			return
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !decoder.More() {
					return
				}

				var data HotWaterNormalized

				if err := decoder.Decode(&data); err == nil {
					out <- struct {
						Data *HotWaterNormalized
						Err  error
					}{Data: &data, Err: nil}
				} else {
					out <- struct {
						Data *HotWaterNormalized
						Err  error
					}{Data: nil, Err: err}
				}
			}
		}
	}(raw)

	return out, nil
}

// ParseHotWaterNormalized возвращает канал, в который будет записывать нормализованные данные по горячей воде,
// полученные в ответе API при вызове метода /api/v2/data/normalized.
//
// body - ответ указанного метода API
//
// Чтение массива данных будет выполняться до полного чтения элементов массива
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseHotWaterNormalized(body []byte) (<-chan struct {
	Data *HotWaterNormalized
	Err  error
}, error) {
	return ParseHotWaterNormalizedWithContext(context.TODO(), body)
}

// ParseHeatNormalizedWithContext возвращает канал, в который будет записывать нормализованные данные по теплу,
// полученные в ответе API при вызове метода /api/v2/data/normalized.
//
// body - ответ указанного метода API; ctx - контекст, с помощью которого можно при необходимости остановить чтение
// массива нормализованных данных
//
// Чтение массива данных будет выполняться до полного чтения элементов массива или до вызова функции cancel отмены
// контекста
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseHeatNormalizedWithContext(ctx context.Context, body []byte) (<-chan struct {
	Data *HeatNormalized
	Err  error
}, error) {
	raw, err := getBytes(body, normalizedHeatPath)

	if err != nil {
		return nil, err
	}

	out := make(chan struct {
		Data *HeatNormalized
		Err  error
	})

	go func(b []byte) {
		defer close(out)

		decoder := json.NewDecoder(bytes.NewReader(b))

		_, err := decoder.Token()

		if err != nil {
			return
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !decoder.More() {
					return
				}

				var data HeatNormalized

				if err := decoder.Decode(&data); err == nil {
					out <- struct {
						Data *HeatNormalized
						Err  error
					}{Data: &data, Err: nil}
				} else {
					out <- struct {
						Data *HeatNormalized
						Err  error
					}{Data: nil, Err: err}
				}
			}
		}
	}(raw)

	return out, nil
}

// ParseHeatNormalized возвращает канал, в который будет записывать нормализованные данные по теплу, полученные в ответе
// API при вызове метода /api/v2/data/normalized.
//
// body - ответ указанного метода API
//
// Чтение массива данных будет выполняться до полного чтения элементов массива
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseHeatNormalized(body []byte) (<-chan struct {
	Data *HeatNormalized
	Err  error
}, error) {
	return ParseHeatNormalizedWithContext(context.TODO(), body)
}

// ParseElectricityNormalizedWithContext возвращает канал, в который будет записывать нормализованные данные по
// электричеству, полученные в ответе API при вызове метода /api/v2/data/normalized.
//
// body - ответ указанного метода API; ctx - контекст, с помощью которого можно при необходимости остановить чтение
// массива нормализованных данных
//
// Чтение массива данных будет выполняться до полного чтения элементов массива или до вызова функции cancel отмены
// контекста
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseElectricityNormalizedWithContext(ctx context.Context, body []byte) (<-chan struct {
	Data *ElectricityNormalized
	Err  error
}, error) {
	raw, err := getBytes(body, normalizedElectricityPath)

	if err != nil {
		return nil, err
	}

	out := make(chan struct {
		Data *ElectricityNormalized
		Err  error
	})

	go func(b []byte) {
		defer close(out)

		decoder := json.NewDecoder(bytes.NewReader(b))

		_, err := decoder.Token()

		if err != nil {
			out <- struct {
				Data *ElectricityNormalized
				Err  error
			}{Data: nil, Err: err}

			return
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !decoder.More() {
					return
				}

				var data ElectricityNormalized

				if err := decoder.Decode(&data); err == nil {
					out <- struct {
						Data *ElectricityNormalized
						Err  error
					}{Data: &data, Err: nil}
				} else {
					out <- struct {
						Data *ElectricityNormalized
						Err  error
					}{Data: nil, Err: err}
				}
			}
		}
	}(raw)

	return out, nil
}

// ParseElectricityNormalized возвращает канал, в который будет записывать нормализованные данные по электричеству,
// полученные в ответе API при вызове метода /api/v2/data/normalized.
//
// body - ответ указанного метода API
//
// Чтение массива данных будет выполняться до полного чтения элементов массива
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseElectricityNormalized(body []byte) (<-chan struct {
	Data *ElectricityNormalized
	Err  error
}, error) {
	return ParseElectricityNormalizedWithContext(context.TODO(), body)
}

// ParseElectricityCurrentNormalizedWithContext возвращает канал, в который будет записывать нормализованные текущие
// данные по электричеству, полученные в ответе API при вызове метода /api/v2/data/normalized.
//
// body - ответ указанного метода API; ctx - контекст, с помощью которого можно при необходимости остановить чтение
// массива нормализованных данных
//
// Чтение массива данных будет выполняться до полного чтения элементов массива или до вызова функции cancel отмены
// контекста
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseElectricityCurrentNormalizedWithContext(ctx context.Context, body []byte) (<-chan struct {
	Data *ElectricityCurrentNormalized
	Err  error
}, error) {
	raw, err := getBytes(body, normalizedElectricityCurrentPath)

	if err != nil {
		return nil, err
	}

	out := make(chan struct {
		Data *ElectricityCurrentNormalized
		Err  error
	})

	go func(b []byte) {
		defer close(out)

		decoder := json.NewDecoder(bytes.NewReader(b))

		_, err := decoder.Token()

		if err != nil {
			return
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !decoder.More() {
					return
				}

				var data ElectricityCurrentNormalized

				if err := decoder.Decode(&data); err == nil {
					out <- struct {
						Data *ElectricityCurrentNormalized
						Err  error
					}{Data: &data, Err: nil}
				} else {
					out <- struct {
						Data *ElectricityCurrentNormalized
						Err  error
					}{Data: nil, Err: err}
				}
			}
		}
	}(raw)

	return out, nil
}

// ParseElectricityCurrentNormalized возвращает канал, в который будет записывать нормализованные текущие данные по
// электричеству, полученные в ответе API при вызове метода /api/v2/data/normalized.
//
// body - ответ указанного метода API
//
// Чтение массива данных будет выполняться до полного чтения элементов массива
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseElectricityCurrentNormalized(body []byte) (<-chan struct {
	Data *ElectricityCurrentNormalized
	Err  error
}, error) {
	return ParseElectricityCurrentNormalizedWithContext(context.TODO(), body)
}

// ParseGasNormalizedWithContext возвращает канал, в который будет записывать нормализованные данные по газу,
// полученные в ответе API при вызове метода /api/v2/data/normalized.
//
// body - ответ указанного метода API; ctx - контекст, с помощью которого можно при необходимости остановить чтение
// массива нормализованных данных
//
// Чтение массива данных будет выполняться до полного чтения элементов массива или до вызова функции cancel отмены
// контекста
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseGasNormalizedWithContext(ctx context.Context, body []byte) (<-chan struct {
	Data *GasNormalized
	Err  error
}, error) {
	raw, err := getBytes(body, normalizedGasPath)

	if err != nil {
		return nil, err
	}

	out := make(chan struct {
		Data *GasNormalized
		Err  error
	})

	go func(b []byte) {
		defer close(out)

		decoder := json.NewDecoder(bytes.NewReader(b))

		_, err := decoder.Token()

		if err != nil {
			return
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !decoder.More() {
					return
				}

				var data GasNormalized

				if err := decoder.Decode(&data); err == nil {
					out <- struct {
						Data *GasNormalized
						Err  error
					}{Data: &data, Err: nil}
				} else {
					out <- struct {
						Data *GasNormalized
						Err  error
					}{Data: nil, Err: err}
				}
			}
		}
	}(raw)

	return out, nil
}

// ParseGasNormalized возвращает канал, в который будет записывать нормализованные данные по газу, полученные в ответе
// API при вызове метода /api/v2/data/normalized.
//
// body - ответ указанного метода API
//
// Чтение массива данных будет выполняться до полного чтения элементов массива
//
// Чтобы прочитать сообщения о результатах обработки запроса, необходимо воспользоваться методом Parse
func ParseGasNormalized(body []byte) (<-chan struct {
	Data *GasNormalized
	Err  error
}, error) {
	return ParseGasNormalizedWithContext(context.TODO(), body)
}
