package domain

type Result struct {
	Value int
	Index int
}

// Interface for a domain-service
type IValueService interface {
	// Add Values
	AddValues([]int) error
	// Return value index
	IndexOf(int) (Result, error)
}
