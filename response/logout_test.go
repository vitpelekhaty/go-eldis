package response

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

func TestParseLogoutResponse_Result(t *testing.T) {
	var cases = [1]struct {
		path       string
		wantResult bool
	}{
		{path: "/testdata/logoutResponse.json", wantResult: true},
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

		resp, err := ParseLogoutResponse(body)

		if err != nil {
			t.Errorf("%s: %q", test.path, err)
		}

		if resp.Result() != test.wantResult {
			t.Errorf("%s: broken test (?)", test.path)
		}
	}
}

func TestParseLogoutResponse_Messages(t *testing.T) {
	var paths = [1]string{"/testdata/logoutResponse.json"}

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

		resp, err := ParseLogoutResponse(body)

		if err != nil {
			t.Errorf("%s: %q", p, err)
		}

		if len(resp.Messages()) != 1 {
			t.Errorf("%s: invalid response format", p)
		}
	}
}

func TestParseLogoutResponseWithEmptyBody(t *testing.T) {
	var body []byte
	_, err := ParseLogoutResponse(body)

	if err != errEmptyBody {
		t.Error("errEmptyBody error expected")
	}
}
