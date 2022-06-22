package fcm

import (
	"flag"
	"net/http"
	"os"
	"testing"
	"time"
)

var httpClient = &http.Client{Transport: &http.Transport{TLSHandshakeTimeout: 60 * time.Second}}

func TestMain(m *testing.M) {
	defer httpClient.CloseIdleConnections()
	flag.Parse()

	os.Exit(m.Run())
}
