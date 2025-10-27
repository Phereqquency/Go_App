package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
)

type Msg struct {
	Message string `json:"message"`
}

var (
	messages []Msg
	mu       sync.Mutex
)

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

	mu.Lock()
	messages = append(messages, m)
	resp := map[string]interface{}{
		"messages": messages,
	}
	mu.Unlock()

	json.NewEncoder(w).Encode(resp)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/echo", echoHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
