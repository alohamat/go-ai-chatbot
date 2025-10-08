package middlewares

import (
	"net/http"
	"github.com/rs/cors"
)

func CorsMiddleware(handler http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"https://react-ai-frontend-gamma.vercel.app", "http://localhost:8080"},
		AllowedMethods: []string{"POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(handler)
}