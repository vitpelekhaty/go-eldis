package eldis

import (
	"flag"
	"testing"
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
}

const (
	layoutQuery = `02.01.2006 15:04`
)

func TestConnection(t *testing.T) {
}
