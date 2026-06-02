package main

import (
	"github.com/Munir/airtime-api/internal/config"
	"github.com/Munir/airtime-api/internal/handlers"
	"github.com/Munir/airtime-api/internal/service"
	"github.com/Munir/airtime-api/internal/trueafrican"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	client := &trueafrican.Client{
		URL:      cfg.APIURL,
		Username: cfg.User,
		Password: cfg.Pass,
	}

	svc := &service.AirtimeService{
		Client: client,
	}

	handler := &handlers.Handler{
		Service: svc,
	}

	router := gin.Default()

	router.POST("/api/v1/airtime/purchase", handler.Purchase)

	router.Run(":" + cfg.AppPort)
}
