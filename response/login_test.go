package response

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

func TestParseLoginResponse_Result(t *testing.T) {
	var cases = [1]struct {
		path       string
		wantResult bool
	}{
		{path: "/testdata/loginResponse.json", wantResult: true},
	}

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	for _, test := range cases {
		path := filepath.Join(filepath.Dir(file), test.path)

		body, err := ioutil.ReadFile(path)

		if err != nil {
			t.Fatalf("%s: %q", test.path, err)
		}

		resp, err := ParseLoginResponse(body)

		if err != nil {
			t.Errorf("%s: %q", test.path, err)
		}

		if resp.Result() != test.wantResult {
			t.Errorf("%s: broken test (?)", test.path)
		}
	}
}

func TestParseLoginResponse_Messages(t *testing.T) {
	var paths = [1]string{"/testdata/loginResponse.json"}

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	for _, p := range paths {
		path := filepath.Join(filepath.Dir(file), p)

		body, err := ioutil.ReadFile(path)

		if err != nil {
			t.Fatalf("%s: %q", p, err)
		}

		resp, err := ParseLoginResponse(body)

		if err != nil {
			t.Errorf("%s: %q", p, err)
		}

		if len(resp.Messages()) != 1 {
			t.Errorf("%s: invalid response format", p)
		}
	}
}

func TestParseLoginResponseWithEmptyBody(t *testing.T) {
	var body []byte
	_, err := ParseLoginResponse(body)

	if err != errEmptyBody {
		t.Error("errEmptyBody error expected")
	}
}
