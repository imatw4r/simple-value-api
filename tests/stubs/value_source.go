package test_stubs

type InMemorySource struct{}

func (ms *InMemorySource) Load() ([]int, error) {
	return []int{700, 750, 1000, 1050, 1100, 1200, 1900}, nil
}

func NewValueSource() *InMemorySource {
	return &InMemorySource{}
}
