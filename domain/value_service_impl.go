package domain

import (
	"math"
	"sort"
)

type ValueService struct {
	values                   []int
	acceptableDiffPercentage float64
}

func isAcceptableValue(initialValue int, givenValue int, acceptableDiffPercentage float64) bool {
	percentageDifference := math.Abs(float64(initialValue-givenValue)) / float64(givenValue) * 100
	return percentageDifference <= acceptableDiffPercentage
}

func (s *ValueService) IndexOf(value int) (Result, error) {
	index := sort.Search(len(s.values), func(i int) bool { return s.values[i] >= value })

	valueFound := index < len(s.values) && s.values[index] == value
	if valueFound {
		return Result{Value: s.values[index], Index: index}, nil
	}

	// Check if other value can be returned
	left_value := s.values[index-1]
	right_value := s.values[index+1]

	if isAcceptableValue(value, left_value, s.acceptableDiffPercentage) {
		return Result{Value: left_value, Index: index - 1}, nil
	} else if isAcceptableValue(value, right_value, s.acceptableDiffPercentage) {
		return Result{Value: right_value, Index: index + 1}, nil
	} else {
		return Result{Value: -1, Index: -1}, nil
	}

}

func (vs *ValueService) AddValues(values []int) error {
	vs.values = append(vs.values, values...)
	return nil
}

func NewValueService(valueDiffPercentage float64) *ValueService {
	return &ValueService{values: []int{}, acceptableDiffPercentage: valueDiffPercentage}
}
