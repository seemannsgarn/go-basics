// server1 this is minimum "echo"-server

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // every request call handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler is returning a peace of path from URL request
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
