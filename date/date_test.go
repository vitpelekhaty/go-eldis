package date

import (
	"testing"
)

var typeToString = map[Type]string{
	Unknown:        "",
	Date:           strDate,
	OnEndOfArchive: strOnEndOfArchive,
	WithTimeBias:   strWithTimeBias,
}

var stringToType = map[string]Type{
	strDate:           Date,
	strOnEndOfArchive: OnEndOfArchive,
	strWithTimeBias:   WithTimeBias,
}

func TestType_String(t *testing.T) {
	for a, want := range typeToString {
		s := a.String()

		if s != want {
			t.Errorf("should be %s, but have %s", want, s)
		}
	}
}

func TestParse(t *testing.T) {
	for s, want := range stringToType {
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

	if _, ok := err.(*ErrorUnknownType); !ok {
		t.Fail()
	}
}
