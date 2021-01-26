package eldis

import (
	"time"
)

const requestTimeLayout = `02.01.2006 15:04:05`

// RequestTime описывает формат времени, принятый в запросах к АИСКУТЭ ЭЛДИС
type RequestTime time.Time

// String возвращает строковое представление типа RequestTime
func (rt *RequestTime) String() string {
	t := time.Time(*rt)
	return t.Format(requestTimeLayout)
}
