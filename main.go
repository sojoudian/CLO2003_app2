package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Welcome to the Go API\n")
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var msg Message
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid read request body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &msg)
	if err != nil {
		http.Error(w, "Invalid read request body", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content/Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func main() {
	http.HandleFunc("/", getHandler)
	http.HandleFunc("/post", postHandler)
	pNumber := ":8091"
	fmt.Printf("Server is running on the port: %s\n", pNumber)
	http.ListenAndServe(pNumber, nil)
}
