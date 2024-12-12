package internal

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"railway-go-api-template/internal/config"
	"railway-go-api-template/internal/routes"
	"syscall"
	"time"
)

func Run() {
	// Load the env file
	config.LoadEnv()

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	cors := config.Cors()

	port := ":" + config.GetEnvWithKey("PORT", "8000")

	server := &http.Server{
		Addr:    port,
		Handler: cors.Handler(mux),
	}

	go func() {
		log.Printf("Server started: http://localhost%s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
