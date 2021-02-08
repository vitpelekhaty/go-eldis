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

func TestGetRawDataBytes(t *testing.T) {
	const filename = "/testdata/rawDataResponse.json"

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	path := filepath.Join(filepath.Dir(file), filename)

	body, err := ioutil.ReadFile(path)

	if err != nil {
		t.Errorf("%s: %q", filename, err)
	}

	b, err := getBytes(body, rawDataPath)

	if err != nil {
		t.Fatalf("%s: %q", filename, err)
	}

	if len(b) == 0 {
		t.Fail()
	}
}

func TestGetBytesWithEmptyBody(t *testing.T) {
	var body []byte

	t.Run(fmt.Sprintf("empty.%s", normalizedWasteWaterPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedWasteWaterPath)

		if err != errEmptyBody {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty.%s", normalizedColdWaterPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedWasteWaterPath)

		if err != errEmptyBody {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty.%s", normalizedHotWaterPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedWasteWaterPath)

		if err != errEmptyBody {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty.%s", normalizedHeatPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedHeatPath)

		if err != errEmptyBody {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty.%s", normalizedElectricityPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedElectricityPath)

		if err != errEmptyBody {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty.%s", normalizedElectricityCurrentPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedElectricityCurrentPath)

		if err != errEmptyBody {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty.%s", normalizedGasPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedGasPath)

		if err != errEmptyBody {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty.%s", listForDevelopmentPath), func(t *testing.T) {
		_, err := getBytes(body, listForDevelopmentPath)

		if err != errEmptyBody {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("empty.%s", uomPath), func(t *testing.T) {
		_, err := getBytes(body, uomPath)

		if err != errEmptyBody {
			t.Fail()
		}
	})
}

func TestGetBytesWithEmptyResponse(t *testing.T) {
	var body = []byte("{}")

	t.Run(fmt.Sprintf("emptyResponse.%s", normalizedWasteWaterPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedWasteWaterPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("emptyResponse.%s", normalizedColdWaterPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedWasteWaterPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("emptyResponse.%s", normalizedHotWaterPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedWasteWaterPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("emptyResponse.%s", normalizedHeatPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedHeatPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("emptyResponse.%s", normalizedElectricityPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedElectricityPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("emptyResponse.%s", normalizedElectricityCurrentPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedElectricityCurrentPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("emptyResponse.%s", normalizedGasPath), func(t *testing.T) {
		_, err := getBytes(body, normalizedGasPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("emptyResponse.%s", listForDevelopmentPath), func(t *testing.T) {
		_, err := getBytes(body, listForDevelopmentPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run(fmt.Sprintf("emptyResponse.%s", uomPath), func(t *testing.T) {
		_, err := getBytes(body, uomPath)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}

func TestGet_SectionListForDevelopment(t *testing.T) {
	filename := "/testdata/listForDevelopmentResponse.json"

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	path := filepath.Join(filepath.Dir(file), filename)

	body, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatalf("%s: %q", filename, err)
	}

	t.Run("Get.SectionListForDevelopment", func(t *testing.T) {
		b, err := Get(SectionListForDevelopment, body)

		if err != nil {
			t.Errorf("%s: %q", filename, err)
		}

		if len(b) == 0 {
			t.Fail()
		}
	})

	t.Run("Get.SectionListForDevelopment.WithEmptyBody", func(t *testing.T) {
		var emptyBody []byte

		_, err := Get(SectionListForDevelopment, emptyBody)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("Get.SectionListForDevelopment.WithEmptyResponse", func(t *testing.T) {
		var emptyBody = []byte("{}")

		_, err := Get(SectionListForDevelopment, emptyBody)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}

func TestGet_SectionNormalizedWasteWaterData(t *testing.T) {
	t.Run("Get.SectionNormalizedWasteWaterData.WithEmptyBody", func(t *testing.T) {
		var emptyBody []byte

		_, err := Get(SectionNormalizedWasteWaterData, emptyBody)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("Get.SectionNormalizedWasteWaterData.WithEmptyResponse", func(t *testing.T) {
		var emptyBody = []byte("{}")

		_, err := Get(SectionNormalizedWasteWaterData, emptyBody)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}

func TestGet_SectionNormalizedColdWaterData(t *testing.T) {
	t.Run("Get.SectionNormalizedColdWaterData.WithEmptyBody", func(t *testing.T) {
		var emptyBody []byte

		_, err := Get(SectionNormalizedColdWaterData, emptyBody)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("Get.SectionNormalizedColdWaterData.WithEmptyResponse", func(t *testing.T) {
		var emptyBody = []byte("{}")

		_, err := Get(SectionNormalizedColdWaterData, emptyBody)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}

func TestGet_SectionNormalizedHotWaterData(t *testing.T) {
	filename := "/testdata/dataNormalizedResponse.json"

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	path := filepath.Join(filepath.Dir(file), filename)

	body, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatalf("%s: %q", filename, err)
	}

	t.Run("Get.SectionNormalizedHotWaterData", func(t *testing.T) {
		b, err := Get(SectionNormalizedHotWaterData, body)

		if err != nil {
			t.Errorf("%s: %q", filename, err)
		}

		if len(b) == 0 {
			t.Fail()
		}
	})

	t.Run("Get.SectionNormalizedHotWaterData.WithEmptyBody", func(t *testing.T) {
		var emptyBody []byte

		_, err := Get(SectionNormalizedHotWaterData, emptyBody)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("Get.SectionNormalizedHotWaterData.WithEmptyResponse", func(t *testing.T) {
		var emptyBody = []byte("{}")

		_, err := Get(SectionNormalizedColdWaterData, emptyBody)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}

func TestGet_SectionNormalizedHeatData(t *testing.T) {
	filename := "/testdata/dataNormalizedResponse.json"

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	path := filepath.Join(filepath.Dir(file), filename)

	body, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatalf("%s: %q", filename, err)
	}

	t.Run("Get.SectionNormalizedHeatData", func(t *testing.T) {
		b, err := Get(SectionNormalizedHeatData, body)

		if err != nil {
			t.Errorf("%s: %q", filename, err)
		}

		if len(b) == 0 {
			t.Fail()
		}
	})

	t.Run("Get.SectionNormalizedHeatData.WithEmptyBody", func(t *testing.T) {
		var emptyBody []byte

		_, err := Get(SectionNormalizedHeatData, emptyBody)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("Get.SectionNormalizedHeatData.WithEmptyResponse", func(t *testing.T) {
		var emptyBody = []byte("{}")

		_, err := Get(SectionNormalizedHeatData, emptyBody)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}

func TestGet_SectionNormalizedElectricityData(t *testing.T) {
	t.Run("Get.SectionNormalizedElectricityData.WithEmptyBody", func(t *testing.T) {
		var emptyBody []byte

		_, err := Get(SectionNormalizedElectricityData, emptyBody)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("Get.SectionNormalizedElectricityData.WithEmptyResponse", func(t *testing.T) {
		var emptyBody = []byte("{}")

		_, err := Get(SectionNormalizedElectricityData, emptyBody)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})

	t.Run("Get.SectionNormalizedElectricityData.WithEmptyResponse", func(t *testing.T) {
		var emptyBody = []byte("{}")

		_, err := Get(SectionNormalizedElectricityData, emptyBody)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}

func TestGet_SectionNormalizedElectricityCurrentData(t *testing.T) {
	t.Run("Get.SectionNormalizedElectricityCurrentData.WithEmptyBody", func(t *testing.T) {
		var emptyBody []byte

		_, err := Get(SectionNormalizedElectricityCurrentData, emptyBody)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("Get.SectionNormalizedElectricityCurrentData.WithEmptyResponse", func(t *testing.T) {
		var emptyBody = []byte("{}")

		_, err := Get(SectionNormalizedElectricityCurrentData, emptyBody)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}

func TestGet_SectionNormalizedGasData(t *testing.T) {
	t.Run("Get.SectionNormalizedGasData.WithEmptyBody", func(t *testing.T) {
		var emptyBody []byte

		_, err := Get(SectionNormalizedGasData, emptyBody)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("Get.SectionNormalizedGasData.WithEmptyResponse", func(t *testing.T) {
		var emptyBody = []byte("{}")

		_, err := Get(SectionNormalizedGasData, emptyBody)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}

func TestGet_SectionUoM(t *testing.T) {
	filename := "/testdata/uomlistResponse.json"

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	path := filepath.Join(filepath.Dir(file), filename)

	body, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatalf("%s: %q", filename, err)
	}

	t.Run("Get.SectionUoM", func(t *testing.T) {
		b, err := Get(SectionUoM, body)

		if err != nil {
			t.Errorf("%s: %q", filename, err)
		}

		if len(b) == 0 {
			t.Fail()
		}
	})

	t.Run("Get.SectionUoM.WithEmptyBody", func(t *testing.T) {
		var emptyBody []byte

		_, err := Get(SectionUoM, emptyBody)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("Get.SectionUoM.WithEmptyResponse", func(t *testing.T) {
		var emptyBody = []byte("{}")

		_, err := Get(SectionUoM, emptyBody)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}

func TestGet_SectionRawData(t *testing.T) {
	filename := "/testdata/rawDataResponse.json"

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	path := filepath.Join(filepath.Dir(file), filename)

	body, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatalf("%s: %q", filename, err)
	}

	t.Run("Get.SectionRawData", func(t *testing.T) {
		b, err := Get(SectionRawData, body)

		if err != nil {
			t.Errorf("%s: %q", filename, err)
		}

		if len(b) == 0 {
			t.Fail()
		}
	})

	t.Run("Get.SectionRawData.WithEmptyBody", func(t *testing.T) {
		var emptyBody []byte

		_, err := Get(SectionRawData, emptyBody)

		if err != errEmptyBody {
			t.Error("errEmptyBody error expected")
		}
	})

	t.Run("Get.SectionRawData.WithEmptyResponse", func(t *testing.T) {
		var emptyBody = []byte("{}")

		_, err := Get(SectionRawData, emptyBody)

		if _, ok := err.(*PathError); !ok {
			t.Fail()
		}
	})
}

func TestGet_UnknownSection(t *testing.T) {
	var body []byte

	_, err := Get(255, body)

	if err != errUnavailableForSection {
		t.Error("errUnavailableForSection error expected")
	}
}
