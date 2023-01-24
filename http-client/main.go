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
var http2Only, http3Only bool

var ResultHTTP2, ResultHTTP3 models.Result

func init() {
	flag.Int64Var(&count, "count", 100, "Number of times to run the test")
	flag.StringVar(&format, "format", "json", "output format (json or csv)")
	flag.BoolVar(&http2Only, "http2", false, "only http2 test")
	flag.BoolVar(&http3Only, "http3", false, "only http3 test")
	verbose := flag.Bool("verbose", false, "display verbose output")
	flag.Parse()

	if *verbose {
		output = os.Stdout
	} else {
		output = io.Discard
	}
}

func init() {
	ResultHTTP2.Protocol = "http2"
	ResultHTTP3.Protocol = "http3"
}

func main() {
	shouldWriteHeader := false

	for i := int64(0); i < count; i++ {
		runTest(http2Only, http3Only)
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

func runTest(http2Only, http3Only bool) {
	if http2Only {
		conn2 := connector.Http2(output)
		ResultHTTP2.TimeMicroSeconds = append(ResultHTTP2.TimeMicroSeconds, conn2)
		return
	}
	if http3Only {
		conn3 := connector.Http3(output)
		ResultHTTP3.TimeMicroSeconds = append(ResultHTTP3.TimeMicroSeconds, conn3)
		return
	}

	conn2 := connector.Http2(output)
	ResultHTTP2.TimeMicroSeconds = append(ResultHTTP2.TimeMicroSeconds, conn2)
	conn3 := connector.Http3(output)
	ResultHTTP3.TimeMicroSeconds = append(ResultHTTP3.TimeMicroSeconds, conn3)
}
