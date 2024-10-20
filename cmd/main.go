package main

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gtihub.com/nishanth-gowda/core-service/api"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "error", err)
		os.Exit(1)
	}

	// Get the HTTP port from the environment variable
	httpPortStr := os.Getenv("HTTP_PORT")
	if httpPortStr == "" {
		slog.Error("HTTP_PORT environment variable is not set")
		os.Exit(1)
	}

	// Convert the port string to an integer
	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		slog.Error("Invalid HTTP_PORT value", "error", err)
		os.Exit(1)
	}

	// Create a new Gin router
	app := gin.New()

	// Route
	api.Route(app.Group("/v1"))

	// Log the port the server is starting on
	slog.Info("Starting server on port", "port", httpPort)

	// Run on the specified port
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(httpPort),
		Handler: app,
	}

	// Start the server and log any errors
	if err := server.ListenAndServe(); err != nil {
		slog.Error("Error starting server", "error", err)
	}
}