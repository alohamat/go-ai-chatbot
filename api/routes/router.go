package routes

import (
	"github.com/alohamat/go-ai-chatbot/handlers"
	"github.com/alohamat/go-ai-chatbot/middlewares"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(middlewares.CorsMiddleware)
	
	router.HandleFunc("/api/send", handlers.SendHandler).Methods("POST", "OPTIONS")
	
	return router
}