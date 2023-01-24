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

var count int64
var format string
var output io.Writer
var http2Url, http3Url string

var ResultHTTP2, ResultHTTP3 models.Result

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

	if http2Url == "" && http3Url == "" {
		log.Fatal("url is required")
	}
}

func init() {
	ResultHTTP2.Protocol = "http2"
	ResultHTTP3.Protocol = "http3"
}

func main() {
	shouldWriteHeader := false

	for i := int64(0); i < count; i++ {
		err := runTest(http2Url, http3Url)
		if err != nil {
			log.Fatal(err)
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

func runTest(http2Url, http3Url string) error {
	if http2Url != "" {
		conn2, err := connector.Http2(http2Url, output)
		if err != nil {
			return err
		}
		ResultHTTP2.TimeMicroSeconds = append(ResultHTTP2.TimeMicroSeconds, conn2)
	}
	if http3Url != "" {
		conn3, err := connector.Http3(http3Url, output)
		if err != nil {
			return err
		}
		ResultHTTP3.TimeMicroSeconds = append(ResultHTTP3.TimeMicroSeconds, conn3)
	}

	return nil
}
