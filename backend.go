// backend.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from the local backend at %s!", r.URL.Path)
	})

	log.Println("✅ Backend running on http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}
