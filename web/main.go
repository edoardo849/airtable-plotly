package main

import (
	"flag"
	"net/http"
)

var (
	addr string
)

func init() {
	// Tie the command-line flag to the intervalFlag variable and
	flag.StringVar(&addr, "addr", ":4040", "address port")
	flag.Parse()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(addr, nil)
}
