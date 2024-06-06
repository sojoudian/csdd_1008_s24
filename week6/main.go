package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}
	fmt.Fprintln(w, "Welcome to the Go API Server!")
}

// go mod init github.com/sojoudian/SimpleAPIServer
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	log.Println("Starting server on :3001")
	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
