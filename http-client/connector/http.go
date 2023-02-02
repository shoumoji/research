package connector

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"github.com/lucas-clemente/quic-go/http3"
)

func Http2(url string, output io.Writer) (int64, error) {
	http2Before := time.Now()

	tr := &http.Transport{
		ForceAttemptHTTP2:  true,
		DisableCompression: true,
		TLSClientConfig: &tls.Config{
			ClientSessionCache:          tls.NewLRUClientSessionCache(0),
			MinVersion:                  tls.VersionTLS13,
			MaxVersion:                  tls.VersionTLS13,
			DynamicRecordSizingDisabled: false,
			Renegotiation:               0,
			KeyLogWriter:                output,
			InsecureSkipVerify:          true,
			// x25519
			CurvePreferences: []tls.CurveID{29},
			// TLS_AES_128_GCM_SHA256
			CipherSuites: []uint16{0x1301},
		},
	}
	client := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1, err
	}

	resp, err := client.Transport.RoundTrip(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		return -1, err
	}

	return time.Since(http2Before).Microseconds(), nil
}

func Http3(url string, output io.Writer) (int64, error) {
	http3Before := time.Now()

	r := http3.RoundTripper{
		DisableCompression: true,
		TLSClientConfig: &tls.Config{
			ClientSessionCache:          tls.NewLRUClientSessionCache(0),
			MinVersion:                  tls.VersionTLS13,
			MaxVersion:                  tls.VersionTLS13,
			DynamicRecordSizingDisabled: false,
			Renegotiation:               0,
			KeyLogWriter:                output,
			InsecureSkipVerify:          true,
			// x25519
			CurvePreferences: []tls.CurveID{29},
			// TLS_AES_128_GCM_SHA256
			CipherSuites: []uint16{0x1301},
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1, err
	}

	resp, err := r.RoundTrip(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		return -1, err
	}

	return time.Since(http3Before).Microseconds(), nil
}
