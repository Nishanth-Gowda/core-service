package main

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gtihub.com/nishanth-gowda/core-service/api"
	"gtihub.com/nishanth-gowda/core-service/config"
)

func main() {

	config.LoanConfig()


	app := gin.New()

	// Route
	api.Route(app.Group("/v1"))

	// Run on PORT 8090
	server := &http.Server{
		Addr:    ":" + viper.GetString("HTTP_PORT"),
		Handler: app,
	}

	if err := server.ListenAndServe(); err != nil {
		slog.Any("Error starting server", err)
	}


}