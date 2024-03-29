package response

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

func TestParseCustomResponse(t *testing.T) {
	var cases = [8]string{
		"/testdata/loginResponse.json",
		"/testdata/logoutResponse.json",
		"/testdata/listForDevelopmentResponse.json",
		"/testdata/uomlistResponse.json",
		"/testdata/dataNormalizedResponse.json",
		"/testdata/dataNormalizedEmptyResponse.json",
		"/testdata/rawDataResponse.json",
		"/testdata/empty.json",
	}

	_, file, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal(errors.New("runtime.Caller error"))
	}

	for _, test := range cases {
		path := filepath.Join(filepath.Dir(file), test)

		body, err := ioutil.ReadFile(path)

		if err != nil {
			t.Fatalf("%s: %q", test, err)
		}

		resp, err := Parse(body)

		if err != nil {
			t.Errorf("%s: %q", test, err)
		}

		if len(resp.Messages()) > 1 {
			t.Errorf("%s: invalid response format", test)
		}
	}
}

func TestParseEmptyCustomResponse(t *testing.T) {
	var body []byte

	_, err := Parse(body)

	if err != errEmptyBody {
		t.Fatal("errEmptyBody error expected")
	}
}

func TestParseEmptyCustomResponse2(t *testing.T) {
	var body = []byte("{}")

	r, err := Parse(body)

	if err != nil {
		t.Fatalf(`Parse("{}" error: %q)`, err)
	}

	if len(r.Messages()) != 0 {
		t.Fatalf("Fail: 0 messages expected, but have %d message(s)", len(r.Messages()))
	}
}
