package response

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

func TestParseNormalizedEmptyData(t *testing.T) {
	const filename = "/testdata/dataNormalizedEmptyResponse.json"

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	path := filepath.Join(filepath.Dir(file), filename)

	body, err := ioutil.ReadFile(path)

	if err != nil {
		t.Errorf("%s: %q", filename, err)
	}

	ctx := context.TODO()

	t.Run(fmt.Sprintf("empty %s", normalizedWasteWaterPath), func(t *testing.T) {
		_, err := ParseWasteWaterNormalizedWithContext(ctx, body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty %s", normalizedColdWaterPath), func(t *testing.T) {
		_, err := ParseColdWaterNormalizedWithContext(ctx, body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty %s", normalizedHotWaterPath), func(t *testing.T) {
		_, err := ParseHotWaterNormalizedWithContext(ctx, body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty %s", normalizedHeatPath), func(t *testing.T) {
		_, err := ParseHeatNormalizedWithContext(ctx, body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty %s", normalizedElectricityPath), func(t *testing.T) {
		_, err := ParseElectricityNormalizedWithContext(ctx, body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty %s", normalizedElectricityCurrentPath), func(t *testing.T) {
		_, err := ParseElectricityCurrentNormalizedWithContext(ctx, body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty %s", normalizedGasPath), func(t *testing.T) {
		_, err := ParseGasNormalizedWithContext(ctx, body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}

func TestParseHotWaterNormalizedWithContext(t *testing.T) {
	var cases = [1]struct {
		path  string
		count int
	}{
		{path: "/testdata/dataNormalizedResponse.json", count: 120},
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

		items, err := ParseHotWaterNormalizedWithContext(ctx, body)

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
			t.Errorf("%s: %d point(s), but %d point(s) expected (broken test?)", test.path, count, test.count)
		}
	}
}
