package test_domain

import (
	"fmt"
	"testing"

	"value-app/pkg/domain"
	test "value-app/tests"
	stub "value-app/tests/stubs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ValueServiceSuite struct {
	suite.Suite
	Service domain.IValueService
	Source  domain.IValueSource
}

func (s *ValueServiceSuite) SetupTest() {
	values := []int{700, 750, 1000, 1050, 1100, 1200, 1900}
	s.Source = stub.NewValueSource(values)
	s.Service = test.GetValueService(s.Source)
}

func TestRunValueServiceSuite(t *testing.T) {
	suite.Run(t, new(ValueServiceSuite))
}

func (suite *ValueServiceSuite) TestValueServiceDirectValueMatch() {
	svc := suite.Service
	var tests = []struct {
		value         int
		expectedValue int
		expectedIndex int
	}{
		{700, 700, 0},
		{750, 750, 1},
		{1000, 1000, 2},
		{1050, 1050, 3},
		{1900, 1900, 6},
	}
	for _, tt := range tests {
		suite.Suite.T().Run(fmt.Sprintf("Value %d Index %d", tt.value, tt.expectedIndex), func(t *testing.T) {
			result, _ := svc.IndexOf(tt.value)
			assert.Equal(t, result.Index, tt.expectedIndex, fmt.Sprintf("Got Index %d, Expected %d", result.Index, tt.expectedIndex))
			assert.Equal(t, result.Value, tt.expectedValue, fmt.Sprintf("Got Value %d, Expected %d", result.Value, tt.expectedValue))
		})
	}
}
func (suite *ValueServiceSuite) TestValueServiceAdjacentValueMatch() {
	svc := suite.Service
	var tests = []struct {
		name          string
		value         int
		expectedValue int
		expectedIndex int
	}{
		{"Match with left value while in between", 780, 750, 1},
		{"Match with right value while in between", 1150, 1100, 4},
		{"Match with edge left value", 730, 700, 0},
		{"Match with edge right value", 1800, 1900, 6},
	}
	for _, tt := range tests {
		suite.Suite.T().Run(tt.name, func(t *testing.T) {
			result, _ := svc.IndexOf(tt.value)
			assert.Equal(t, result.Index, tt.expectedIndex, fmt.Sprintf("Receive Index: %d, Expected: %d", result.Index, tt.expectedIndex))
			assert.Equal(t, result.Value, tt.expectedValue, fmt.Sprintf("Receive Value: %d, Expected: %d", result.Value, tt.expectedValue))
		})
	}
}

func (suite *ValueServiceSuite) TestValueServiceOnValueNotFound() {
	svc := suite.Service
	var tests = []struct {
		name          string
		value         int
		expectedValue int
		expectedIndex int
	}{
		{"Negative value", -1324, -1, -1},
		{"Value much lower than any existing", 1, -1, -1},
		{"Value much bigger than any existing", 1000000, -1, -1},
	}
	for _, tt := range tests {
		suite.Suite.T().Run(fmt.Sprintf(tt.name, tt.value, tt.expectedIndex), func(t *testing.T) {
			result, _ := svc.IndexOf(tt.value)
			assert.Equal(t, result.Index, tt.expectedIndex, fmt.Sprintf("Receive Index %d, Expected %d", result.Index, tt.expectedIndex))
			assert.Equal(t, result.Value, tt.expectedValue, fmt.Sprintf("Receive Value %d, Expected %d", result.Value, tt.expectedValue))
		})
	}
}
