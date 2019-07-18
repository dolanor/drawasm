package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	dir := "./static"
	staticFS := http.FileServer(http.Dir(dir))

	host := ":9999"
	fmt.Printf("Serving on %s\n", host)
	http.ListenAndServe(":9999", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers for wasm
		w.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(r.URL.Path, ".wasm") {
			w.Header().Set("Content-Type", "application/wasm")
		}
		staticFS.ServeHTTP(w, r)
	}))
}
