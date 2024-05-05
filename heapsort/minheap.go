package heapsort

type MinHeap struct {
	slice []int
}

func NewMinHeap() MinHeap {
	return MinHeap{
		slice: []int{},
	}
}

func NewMinHeapFrom(input []int) MinHeap {
	heap := NewMinHeap()
	for _, i := range input {
		heap.Insert(i)
	}

	return heap
}

func (h *MinHeap) Insert(key int) *MinHeap {
	newSlice := append(h.slice, key)
	h.slice = minheapifyUp(newSlice, len(newSlice)-1)

	return h

}

func minheapifyUp(slice []int, from int) []int {
	p := parent(from)
	_ = p
	// the only different with the heapifuUp for the MaxHeap is the comparaison sign
	for slice[parent(from)] > slice[from] {
		slice = swap(slice, parent(from), from)
		from = parent(from)
	}

	return slice
}

func (h MinHeap) ToSlice() []int {
	return h.slice
}
