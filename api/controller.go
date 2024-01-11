package api

import (
	"fmt"
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

func (vc *ValueController) GetValueIndex(c *gin.Context) {
	var request GetValueIndexRequest
	var response GetValueIndexResponse

	if err := c.ShouldBindUri(&request); err != nil {
		errorMsg := err.Error()
		log.Warn(errorMsg)
		response = GetValueIndexResponse{
			Index:        -1,
			Value:        -1,
			ErrorMessage: errorMsg,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	log.Infof("Received request with Value %s", request.Value)

	valueInt, err := convertToInt(request.Value)

	if err != nil {
		errorMsg := fmt.Sprintf("Value %s is not a valid integer", request.Value)
		log.Warn(errorMsg)
		response = GetValueIndexResponse{
			Index:        -1,
			Value:        -1,
			ErrorMessage: errorMsg,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := vc.svc.IndexOf(valueInt)

	if err != nil {
		errorMsg := "Failed to retrieve an index"
		log.Warn(errorMsg)
		response = GetValueIndexResponse{
			Index:        -1,
			Value:        -1,
			ErrorMessage: errorMsg,
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if result.Index == -1 {
		errorMsg := "Index not found"
		log.Info(errorMsg)
		response = GetValueIndexResponse{
			Index:        -1,
			Value:        valueInt,
			ErrorMessage: errorMsg,
		}
		c.JSON(http.StatusOK, response)
		return
	}

	log.Info("Index found")
	response = GetValueIndexResponse{
		Value:        result.Value,
		Index:        result.Index,
		ErrorMessage: "",
	}

	c.JSON(http.StatusOK, response)
}

func NewValueController(svc domain.IValueService) *ValueController {
	return &ValueController{
		svc: svc,
	}
}

func convertToInt(valueStr string) (int, error) {
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return -1, err
	}
	return value, nil
}
