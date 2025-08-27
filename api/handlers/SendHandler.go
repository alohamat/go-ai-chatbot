package handlers

import (
	"net/http"
	"encoding/json"
	"log"

	"github.com/alohamat/go-ai-chatbot/services"
)

type prompt struct {
	Prompt string `json:"prompt"`
}

func SendHandler(w http.ResponseWriter, r *http.Request) {
	var req prompt
	err := json.NewDecoder(r.Body).Decode(&req)
	if (err != nil) {
		http.Error(w, "error decoding json", http.StatusInternalServerError)
		return
	}
	log.Printf("json received: %s", req.Prompt)
	result := services.AITextSend(req.Prompt)
	log.Printf("ai result %s", result)
	
	json.NewEncoder(w).Encode(result)
}