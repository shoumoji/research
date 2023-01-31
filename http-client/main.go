package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/shoumoji/research/http3-client/connector"
	"github.com/shoumoji/research/http3-client/models"
	"github.com/shoumoji/research/http3-client/view"
)

var (
	count              int64
	format             string
	output             io.Writer
	http2Url, http3Url string

	ResultHTTP2, ResultHTTP3 models.Result
)

func init() {
	flag.Int64Var(&count, "count", 100, "Number of times to run the test")
	flag.StringVar(&format, "format", "json", "output format (json or csv)")
	flag.StringVar(&http2Url, "http2", "", "http2 url")
	flag.StringVar(&http3Url, "http3", "", "http3 url")

	verbose := flag.Bool("verbose", false, "display verbose output")
	flag.Parse()

	if *verbose {
		output = os.Stdout
	} else {
		output = io.Discard
	}

	if count < int64(1) {
		log.Fatal("count must be greater than 0")
	}

	if http2Url == "" && http3Url == "" {
		log.Fatal("url is required")
	}
}

func init() {
	ResultHTTP2.Protocol = "http2"
	ResultHTTP3.Protocol = "http3"

	ResultHTTP2.Count = count
	ResultHTTP3.Count = count
}

func main() {
	shouldWriteHeader := true

	for i := int64(0); i < count; i++ {
		if http2Url != "" {
			if runTestHTTP2(http2Url) {
				ResultHTTP2.ConnectSuccessCount++
			}
		}

		if http3Url != "" {
			if runTestHTTP3(http3Url) {
				ResultHTTP3.ConnectSuccessCount++
			}
		}
	}

	switch format {
	case "json":
		view.OutputJSON(&ResultHTTP2, &ResultHTTP3)
	case "csv":
		view.OutputCSV(shouldWriteHeader, count, &ResultHTTP2, &ResultHTTP3)
	default:
		log.Fatal("invalid format")
	}

}

func runTestHTTP2(http2Url string) (successConnect bool) {
	conn2, err := connector.Http2(http2Url, output)
	if err != nil {
		return false
	}
	ResultHTTP2.TimeMicroSeconds = append(ResultHTTP2.TimeMicroSeconds, conn2)

	return true
}

func runTestHTTP3(http3Url string) (successConnect bool) {
	conn3, err := connector.Http3(http3Url, output)
	if err != nil {
		return false
	}
	ResultHTTP3.TimeMicroSeconds = append(ResultHTTP3.TimeMicroSeconds, conn3)

	return true
}
