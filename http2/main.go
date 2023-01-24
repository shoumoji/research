package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {}))

	w := os.Stdout

	server := http.Server{
		Addr: "127.0.0.1:18000",
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
