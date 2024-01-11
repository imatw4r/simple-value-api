package domain

import (
	"math"
	"sort"

	log "github.com/sirupsen/logrus"
)

type ValueService struct {
	values                   []int
	acceptableDiffPercentage float64
}

func isAcceptableValue(initialValue int, givenValue int, acceptableDiffPercentage float64) bool {
	valueDifference := math.Abs(float64(initialValue - givenValue))
	acceptableDifference := math.Abs(float64(initialValue) * (acceptableDiffPercentage / 100))
	log.Debugf("Value difference: %f", valueDifference)
	log.Debugf("Acceptable difference: %f", acceptableDifference)
	return valueDifference <= acceptableDifference
}

// IndexOf returns an Index of given Value.
// If the Value is not found, two adjacent Values are checked.
// If any of two Values is within acceptableDiffPercentage`
// from the given Value, its Index is returned instead,
// otherwise Index -1 is returned.
// Left Value Index has higher priority over right Value.
func (s *ValueService) IndexOf(value int) (Result, error) {
	index := sort.Search(len(s.values), func(i int) bool { return s.values[i] >= value })
	log.Debugf("Processing Value: %d", value)
	log.Debugf("Initial Index: %d", index)

	valueFound := 0 <= index && index < len(s.values) && s.values[index] == value
	if valueFound {
		return Result{Value: value, Index: index}, nil
	}

	// Consider not found idx (-1) to be edge idx also
	isLeftEdgeIdx := index <= 0
	isRightEdgeIdx := index == len(s.values)

	resultNotFound := Result{Value: -1, Index: -1}

	if isLeftEdgeIdx {
		// Check only right side
		log.Debugf("Value on the left edge")
		right_value := s.values[index+1]
		if isAcceptableValue(value, right_value, s.acceptableDiffPercentage) {
			log.Debugf("Final Index: %d", index+1)
			return Result{Index: index + 1, Value: right_value}, nil
		}
		log.Debugf("Final Index: %d", resultNotFound.Index)
		return resultNotFound, nil
	} else if isRightEdgeIdx {
		// Check only left side
		log.Debugf("Value on the right edge.")
		left_value := s.values[index-1]
		if isAcceptableValue(value, left_value, s.acceptableDiffPercentage) {
			log.Debugf("Final Index: %d", index-1)
			return Result{Index: index - 1, Value: left_value}, nil
		}
		log.Debugf("Final Index: %d", resultNotFound.Index)
		return resultNotFound, nil
	}

	// Index in between other values, so check both
	log.Debugf("Value in between other values.")
	left_value := s.values[index-1]
	right_value := s.values[index]

	if isAcceptableValue(value, left_value, s.acceptableDiffPercentage) {
		log.Debugf("Final Index: %d", index-1)
		return Result{Index: index - 1, Value: left_value}, nil
	} else if isAcceptableValue(value, right_value, s.acceptableDiffPercentage) {
		log.Debugf("Final Index: %d", index)
		return Result{Index: index, Value: right_value}, nil
	}

	log.Debugf("Final Index: %d", resultNotFound.Index)
	return resultNotFound, nil
}

func (vs *ValueService) AddValues(values []int) error {
	vs.values = append(vs.values, values...)
	return nil
}

func NewValueService(valueDiffPercentage float64) *ValueService {
	return &ValueService{values: []int{}, acceptableDiffPercentage: valueDiffPercentage}
}
