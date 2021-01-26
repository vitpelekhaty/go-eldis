package eldis

// DateType тип используемой даты для поиска данных
type DateType byte

const (
	// Date дата архива, приборная
	Date DateType = iota
	// DateOnEndOfArchive дата архива на конец часа
	DateOnEndOfArchive
	// DateWithTimeBias дата архива на конец часа с учётом времени на приборе и часового пояса
	DateWithTimeBias
)

// String возвращает строковое представление значения типа даты для поиска данных
func (dt DateType) String() string {
	switch dt {
	case DateOnEndOfArchive:
		return "dateOnEndOfArchive"
	case DateWithTimeBias:
		return "dateWithTimeBias"
	default:
		return "date"
	}
}
