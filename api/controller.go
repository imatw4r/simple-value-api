package api

import (
	"net/http"
	"strconv"
	"value-app/domain"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type IValueController interface {
	GetValueIndex(c *gin.Context)
}

type ValueController struct {
	svc domain.IValueService
}

func validateValue(valueStr string) (int, error) {
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return -1, err
	}
	return value, nil
}

func (vc *ValueController) GetValueIndex(c *gin.Context) {
	valueStr := c.Param("value")
	valueInt, err := validateValue(valueStr)
	log.Debugf("Looking for value of %d", valueInt)

	if err != nil {
		log.Warn("Value is not a valid integer")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value provided"})
		return
	}

	result, err := vc.svc.IndexOf(valueInt)

	if err != nil {
		log.Warn("Failed to retrieve an index")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve the index"})
		return
	}

	if result.Index == -1 {
		log.Info("Value index not found")
		response := gin.H{
			"index":               result.Index,
			"corresponding_value": result.Value,
			"message":             "NotFound",
		}
		c.JSON(http.StatusOK, response)
		return
	}

	log.Info("Value index found")
	log.Debugf("Value: %d Index: %d", result.Value, result.Index)
	response := gin.H{
		"index":               result.Index,
		"corresponding_value": result.Value,
		"message":             "Success",
	}

	c.JSON(http.StatusOK, response)
}

func NewValueController(svc domain.IValueService) *ValueController {
	return &ValueController{
		svc: svc,
	}
}
