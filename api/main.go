package main

import (
	"context"
	"flag"
	"log"
	"net/http"
)

var (
	addr   string
	apiKey string
)

func init() {
	// Tie the command-line flag to the intervalFlag variable and

	flag.StringVar(&addr, "addr", ":8080", "address port")
	flag.StringVar(&apiKey, "apiKey", "", "Airtable's api key")
	flag.Parse()
}

func main() {
	log.Print(apiKey)
	s := &Server{
		APIKey: apiKey,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/ideas/", withCORS(withAPIKey(s.handlePolls)))
	log.Println("Starting web server on", addr)
	http.ListenAndServe(":8080", mux)
	log.Println("Stopping...")
}

// Server is the API server.
type Server struct {
	APIKey string
}

func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Location")
		fn(w, r)
	}
}

type contextKey struct {
	name string
}

var contextKeyAPIKey = &contextKey{"api-key"}

// APIKey initiates a new Api key
func APIKey(ctx context.Context) (string, bool) {
	key := ctx.Value(contextKeyAPIKey)
	if key == nil {
		return "", false
	}
	keystr, ok := key.(string)
	return keystr, ok
}

func withAPIKey(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if !isValidAPIKey(key) {
			respondErr(w, r, http.StatusUnauthorized, "invalid API key")
			return
		}
		ctx := context.WithValue(r.Context(), contextKeyAPIKey, key)
		fn(w, r.WithContext(ctx))
	}
}

func isValidAPIKey(key string) bool {
	return key == "abc123"
}
