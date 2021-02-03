package date

// Type тип используемой даты для поиска данных
type Type byte

const (
	// Unknown неизвестный тип даты
	Unknown Type = iota
	// Date дата архива, приборная
	Date Type = iota
	// OnEndOfArchive дата архива на конец часа
	OnEndOfArchive
	// WithTimeBias дата архива на конец часа с учётом времени на приборе и часового пояса
	WithTimeBias
)

const (
	strDate           = "date"
	strOnEndOfArchive = "dateOnEndOfArchive"
	strWithTimeBias   = "dateWithTimeBias"
)

// String возвращает строковое представление значения типа даты для поиска данных
func (dt Type) String() string {
	switch dt {
	case Date:
		return strDate
	case OnEndOfArchive:
		return strOnEndOfArchive
	case WithTimeBias:
		return strWithTimeBias
	default:
		return ""
	}
}

// Parse возвращает значение типа Type, соответствующее значению строки s
func Parse(s string) (Type, error) {
	switch s {
	case strDate:
		return Date, nil
	case strOnEndOfArchive:
		return OnEndOfArchive, nil
	case strWithTimeBias:
		return WithTimeBias, nil
	default:
		return Unknown, &ErrorUnknownType{Value: s}
	}
}
