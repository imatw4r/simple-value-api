package api

import (
	"context"
	"value-app/common"
	"value-app/domain"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RunWebserver(ctx context.Context, cancel context.CancelFunc, config *common.GlobalConfig, svc domain.IValueService) {
	log.Infof("Initiating web application...")
	controller := NewValueController(svc)
	router := gin.Default()
	router.GET("/index/:value", controller.GetValueIndex)

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
