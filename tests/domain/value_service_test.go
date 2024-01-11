package test_domain

import (
	"fmt"
	"testing"

	test "value-app/tests"

	"github.com/stretchr/testify/assert"
)

func TestValueServiceDirectValueMatch(t *testing.T) {
	svc := test.GetValueService()
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
		t.Run(fmt.Sprintf("Value %d Index %d", tt.value, tt.expectedIndex), func(t *testing.T) {
			result, _ := svc.IndexOf(tt.value)
			assert.Equal(t, result.Index, tt.expectedIndex, fmt.Sprintf("Got Index %d, Expected %d", result.Index, tt.expectedIndex))
			assert.Equal(t, result.Value, tt.expectedValue, fmt.Sprintf("Got Value %d, Expected %d", result.Value, tt.expectedValue))
		})
	}
}
func TestValueServiceAdjacentValueMatch(t *testing.T) {
	svc := test.GetValueService()
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
		t.Run(tt.name, func(t *testing.T) {
			result, _ := svc.IndexOf(tt.value)
			assert.Equal(t, result.Index, tt.expectedIndex, fmt.Sprintf("Receive Index: %d, Expected: %d", result.Index, tt.expectedIndex))
			assert.Equal(t, result.Value, tt.expectedValue, fmt.Sprintf("Receive Value: %d, Expected: %d", result.Value, tt.expectedValue))
		})
	}
}

func TestValueServiceOnValueNotFound(t *testing.T) {
	svc := test.GetValueService()
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
		t.Run(fmt.Sprintf(tt.name, tt.value, tt.expectedIndex), func(t *testing.T) {
			result, _ := svc.IndexOf(tt.value)
			assert.Equal(t, result.Index, tt.expectedIndex, fmt.Sprintf("Receive Index %d, Expected %d", result.Index, tt.expectedIndex))
			assert.Equal(t, result.Value, tt.expectedValue, fmt.Sprintf("Receive Value %d, Expected %d", result.Value, tt.expectedValue))
		})
	}
}
