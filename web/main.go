package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":8082", nil)
}
