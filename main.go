package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Msg struct {
	Message string `json:"message"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "use POST"})
		return
	}

	var m Msg
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
		return
	}

	resp := map[string]string{"reply": "You sent: " + m.Message}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/api/echo", echoHandler)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
