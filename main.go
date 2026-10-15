// maintenance-in-progress answers every request with a static page
// saying the site is down for maintenance, as a 503 so clients and
// crawlers know it's temporary. See handler.
package main

import (
	_ "embed"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed index.html
var page []byte

// retryAfter is how long clients are told to wait, in seconds; the
// page reloads itself on the same beat.
const retryAfter = "60"

// handler serves the page on any path and method, never cached, so
// the site's own pages come back as soon as the maintenance ends.
func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Retry-After", retryAfter)
		w.WriteHeader(http.StatusServiceUnavailable)

		if r.Method != http.MethodHead {
			w.Write(page)
		}
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           handler(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("serving the maintenance page on :%s", port)
	log.Fatal(srv.ListenAndServe())
}
