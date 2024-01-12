package test_api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	api "value-app/pkg/api"
	"value-app/pkg/domain"

	test "value-app/tests"
	stub "value-app/tests/stubs"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ValueAPISuite struct {
	suite.Suite
	Router   *gin.Engine
	Recorder *httptest.ResponseRecorder
	Service  domain.IValueService
	Source   domain.IValueSource
}

func (s *ValueAPISuite) SetupTest() {
	values := []int{700, 750, 1000, 1050, 1100, 1200, 1900}
	s.Source = stub.NewValueSource(values)
	s.Recorder = httptest.NewRecorder()
	s.Service = test.GetValueService(s.Source)
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
				Found:        true,
			},
		},
		{
			"1150", api.GetValueIndexResponse{
				ErrorMessage: "",
				Value:        1100,
				Index:        4,
				Found:        true,
			},
		},
		{
			"700", api.GetValueIndexResponse{
				ErrorMessage: "",
				Value:        700,
				Index:        0,
				Found:        true,
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
				Found:        false,
			},
		},
		{
			"11.41223.123", api.GetValueIndexResponse{
				ErrorMessage: "Value 11.41223.123 is not a valid integer",
				Value:        -1,
				Index:        -1,
				Found:        false,
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
