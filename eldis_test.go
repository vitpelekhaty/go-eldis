// +build integration

package eldis

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/vitpelekhaty/httptracer"
)

var (
	username           string
	password           string
	apiKey             string
	rawURL             string
	insecureSkipVerify bool
	bodies             bool
	tracePath          string
	strTimeout         string
	limit              uint
	strStart           string
	strEnd             string
	strDataArchive     string
	strDateType        string
)

func init() {
	flag.StringVar(&username, "username", "", "username")
	flag.StringVar(&password, "password", "", "password")
	flag.StringVar(&apiKey, "key", "", "API key")
	flag.StringVar(&rawURL, "api-url", "", "API URL")
	flag.BoolVar(&insecureSkipVerify, "insecure-skip-verify", false, "insecure skip verify")
	flag.StringVar(&tracePath, "trace", "", "write trace into path")
	flag.StringVar(&strTimeout, "timeout", "30s", "timeout")
	flag.BoolVar(&bodies, "bodies", false, "write bodies into trace")
	flag.UintVar(&limit, "limit", 0, "limit number of points")
	flag.StringVar(&strStart, "from", "", "a beginning of a measurement period")
	flag.StringVar(&strEnd, "to", "", "end of measurement period")
	flag.StringVar(&strDataArchive, "archive", "HourArchive", "type of archive")
	flag.StringVar(&strDateType, "date", "date", "type of date (for DataNormalized())")
}

/*
const (
	layoutQuery = `02.01.2006 15:04`
)
*/

func TestConnection_RawData(t *testing.T) {
	_, err := url.Parse(rawURL)

	if err != nil {
		t.Fatal(err)
	}
	/*
		start, err := time.Parse(layoutQuery, strStart)

		if err != nil {
			t.Fatal(err)
		}

		end, err := time.Parse(layoutQuery, strEnd)

		if err != nil {
			t.Fatal(err)
		}

		archiveType, err := archive.Parse(strDataArchive)

		if err != nil {
			t.Fatal(err)
		}
	*/
	timeout, err := time.ParseDuration(strTimeout)

	if err != nil {
		t.Fatal(err)
	}

	client := setupHTTPClient(timeout*time.Second, insecureSkipVerify)

	if strings.TrimSpace(tracePath) != "" {
		f, err := os.Create(tracePath)

		if err != nil {
			t.Fatal(err)
		}

		defer func() {
			if _, err := f.WriteString("]"); err != nil {
				t.Error(err)
			}

			if err := f.Close(); err != nil {
				t.Error(err)
			}
		}()

		_, err = f.WriteString("[")

		if err != nil {
			t.Fatal(err)
		}

		emptyTrace := true

		callbackFunc := func(entry *httptracer.Entry) {
			if entry == nil {
				return
			}

			b, err := json.Marshal(entry)

			if err != nil {
				t.Fatal(err)
			}

			if !emptyTrace {
				_, err = f.WriteString(",")

				if err != nil {
					t.Fatal(err)
				}
			}

			_, err = f.Write(b)

			if err != nil {
				t.Fatal(err)
			}

			emptyTrace = false
		}

		client = setupTracer(client, setupTracerOptions(bodies, callbackFunc)...)
	}

	c, err := NewConnection(client)

	if err != nil {
		t.Fatal(err)
	}

	err = c.Open(rawURL, WithAuth(username, password, apiKey))

	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := c.Close(); err != nil {
			t.Error(err)
		}
	}()

	uom, err := c.UOMList()

	if err != nil {
		t.Errorf("UoMList() error: %q", err)
	}

	if len(uom) == 0 {
		t.Error("UoMList() error: empty body")
	}

	p, err := c.ListForDevelopment()

	if err != nil {
		t.Fatalf("ListForDevelopment() error: %q", err)
	}

	if len(p) == 0 {
		t.Fatal("ListForDevelopment() error: empty body")
	}

	/*
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		points, err := response.ParseRegPointsWithContext(ctx, p)

		if err != nil {
			t.Fatalf("ParseRawDataWithContext() error: %q", err)
		}

		var pointCount int

		for point := range points {
			if point.Err != nil {
				t.Error(point.Err)

				pointCount++

				continue
			}

			if limit > 0 && pointCount > int(limit) {
				cancel()
			}

			regPoint := point.RegPoint

			t.Logf("RawData() for point %s (sensor %s %s)", regPoint.ID, regPoint.DeviceName, regPoint.SN)

			d, err := c.RawData(point.RegPoint.ID, archiveType, RequestTime(start), RequestTime(end))

			if err != nil {
				t.Errorf("RawData() error: %q", err)

				pointCount++

				continue
			}

			if len(d) == 0 {
				t.Error("RawData() error: empty body")

				pointCount++

				continue
			}

			pointCount++
		}

	*/
}

