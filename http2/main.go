package main

import (
	"crypto/tls"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const size1mb = 1024 * 1024
const size10mb = 1024 * 1024 * 10
const size100mb = 1024 * 1024 * 100
const size1000mb = 1024 * 1024 * 1000

var data1mb, data10mb, data100mb, data1000mb []byte

func init() {
	rand.Seed(time.Now().UnixNano())

	data1mb = make([]byte, size1mb)
	if _, err := rand.Read(data1mb); err != nil {
		log.Fatal(err)
	}

	data10mb = make([]byte, size10mb)
	if _, err := rand.Read(data10mb); err != nil {
		log.Fatal(err)
	}

	data100mb = make([]byte, size100mb)
	if _, err := rand.Read(data100mb); err != nil {
		log.Fatal(err)
	}

	data1000mb = make([]byte, size1000mb)
	if _, err := rand.Read(data1000mb); err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {}))
	mux.Handle("/1mb", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write(data1mb)
	}))
	mux.Handle("/10mb", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write(data10mb)
	}))
	mux.Handle("/100mb", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write(data100mb)
	}))
	mux.Handle("/1000mb", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write(data1000mb)
	}))

	w := os.Stdout

	server := http.Server{
		Addr: ":18000",
		TLSConfig: &tls.Config{
			MinVersion:   tls.VersionTLS13,
			MaxVersion:   tls.VersionTLS13,
			KeyLogWriter: w,
		},
		Handler: mux,
	}

	err := server.ListenAndServeTLS("../cert/server.crt", "../cert/private.key")
	if err != nil {
		log.Fatal(err)
	}
}
