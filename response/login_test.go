package response

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

func TestParseLoginResponse(t *testing.T) {
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
