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

const (
	strCurrentValues      = "CurrentValues"
	strMinuteArchive      = "MinuteArchive"
	strHourArchive        = "HourArchive"
	strDailyArchive       = "DailyArchive"
	strMonthLongArchive   = "MonthLongArchive"
	strTotalCurrentValues = "TotalCurrentValues"
	strIntervalArchive    = "IntervalArchive"
	strHalfHourArchive    = "HalfHourArchive"
	strDecadeArchive      = "DecadeArchive"
	strCurrentArchived    = "CurrentArchived"
)

// String возвращает строковое представление значения типа архива показаний
func (a DataArchive) String() string {
	switch a {
	case CurrentValues:
		return strCurrentValues
	case MinuteArchive:
		return strMinuteArchive
	case HourArchive:
		return strHourArchive
	case DailyArchive:
		return strDailyArchive
	case MonthLongArchive:
		return strMonthLongArchive
	case TotalCurrentValues:
		return strTotalCurrentValues
	case IntervalArchive:
		return strIntervalArchive
	case HalfHourArchive:
		return strHalfHourArchive
	case DecadeArchive:
		return strDecadeArchive
	case CurrentArchived:
		return strCurrentArchived
	default:
		return ""
	}
}

// Parse возвращает значение типа DataArchive, соответствующее значению строки s
func Parse(s string) (DataArchive, error) {
	switch s {
	case strCurrentValues:
		return CurrentValues, nil
	case strMinuteArchive:
		return MinuteArchive, nil
	case strHourArchive:
		return HourArchive, nil
	case strDailyArchive:
		return DailyArchive, nil
	case strMonthLongArchive:
		return MonthLongArchive, nil
	case strTotalCurrentValues:
		return TotalCurrentValues, nil
	case strIntervalArchive:
		return IntervalArchive, nil
	case strHalfHourArchive:
		return HalfHourArchive, nil
	case strDecadeArchive:
		return DecadeArchive, nil
	case strCurrentArchived:
		return CurrentArchived, nil
	default:
		return UnknownArchive, &ErrorUnknownArchive{Value: s}
	}
}
