package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/NarHakobyan/go-httpbin"
)

var (
	host = flag.String("host", ":8080", "<host:port>")
	printLogs = flag.Bool("log", false, "<boolean>")
)

func main() {
	flag.Parse()

	log.Printf("httpbin listening on %s", *host)
	log.Fatal(http.ListenAndServe(*host, httpbin.GetMux(printLogs)))
}
