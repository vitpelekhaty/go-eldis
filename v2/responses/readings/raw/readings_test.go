package raw

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vitpelekhaty/go-eldis/v2/responses"
)

const pathRawDataResponse = "../../../testdata/rawDataResponse.json"

func TestParse(t *testing.T) {
	path := fullpath(pathRawDataResponse)

	b, err := os.ReadFile(path)
	require.NoError(t, err)

	sb := responses.Extract(responses.SectionRaw, bytes.NewBuffer(b))
	require.NotEqual(t, 0, len(sb))

	items, err := Parse(context.Background(), bytes.NewReader(sb))
	require.NoError(t, err)

	var count int

	for item := range items {
		require.NoError(t, item.E, "item", item)
		count++
	}

	require.NotEqual(t, 0, count)
}

func fullpath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	fn, _ := filepath.Abs(path)

	return fn
}
