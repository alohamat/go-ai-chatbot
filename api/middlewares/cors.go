package middlewares

import (
	"net/http"
	"github.com/rs/cors"
)

func CorsMiddleware(handler http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
		AllowCredentials: false,
	}).Handler(handler)
}