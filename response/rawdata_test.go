package response

import (
	"context"
	"errors"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestRawValue(t *testing.T) {
	const (
		stringValue   = "test"
		intValue      = 100
		floatValue    = 3.1415926535
		rawTimeLayout = "2006-01-02T03:04:05"
	)

	timeValue := time.Date(2020, 12, 12, 3, 0, 0, 0, time.UTC)

	t.Run("RawValue.String", func(t *testing.T) {
		value := stringValue
		raw := RawValue(value)

		s, err := raw.String()

		if err != nil {
			t.Fatal(err)
		}

		if s != value {
			t.Fail()
		}
	})

	t.Run("RawValue.AsInt64", func(t *testing.T) {
		value := strconv.Itoa(intValue)
		raw := RawValue(value)

		i, err := raw.AsInt()

		if err != nil {
			t.Fatal(err)
		}

		if i != intValue {
			t.Fail()
		}
	})

	t.Run("RawValue.AsInt64.WithError", func(t *testing.T) {
		raw := RawValue(stringValue)

		i, _ := raw.AsInt()

		if i != -1 {
			t.Fail()
		}
	})

	t.Run("RawValue.AsFloat", func(t *testing.T) {
		value := strconv.FormatFloat(floatValue, 'f', -1, 64)
		raw := RawValue(value)

		f, err := raw.AsFloat()

		if err != nil {
			t.Fatal(err)
		}

		if f != floatValue {
			t.Fail()
		}
	})

	t.Run("RawValue.AsFloat.WithInvalidSep", func(t *testing.T) {
		value := strconv.FormatFloat(floatValue, 'f', -1, 64)
		value = strings.ReplaceAll(value, ".", ",")

		raw := RawValue(value)

		f, err := raw.AsFloat()

		if err != nil {
			t.Fatal(err)
		}

		if f != floatValue {
			t.Fail()
		}
	})

	t.Run("RawValue.AsTime", func(t *testing.T) {
		value := timeValue.Format(rawTimeLayout)
		raw := RawValue(value)

		tt, err := raw.AsTime()

		if err != nil {
			t.Fatal(err)
		}

		if tt != timeValue {
			t.Fail()
		}
	})
}

func TestParseRawDataWithContext(t *testing.T) {
	var cases = [1]struct {
		path  string
		count int
	}{
		{path: "/testdata/rawDataResponse.json", count: 120},
	}

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	ctx := context.TODO()

	for _, test := range cases {
		count := 0

		path := filepath.Join(filepath.Dir(file), test.path)

		body, err := ioutil.ReadFile(path)

		if err != nil {
			t.Fatalf("%s: %q", test.path, err)
		}

		items, err := ParseRawDataWithContext(ctx, body)

		if err != nil {
			t.Fatalf("%s: %q", test.path, err)
		}

		for item := range items {
			if item.Err != nil {
				t.Errorf("%s: %q", test.path, err)
			} else {
				count++
			}
		}

		if test.count != count {
			t.Errorf("%s: %d item(s), but %d item(s) expected (broken test?)", test.path, count, test.count)
		}
	}
}

func TestParseRawDataWithCancel(t *testing.T) {
	filename := "/testdata/rawDataResponse.json"
	const expectedCount = 120

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	path := filepath.Join(filepath.Dir(file), filename)

	body, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatalf("%s: %q", filename, err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	items, err := ParseRawDataWithContext(ctx, body)

	if err != nil {
		t.Fatalf("%s: %q", filename, err)
	}

	var count int

	for item := range items {
		if item.Err != nil {
			t.Errorf("%s: %q", filename, err)
		}

		count++
		cancel()
	}

	if !(count < expectedCount) {
		t.Errorf("%s: %d step(s), you cannot stop me", filename, count)
	}
}

func TestParseRawData(t *testing.T) {
	var cases = [1]struct {
		path  string
		count int
	}{
		{path: "/testdata/rawDataResponse.json", count: 120},
	}

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	for _, test := range cases {
		count := 0

		path := filepath.Join(filepath.Dir(file), test.path)

		body, err := ioutil.ReadFile(path)

		if err != nil {
			t.Fatalf("%s: %q", test.path, err)
		}

		items, err := ParseRawData(body)

		if err != nil {
			t.Fatalf("%s: %q", test.path, err)
		}

		for item := range items {
			if item.Err != nil {
				t.Errorf("%s: %q", test.path, err)
			} else {
				count++
			}
		}

		if test.count != count {
			t.Errorf("%s: %d item(s), but %d item(s) expected (broken test?)", test.path, count, test.count)
		}
	}
}

func TestParseRawDataWithEmptyBody(t *testing.T) {
	var body []byte

	t.Run("ParseRawDataWithContext.WithEmptyBody", func(t *testing.T) {
		_, err := ParseRawDataWithContext(context.TODO(), body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("ParseRawData.WithEmptyBody", func(t *testing.T) {
		_, err := ParseRawData(body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})
}

func TestParseRawDataWithEmptyBody2(t *testing.T) {
	var body = []byte("{}")

	t.Run("ParseRawDataWithContext.WithEmptyResponse", func(t *testing.T) {
		_, err := ParseRawDataWithContext(context.TODO(), body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("ParseRawData.WithEmptyResponse", func(t *testing.T) {
		_, err := ParseRawData(body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}
