package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Msg структура для сообщений
type Msg struct {
	Message string `json:"message"`
}

var messages []Msg

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

	messages = append(messages, m) // сохраняем в памяти

	json.NewEncoder(w).Encode(map[string]interface{}{
		"reply":    "You sent: " + m.Message,
		"messages": messages,
	})
}

func main() {
	// Статика (HTML, CSS, JS)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/echo", echoHandler)

	// Render даёт порт через переменную PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback для локального запуска
	}

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
