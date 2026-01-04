package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Static frontend
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fileServer)

	// API routes
	mux.HandleFunc("/time", logRequest(timeHandler))
	mux.HandleFunc("/api/info", logRequest(infoHandler))

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
