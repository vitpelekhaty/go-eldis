package response

import (
	"context"
	"errors"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

func TestParseNormalizedWithEmptyNode(t *testing.T) {
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

	t.Run("ParseWasteWaterNormalizedWithContext.WithEmptyNode", func(t *testing.T) {
		_, err := ParseWasteWaterNormalizedWithContext(ctx, body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("ParseWasteWaterNormalized.WithEmptyNode", func(t *testing.T) {
		_, err := ParseWasteWaterNormalized(body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("ParseColdWaterNormalizedWithContext.WithEmptyNode", func(t *testing.T) {
		_, err := ParseColdWaterNormalizedWithContext(ctx, body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("ParseColdWaterNormalized.WithEmptyNode", func(t *testing.T) {
		_, err := ParseColdWaterNormalized(body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("ParseHotWaterNormalizedWithContext.WithEmptyNode", func(t *testing.T) {
		_, err := ParseHotWaterNormalizedWithContext(ctx, body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("ParseHotWaterNormalized.WithEmptyNode", func(t *testing.T) {
		_, err := ParseHotWaterNormalized(body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("ParseHeatNormalizedWithContext.WithEmptyNode", func(t *testing.T) {
		_, err := ParseHeatNormalizedWithContext(ctx, body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("ParseHeatNormalized.WithEmptyNode", func(t *testing.T) {
		_, err := ParseHeatNormalized(body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("ParseElectricityNormalizedWithContext.WithEmptyNode", func(t *testing.T) {
		_, err := ParseElectricityNormalizedWithContext(ctx, body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("ParseElectricityNormalized.WithEmptyNode", func(t *testing.T) {
		_, err := ParseElectricityNormalized(body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("ParseElectricityCurrentNormalizedWithContext.WithEmptyNode", func(t *testing.T) {
		_, err := ParseElectricityCurrentNormalizedWithContext(ctx, body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("ParseElectricityCurrentNormalized.WithEmptyNode", func(t *testing.T) {
		_, err := ParseElectricityCurrentNormalized(body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("ParseGasNormalizedWithContext.WithEmptyNode", func(t *testing.T) {
		_, err := ParseGasNormalizedWithContext(ctx, body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("ParseGasNormalized.WithEmptyNode", func(t *testing.T) {
		_, err := ParseGasNormalized(body)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}

func TestParseNormalizedWithEmptyBody(t *testing.T) {
	var body []byte
	ctx := context.TODO()

	t.Run("ParseWasteWaterNormalizedWithContext.WithEmptyBody", func(t *testing.T) {
		_, err := ParseWasteWaterNormalizedWithContext(ctx, body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("ParseWasteWaterNormalized.WithEmptyBody", func(t *testing.T) {
		_, err := ParseWasteWaterNormalized(body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("ParseColdWaterNormalizedWithContext.WithEmptyBody", func(t *testing.T) {
		_, err := ParseColdWaterNormalizedWithContext(ctx, body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("ParseColdWaterNormalized.WithEmptyBody", func(t *testing.T) {
		_, err := ParseColdWaterNormalized(body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("ParseHotWaterNormalizedWithContext.WithEmptyBody", func(t *testing.T) {
		_, err := ParseHotWaterNormalizedWithContext(ctx, body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("ParseHotWaterNormalized.WithEmptyBody", func(t *testing.T) {
		_, err := ParseHotWaterNormalized(body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("ParseHeatNormalizedWithContext.WithEmptyBody", func(t *testing.T) {
		_, err := ParseHeatNormalizedWithContext(ctx, body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("ParseHeatNormalized.WithEmptyBody", func(t *testing.T) {
		_, err := ParseHeatNormalized(body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("ParseElectricityNormalizedWithContext.WithEmptyBody", func(t *testing.T) {
		_, err := ParseElectricityNormalizedWithContext(ctx, body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("ParseElectricityNormalized.WithEmptyBody", func(t *testing.T) {
		_, err := ParseElectricityNormalized(body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("ParseElectricityCurrentNormalizedWithContext.WithEmptyBody", func(t *testing.T) {
		_, err := ParseElectricityCurrentNormalizedWithContext(ctx, body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("ParseElectricityCurrentNormalized.WithEmptyBody", func(t *testing.T) {
		_, err := ParseElectricityCurrentNormalized(body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("ParseGasNormalizedWithContext.WithEmptyBody", func(t *testing.T) {
		_, err := ParseGasNormalizedWithContext(ctx, body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("ParseGasNormalized.WithEmptyBody", func(t *testing.T) {
		_, err := ParseGasNormalized(body)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
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
			t.Errorf("%s: %d item(s), but %d item(s) expected (broken test?)", test.path, count, test.count)
		}
	}
}

func TestParseHotWaterNormalizedWithCancel(t *testing.T) {
	filename := "/testdata/dataNormalizedResponse.json"
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

	items, err := ParseHotWaterNormalizedWithContext(ctx, body)

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

func TestParseHotWaterNormalized(t *testing.T) {
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

	for _, test := range cases {
		count := 0

		path := filepath.Join(filepath.Dir(file), test.path)

		body, err := ioutil.ReadFile(path)

		if err != nil {
			t.Fatalf("%s: %q", test.path, err)
		}

		items, err := ParseHotWaterNormalized(body)

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

func TestParseHeatNormalizedWithContext(t *testing.T) {
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

		items, err := ParseHeatNormalizedWithContext(ctx, body)

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

func TestParseHeatNormalizedWithCancel(t *testing.T) {
	filename := "/testdata/dataNormalizedResponse.json"
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

	items, err := ParseHeatNormalizedWithContext(ctx, body)

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

func TestParseHeatNormalized(t *testing.T) {
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

	for _, test := range cases {
		count := 0

		path := filepath.Join(filepath.Dir(file), test.path)

		body, err := ioutil.ReadFile(path)

		if err != nil {
			t.Fatalf("%s: %q", test.path, err)
		}

		items, err := ParseHeatNormalized(body)

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
