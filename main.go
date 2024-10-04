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
	fmt.Fprintf(w, "Welcome to the Go API\n")
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	var msg Message
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &msg)
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
