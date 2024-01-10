package test_domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueServiceOnValueFound(t *testing.T) {
	svc := getValueService()
	var tests = []struct {
		value    int
		expected int
	}{
		{10, 0},
		{100, 1},
		{120, 2},
		{140, 3},
		{190, 4},
		{300, 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Value %d Index %d", tt.value, tt.expected), func(t *testing.T) {
			result, _ := svc.IndexOf(tt.value)
			assert.Equal(t, result.Index, tt.expected, fmt.Sprintf("Got Index %d, Expected %d", result.Index, tt.expected))
		})
	}
}
