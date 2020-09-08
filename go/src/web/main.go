package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func greeting(t string) (s string) {
	s = "<b>" + t + "</b>"
	return
}

// hello responds to the request with a plain-text "Code.education Rocks!" message.
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	msg := greeting("Code.education Rocks!")
	fmt.Fprintf(w, msg)
}

func main() {
	// register hello function to handle all requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	// use PORT environment variable, or default to 8000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
