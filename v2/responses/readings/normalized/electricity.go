package normalized

import "github.com/guregu/null"

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
