package test_api

import (
	"value-app/api"
	"value-app/domain"

	"github.com/gin-gonic/gin"
)

func getTestClient(svc domain.IValueService) *gin.Engine {
	return api.CreateApp(svc)
}
