package test_stubs

type InMemorySource struct {
	Values []int
}

func (ms *InMemorySource) Load() ([]int, error) {
	return ms.Values, nil
}

func NewValueSource(values []int) *InMemorySource {
	return &InMemorySource{Values: values}
}
