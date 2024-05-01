package heapsort

type MaxHeap struct {
	slice []int
}

func NewMaxHeap() MaxHeap {
	return MaxHeap{
		slice: make([]int, 0),
	}
}

func (h MaxHeap) Size() int {
	return len(h.slice)
}

func (h *MaxHeap) Insert(key int) *MaxHeap {
	h.slice = append(h.slice, key)
	return h
}

func (h MaxHeap) ToSlice() []int {
	return h.slice
}
