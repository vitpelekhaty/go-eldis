package eldis

import (
	"bytes"
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/vitpelekhaty/go-eldis/v2/responses"
	"github.com/vitpelekhaty/go-eldis/v2/responses/points"
)

var (
	rawURL      = os.Getenv("ELDIS_API_URL")
	username    = os.Getenv("ELDIS_API_USERNAME")
	password    = os.Getenv("ELDIS_API_PASSWORD")
	accessToken = os.Getenv("ELDIS_API_ACCESS_TOKEN")
	pointID     = os.Getenv("ELDIS_POINT_ID")
)

func TestConnect(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.TODO(), time.Second*30)

	defer func() {
		cancelFunc()
	}()

	conn, err := Connect(ctx, rawURL, Credentials{Username: username, Password: password, AccessToken: accessToken})
	require.NoError(t, err)

	err = conn.Close(ctx)
	require.NoError(t, err)
}

func TestConnection_ListForDevelopment(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.TODO(), time.Second*30)

	defer func() {
		cancelFunc()
	}()

	conn, err := Connect(ctx, rawURL, Credentials{Username: username, Password: password, AccessToken: accessToken})
	require.NoError(t, err)

	defer func() {
		err := conn.Close(ctx)
		require.NoError(t, err)
	}()

	b, err := conn.ListForDevelopment(ctx)
	require.NoError(t, err)

	sb := responses.Extract(responses.SectionListForDevelopment, bytes.NewBuffer(b))
	require.NotEqual(t, 0, len(sb))

	items, err := points.Parse(context.Background(), bytes.NewReader(sb))
	require.NoError(t, err)

	var count int

	for item := range items {
		require.NoError(t, item.E, "item", item)
		count++
	}

	require.NotEqual(t, 0, count)
}

func TestConnection_NormalizedReadings(t *testing.T) {
	to := time.Now()
	from := to.Add(-time.Hour * 24 * 7)

	ctx, cancelFunc := context.WithTimeout(context.TODO(), time.Second*30)

	defer func() {
		cancelFunc()
	}()

	conn, err := Connect(ctx, rawURL, Credentials{Username: username, Password: password, AccessToken: accessToken})
	require.NoError(t, err)

	defer func() {
		err := conn.Close(ctx)
		require.NoError(t, err)
	}()

	t.Run("hour_archive", func(t *testing.T) {
		_, err = conn.NormalizedReadings(ctx, pointID, HourArchive, from, to, Date)
		require.NoError(t, err)
	})

	t.Run("daily_archive", func(t *testing.T) {
		_, err = conn.NormalizedReadings(ctx, pointID, DailyArchive, from, to, Date)
		require.NoError(t, err)
	})
}

func TestConnection_RawReadings(t *testing.T) {
	to := time.Now()
	from := to.Add(-time.Hour * 24 * 7)

	ctx, cancelFunc := context.WithTimeout(context.TODO(), time.Second*30)

	defer func() {
		cancelFunc()
	}()

	conn, err := Connect(ctx, rawURL, Credentials{Username: username, Password: password, AccessToken: accessToken})
	require.NoError(t, err)

	defer func() {
		err := conn.Close(ctx)
		require.NoError(t, err)
	}()

	t.Run("hour_archive", func(t *testing.T) {
		_, err = conn.RawReadings(ctx, pointID, HourArchive, from, to)
		require.NoError(t, err)
	})

	t.Run("daily_archive", func(t *testing.T) {
		_, err = conn.RawReadings(ctx, pointID, DailyArchive, from, to)
		require.NoError(t, err)
	})
}
