package connector

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/lucas-clemente/quic-go/http3"
)

func Http2(output io.Writer) int64 {
	http2Before := time.Now()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			ClientSessionCache:          tls.NewLRUClientSessionCache(0),
			MinVersion:                  tls.VersionTLS13,
			MaxVersion:                  tls.VersionTLS13,
			CurvePreferences:            []tls.CurveID{},
			DynamicRecordSizingDisabled: false,
			Renegotiation:               0,
			KeyLogWriter:                output,
			InsecureSkipVerify:          true,
		},
	}
	client := &http.Client{
		Transport: tr,
	}
	req, _ := http.NewRequest("GET", "https://localhost:8081", nil)

	resp, err := client.Transport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Fatal(err)
	}

	return time.Since(http2Before).Microseconds()
}

func Http3(output io.Writer) int64 {
	http3Before := time.Now()

	r := http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			ClientSessionCache:          tls.NewLRUClientSessionCache(0),
			MinVersion:                  tls.VersionTLS13,
			MaxVersion:                  tls.VersionTLS13,
			CurvePreferences:            []tls.CurveID{},
			DynamicRecordSizingDisabled: false,
			Renegotiation:               0,
			KeyLogWriter:                output,
			InsecureSkipVerify:          true,
		},
	}
	req, _ := http.NewRequest("GET", "https://localhost:8081", nil)

	resp, err := r.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(io.Discard, resp.Body)

	return time.Since(http3Before).Microseconds()
}
