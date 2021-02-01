package response

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGetBytesForEmptySection(t *testing.T) {
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

	t.Run(fmt.Sprintf("empty.%s", normalizedWasteWaterPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedWasteWaterPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty.%s", normalizedColdWaterPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedWasteWaterPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty.%s", normalizedHotWaterPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedWasteWaterPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty.%s", normalizedHeatPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedHeatPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty.%s", normalizedElectricityPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedElectricityPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty.%s", normalizedElectricityCurrentPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedElectricityCurrentPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty.%s", normalizedGasPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedGasPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}

func TestGetHotWaterNormalizedBytes(t *testing.T) {
	const filename = "/testdata/dataNormalizedResponse.json"

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	path := filepath.Join(filepath.Dir(file), filename)

	body, err := ioutil.ReadFile(path)

	if err != nil {
		t.Errorf("%s: %q", filename, err)
	}

	b, err := getBytes(body, normalizedHotWaterPath)

	if err != nil {
		t.Fatalf("%s: %q", filename, err)
	}

	if len(b) == 0 {
		t.Fail()
	}
}

func TestGetHeatNormalizedBytes(t *testing.T) {
	const filename = "/testdata/dataNormalizedResponse.json"

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	path := filepath.Join(filepath.Dir(file), filename)

	body, err := ioutil.ReadFile(path)

	if err != nil {
		t.Errorf("%s: %q", filename, err)
	}

	b, err := getBytes(body, normalizedHeatPath)

	if err != nil {
		t.Fatalf("%s: %q", filename, err)
	}

	if len(b) == 0 {
		t.Fail()
	}
}

func TestGetListForDevelopmentBytes(t *testing.T) {
	const filename = "/testdata/listForDevelopmentResponse.json"

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	path := filepath.Join(filepath.Dir(file), filename)

	body, err := ioutil.ReadFile(path)

	if err != nil {
		t.Errorf("%s: %q", filename, err)
	}

	b, err := getBytes(body, listForDevelopmentPath)

	if err != nil {
		t.Fatalf("%s: %q", filename, err)
	}

	if len(b) == 0 {
		t.Fail()
	}
}

func TestGetUoMGroupBytes(t *testing.T) {
	const filename = "/testdata/uomlistResponse.json"

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	path := filepath.Join(filepath.Dir(file), filename)

	body, err := ioutil.ReadFile(path)

	if err != nil {
		t.Errorf("%s: %q", filename, err)
	}

	b, err := getBytes(body, uomPath)

	if err != nil {
		t.Fatalf("%s: %q", filename, err)
	}

	if len(b) == 0 {
		t.Fail()
	}
}