/*
func TestConnection_DataNormalized(t *testing.T) {
	_, err := url.Parse(rawURL)

	if err != nil {
		t.Fatal(err)
	}

	start, err := time.Parse(layoutQuery, strStart)

	if err != nil {
		t.Fatal(err)
	}

	end, err := time.Parse(layoutQuery, strEnd)

	if err != nil {
		t.Fatal(err)
	}

	archiveType, err := archive.Parse(strDataArchive)

	if err != nil {
		t.Fatal(err)
	}

	dateType, err := date.Parse(strDateType)

	if err != nil {
		t.Fatal(err)
	}

	timeout, err := time.ParseDuration(strTimeout)

	if err != nil {
		t.Fatal(err)
	}

	client := setupHTTPClient(timeout*time.Second, insecureSkipVerify)

	if strings.TrimSpace(tracePath) != "" {
		f, err := os.Create(tracePath)

		if err != nil {
			t.Fatal(err)
		}

		defer func() {
			if _, err := f.WriteString("]"); err != nil {
				t.Error(err)
			}

			if err := f.Close(); err != nil {
				t.Error(err)
			}
		}()

		_, err = f.WriteString("[")

		if err != nil {
			t.Fatal(err)
		}

		emptyTrace := true

		callbackFunc := func(entry *httptracer.Entry) {
			if entry == nil {
				return
			}

			b, err := json.Marshal(entry)

			if err != nil {
				t.Fatal(err)
			}

			if !emptyTrace {
				_, err = f.WriteString(",")

				if err != nil {
					t.Fatal(err)
				}
			}

			_, err = f.Write(b)

			if err != nil {
				t.Fatal(err)
			}

			emptyTrace = false
		}

		client = setupTracer(client, setupTracerOptions(bodies, callbackFunc)...)
	}

	c, err := NewConnection(client)

	if err != nil {
		t.Fatal(err)
	}

	err = c.Open(rawURL, WithAuth(username, password, apiKey))

	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := c.Close(); err != nil {
			t.Error(err)
		}
	}()

	uom, err := c.UOMList()

	if err != nil {
		t.Errorf("UoMList() error: %q", err)
	}

	if len(uom) == 0 {
		t.Error("UoMList() error: empty body")
	}

	p, err := c.ListForDevelopment()

	if err != nil {
		t.Fatalf("ListForDevelopment() error: %q", err)
	}

	if len(p) == 0 {
		t.Fatal("ListForDevelopment() error: empty body")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	points, err := response.ParseRegPointsWithContext(ctx, p)

	if err != nil {
		t.Fatalf("ParseRawDataWithContext() error: %q", err)
	}

	var pointCount int

	for point := range points {
		if point.Err != nil {
			t.Error(point.Err)

			pointCount++

			continue
		}

		if limit > 0 && pointCount > int(limit) {
			cancel()
		}

		regPoint := point.RegPoint

		t.Logf("RawData() for point %s (sensor %s %s)", regPoint.ID, regPoint.DeviceName, regPoint.SN)

		d, err := c.DataNormalized(point.RegPoint.ID, archiveType, RequestTime(start), RequestTime(end), dateType)

		if err != nil {
			t.Errorf("RawData() error: %q", err)

			pointCount++

			continue
		}

		if len(d) == 0 {
			t.Error("DataNormalized() error: empty body")

			pointCount++

			continue
		}

		pointCount++
	}
}
*/

func setupHTTPClient(timeout time.Duration, insecureSkipVerify bool) *http.Client {
	client := &http.Client{
		Timeout: timeout,
	}

	if insecureSkipVerify {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client.Transport = transport
	}

	return client
}

func setupTracer(client *http.Client, options ...httptracer.Option) *http.Client {
	return httptracer.Trace(client, options...)
}

func setupTracerOptions(withBodies bool, withCallback httptracer.CallbackFunc) []httptracer.Option {
	options := make([]httptracer.Option, 0)

	if withBodies {
		options = append(options, httptracer.WithBodies(withBodies))
	}

	if withCallback != nil {
		options = append(options, httptracer.WithCallback(withCallback))
	}

	return options
}
