package eldis

// Archive типы архивов показаний
type Archive string

const (
	// CurrentValues текущие значения (электричество)
	CurrentValues Archive = "CurrentValues"

	// MinuteArchive минутный архив
	MinuteArchive Archive = "MinuteArchive"

	// HourArchive часовой архив (ХВС, ГВС, тепло, сточные воды, газ, электричество)
	HourArchive Archive = "HourArchive"

	// DailyArchive суточный архив (ХВС, ГВС, тепло, сточные воды, газ, электричество)
	DailyArchive Archive = "DailyArchive"

	// MonthLongArchive месячный архив (ХВС, ГВС, тепло, сточные воды, газ)
	MonthLongArchive Archive = "MonthLongArchive"

	// TotalCurrentValues итоговые текущие значения (электричество)
	TotalCurrentValues Archive = "TotalCurrentValues"

	// IntervalArchive интервальный архив
	IntervalArchive Archive = "IntervalArchive"

	// HalfHourArchive получасовой архив (электричество)
	HalfHourArchive Archive = "HalfHourArchive"

	// DecadeArchive декадный архив
	DecadeArchive Archive = "DecadeArchive"

	// CurrentArchived текущий архивируемый
	CurrentArchived Archive = "CurrentArchived"
)

var codeArchive = map[Archive]string{
	CurrentValues:      "30001",
	MinuteArchive:      "30002",
	HourArchive:        "30003",
	DailyArchive:       "30004",
	MonthLongArchive:   "30005",
	TotalCurrentValues: "30006",
	IntervalArchive:    "30008",
	HalfHourArchive:    "30017",
	DecadeArchive:      "30022",
	CurrentArchived:    "30023",
}

// DateType тип даты для поиска данных
type DateType string

const (
	// Date приборная дата архива
	Date DateType = "date"

	// DateOnEndOfArchive дата архива на конец часа
	DateOnEndOfArchive DateType = "dateOnEndOfArchive"

	// DateWithTimeBias дата архива на конец часа с учетом времени на приборе и часового пояса
	DateWithTimeBias DateType = "dateWithTimeBias"
)
