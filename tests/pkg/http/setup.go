package test_api

import (
	api "value-app/pkg/api"
	"value-app/pkg/domain"

	"github.com/gin-gonic/gin"
)

func getTestClient(svc domain.IValueService) *gin.Engine {
	return api.CreateApp(svc)
}
