package api

import (
	"context"
	"value-app/config"
	"value-app/domain"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RunWebserver(ctx context.Context, cancel context.CancelFunc, config config.AppConfig, svc domain.IValueService) {
	controller := NewValueController(svc)
	router := gin.Default()
	router.GET("/index/:value", controller.GetValueIndex)

	log.Printf("Starting web server at %s", config.Port)
	go func() {
		err := router.Run(":" + config.Port)
		if err != nil {
			panic(err)
		}
		log.Println("Cancelling context as server exited")
		cancel()
	}()
}
