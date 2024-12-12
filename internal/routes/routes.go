package routes

import (
	"net/http"
	"railway-go-api-template/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	healthHandler := handlers.NewHealthHandler()
	mux.HandleFunc("GET /health", healthHandler.HealthCheck)
}
