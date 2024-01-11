package test_api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	api "value-app/api"
	"value-app/domain"

	test "value-app/tests"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/assert"
)

type ValueAPISuite struct {
	suite.Suite
	Router   *gin.Engine
	Recorder *httptest.ResponseRecorder
	Service  domain.IValueService
}

func (s *ValueAPISuite) SetupTest() {
	s.Recorder = httptest.NewRecorder()
	s.Service = test.GetValueService()
	s.Router = getTestClient(s.Service)
}

func TestRunValueAPISuite(t *testing.T) {
	suite.Run(t, new(ValueAPISuite))
}

func (suite *ValueAPISuite) TestGetValueIndexReturn200() {
	var testCases = []struct {
		Value            string
		ExpectedResponse api.GetValueIndexResponse
	}{
		{
			"1000", api.GetValueIndexResponse{
				ErrorMessage: "",
				Value:        1000,
				Index:        2,
			},
		},
		{
			"1150", api.GetValueIndexResponse{
				ErrorMessage: "",
				Value:        1100,
				Index:        4,
			},
		},
		{
			"700", api.GetValueIndexResponse{
				ErrorMessage: "",
				Value:        700,
				Index:        0,
			},
		},
	}

	for _, tc := range testCases {
		suite.T().Run(fmt.Sprintf("Testing value %s", tc.Value), func(t *testing.T) {
			endpoint := fmt.Sprintf("/index/%s", tc.Value)
			request, err := http.NewRequest(http.MethodGet, endpoint, nil)

			assert.NoError(suite.T(), err)

			suite.Router.ServeHTTP(suite.Recorder, request)

			assert.Equal(suite.T(), 200, suite.Recorder.Result().StatusCode)

			var response api.GetValueIndexResponse
			decoder := json.NewDecoder(suite.Recorder.Body)
			err = decoder.Decode(&response)
			assert.NoError(suite.T(), err)

			assert.Equal(suite.T(), tc.ExpectedResponse, response)
		})
	}
}

func (suite *ValueAPISuite) TestGetValueIndexReturn400() {
	var testCases = []struct {
		Value            string
		ExpectedResponse api.GetValueIndexResponse
	}{
		{
			"fasf", api.GetValueIndexResponse{
				ErrorMessage: "Value fasf is not a valid integer",
				Value:        -1,
				Index:        -1,
			},
		},
		{
			"11.41223.123", api.GetValueIndexResponse{
				ErrorMessage: "Value 11.41223.123 is not a valid integer",
				Value:        -1,
				Index:        -1,
			},
		},
	}

	for _, tc := range testCases {
		suite.T().Run(fmt.Sprintf("Testing value %s", tc.Value), func(t *testing.T) {
			endpoint := fmt.Sprintf("/index/%s", tc.Value)
			request, err := http.NewRequest(http.MethodGet, endpoint, nil)

			assert.NoError(suite.T(), err)

			suite.Router.ServeHTTP(suite.Recorder, request)

			assert.Equal(suite.T(), 400, suite.Recorder.Result().StatusCode)

			var response api.GetValueIndexResponse
			decoder := json.NewDecoder(suite.Recorder.Body)
			err = decoder.Decode(&response)
			assert.NoError(suite.T(), err)

			assert.Equal(suite.T(), tc.ExpectedResponse, response)
		})
	}
}
