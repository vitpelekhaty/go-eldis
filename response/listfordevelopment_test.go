package response

import (
	"context"
	"errors"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

func TestParseRegPoints(t *testing.T) {
	var cases = [1]struct {
		path  string
		count int
	}{
		{path: "/testdata/listForDevelopmentResponse.json", count: 6},
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

		items := ParseRegPoints(body)

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

func TestParseRegPointsWithContext(t *testing.T) {
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	items := ParseRegPointsWithContext(ctx, body)

	for item := range items {
		if item.Err != nil {
			t.Errorf("%s: %q", filename, err)
		}

		cancel()
	}
}
