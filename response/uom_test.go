package response

import (
	"context"
	"errors"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

func TestParseUoMGroupsWithContext(t *testing.T) {
	var cases = [1]struct {
		path  string
		count int
	}{
		{path: "/testdata/uomlistResponse.json", count: 20},
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

		items, err := ParseUoMGroupsWithContext(ctx, body)

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
			t.Errorf("%s: %d group(s), but %d group(s) expected (broken test?)", test.path, count, test.count)
		}
	}
}

func TestParseUoMGroupsWithCancel(t *testing.T) {
	filename := "/testdata/uomlistResponse.json"
	const expectedCount = 20

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

	items, err := ParseUoMGroupsWithContext(ctx, body)

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
