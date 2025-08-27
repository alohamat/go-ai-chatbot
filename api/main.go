package main

import (
	"log"
	"net/http"

	"github.com/alohamat/go-ai-chatbot/routes"
)

func main() {
	r := routes.InitRouter()
	log.Println("server running at 8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}