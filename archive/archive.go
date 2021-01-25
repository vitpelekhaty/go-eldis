package archive

// DataArchive архив данных
type DataArchive int

const (
	// UnknownArchive неопределенный архив
	UnknownArchive DataArchive = 0
	// CurrentValues текущие значения (электричество)
	CurrentValues DataArchive = 30001
	// MinuteArchive минутный архив
	MinuteArchive DataArchive = 30002
	// HourArchive часовой архив (ХВС, ГВС, тепло, сточные воды, газ, электричество)
	HourArchive DataArchive = 30003
	// DailyArchive суточный архив (ХВС, ГВС, тепло, сточные воды, газ, электричество)
	DailyArchive DataArchive = 30004
	// MonthLongArchive месячный архив (ХВС, ГВС, тепло, сточные воды, газ)
	MonthLongArchive DataArchive = 30005
	// TotalCurrentValues итоговые текущие значения (электричество)
	TotalCurrentValues DataArchive = 30006
	// IntervalArchive интервальный архив
	IntervalArchive DataArchive = 30008
	// HalfHourArchive получасовой архив (электричество)
	HalfHourArchive DataArchive = 30017
	// DecadeArchive декадный архив
	DecadeArchive DataArchive = 30022
	// CurrentArchived текущий архивируемый
	CurrentArchived DataArchive = 30023
)

// String возвращает строковое представление значения типа архива показаний
func (a DataArchive) String() string {
	switch a {
	case CurrentValues:
		return "CurrentValues"
	case MinuteArchive:
		return "MinuteArchive"
	case HourArchive:
		return "HourArchive"
	case DailyArchive:
		return "DailyArchive"
	case MonthLongArchive:
		return "MonthLongArchive"
	case TotalCurrentValues:
		return "TotalCurrentValues"
	case IntervalArchive:
		return "IntervalArchive"
	case HalfHourArchive:
		return "HalfHourArchive"
	case DecadeArchive:
		return "DecadeArchive"
	case CurrentArchived:
		return "CurrentArchived"
	default:
		return "UnknownArchive"
	}
}
