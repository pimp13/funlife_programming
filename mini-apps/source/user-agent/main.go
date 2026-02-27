package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, v any, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func main() {
	mx := http.NewServeMux()

	mx.HandleFunc("GET /user-agent/{id}", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		ip := r.RemoteAddr
		queryParam := r.URL.Query()
		pathValueParam := r.PathValue("id")
		WriteJSON(w, map[string]string{
			"userAgent":  userAgent,
			"ip":         ip,
			"pathValue":  pathValueParam,
			"queryValue": queryParam.Get("name"),
		}, http.StatusOK)
	})

	const port = ":9001"
	log.Printf("Server is running on %s\n", port)
	log.Fatalln(http.ListenAndServe(port, mx))
}
