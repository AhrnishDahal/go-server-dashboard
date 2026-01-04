package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().Format(time.RFC1123)))
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"version": "1.0.0",
	})
}
