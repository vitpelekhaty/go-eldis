package archive

import (
	"testing"
)

var archiveToString = map[DataArchive]string{
	UnknownArchive:     "",
	CurrentValues:      strCurrentValues,
	MinuteArchive:      strMinuteArchive,
	HourArchive:        strHourArchive,
	DailyArchive:       strDailyArchive,
	MonthLongArchive:   strMonthLongArchive,
	TotalCurrentValues: strTotalCurrentValues,
	IntervalArchive:    strIntervalArchive,
	HalfHourArchive:    strHalfHourArchive,
	DecadeArchive:      strDecadeArchive,
	CurrentArchived:    strCurrentArchived,
}

var stringToArchive = map[string]DataArchive{
	strCurrentValues:      CurrentValues,
	strMinuteArchive:      MinuteArchive,
	strHourArchive:        HourArchive,
	strDailyArchive:       DailyArchive,
	strMonthLongArchive:   MonthLongArchive,
	strTotalCurrentValues: TotalCurrentValues,
	strIntervalArchive:    IntervalArchive,
	strHalfHourArchive:    HalfHourArchive,
	strDecadeArchive:      DecadeArchive,
	strCurrentArchived:    CurrentArchived,
}

func TestDataArchive_String(t *testing.T) {
	for a, want := range archiveToString {
		s := a.String()

		if s != want {
			t.Errorf("should be %s, but have %s", want, s)
		}
	}
}

func TestParse(t *testing.T) {
	for s, want := range stringToArchive {
		a, err := Parse(s)

		if err != nil {
			t.Errorf("Parse(%s): %q", s, err)
		}

		if a != want {
			t.Errorf("Parse(%s): should be %q, but have %q", s, want, a)
		}
	}
}

func TestParse2(t *testing.T) {
	_, err := Parse("test")

	if _, ok := err.(*ErrorUnknownArchive); !ok {
		t.Fail()
	}
}
