package test_stubs

type MemorySource struct{}

func (ms *MemorySource) Load() ([]int, error) {
	return []int{10, 100, 120, 140, 190, 300}, nil
}

func NewMemoryValueSource() *MemorySource {
	return &MemorySource{}
}
