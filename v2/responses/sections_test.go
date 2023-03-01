package responses

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	pathNormalizedEmptyResponse    = "../testdata/dataNormalizedEmptyResponse.json"
	pathNormalizedResponse         = "../testdata/dataNormalizedResponse.json"
	pathListForDevelopmentResponse = "../testdata/listForDevelopmentResponse.json"
	pathRawDataResponse            = "../testdata/rawDataResponse.json"
)

func TestExtract(t *testing.T) {
	t.Run("pathNormalizedEmptyResponse", func(t *testing.T) {
		path := fullpath(pathNormalizedEmptyResponse)

		b, err := os.ReadFile(path)
		require.NoError(t, err)

		sb := Extract(SectionNormalizedHotWater, bytes.NewBuffer(b))
		require.Equal(t, 0, len(sb))
	})

	t.Run("pathNormalizedResponse", func(t *testing.T) {
		path := fullpath(pathNormalizedResponse)

		b, err := os.ReadFile(path)
		require.NoError(t, err)

		sb := Extract(SectionNormalizedHotWater, bytes.NewBuffer(b))
		require.NotEqual(t, 0, len(sb))
	})

	t.Run("pathListForDevelopmentResponse", func(t *testing.T) {
		path := fullpath(pathListForDevelopmentResponse)

		b, err := os.ReadFile(path)
		require.NoError(t, err)

		sb := Extract(SectionListForDevelopment, bytes.NewBuffer(b))
		require.NotEqual(t, 0, len(sb))
	})

	t.Run("pathRawDataResponse", func(t *testing.T) {
		path := fullpath(pathRawDataResponse)

		b, err := os.ReadFile(path)
		require.NoError(t, err)

		sb := Extract(SectionRaw, bytes.NewBuffer(b))
		require.NotEqual(t, 0, len(sb))
	})
}

func fullpath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	fn, _ := filepath.Abs(path)

	return fn
}
