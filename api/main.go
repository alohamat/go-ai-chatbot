package main

import (
	"log"
	"net/http"

	"github.com/alohamat/go-ai-chatbot/routes"
)

func main() {
	r := routes.InitRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("server running at port", port)
	log.Fatalln(http.ListenAndServe(":"+port, r))
}