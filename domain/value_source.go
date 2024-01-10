package domain

// Interface for loading number values
type IValueSource interface {
	Load() ([]int, error)
}
