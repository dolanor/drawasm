package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	staticFS := http.FileServer(http.Dir("./static"))

	host := "localhost:9999"
	fmt.Printf("Serving on http://%s\n", host)
	http.ListenAndServe(host, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers for wasm
		w.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(r.URL.Path, ".wasm") {
			w.Header().Set("Content-Type", "application/wasm")
		}
		staticFS.ServeHTTP(w, r)
	}))
}
