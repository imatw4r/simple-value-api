package api

import (
	"context"
	"value-app/config"
	"value-app/pkg/domain"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateApp(svc domain.IValueService) *gin.Engine {
	controller := NewValueController(svc)
	router := gin.Default()
	router.GET("/index/:value", controller.GetValueIndex)
	return router
}

func RunWebserver(ctx context.Context, cancel context.CancelFunc, config *config.GlobalConfig, svc domain.IValueService) {
	log.Infof("Initiating web application...")
	router := CreateApp(svc)
	log.Infof("Starting web server at %s", config.App.Port)
	go func() {
		err := router.Run(":" + config.App.Port)
		if err != nil {
			panic(err)
		}
		log.Println("Cancelling context as server exited")
		cancel()
	}()
}
