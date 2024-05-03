package heapsort

type MaxHeap struct {
	slice []int
}

func NewMaxHeap() MaxHeap {
	return MaxHeap{
		slice: make([]int, 0),
	}
}

func NewMaxHeapFrom(input []int) MaxHeap {
	heap := NewMaxHeap()
	for _, i := range input {
		heap.Insert(i)
	}

	return heap
}

func (h MaxHeap) Size() int {
	return len(h.slice)
}

func (h *MaxHeap) Insert(key int) *MaxHeap {
	newSlice := append(h.slice, key)
	h.slice = heapifyUp(newSlice, len(newSlice)-1)

	return h
}

func (h *MaxHeap) Pull() int {
	if len(h.slice) == 0 {
		return -1
	}

	if len(h.slice) == 1 {
		max := h.slice[0]
		h.slice = nil
		return max
	}

	max := h.slice[0]
	lastIndex := len(h.slice) - 1
	h.slice[0], h.slice[lastIndex] = h.slice[lastIndex], h.slice[0]
	h.slice = NewMaxHeapFrom(h.slice[:lastIndex]).ToSlice()

	return max
}

func (h MaxHeap) ToSlice() []int {
	return h.slice
}

func (h MaxHeap) Max() int {
	if len(h.slice) == 0 {
		return -1
	}

	return h.slice[0]
}

func (h MaxHeap) Sort() []int {
	sorted := make([]int, len(h.slice))
	for i := len(h.slice) - 1; i >= 0; i-- {
		max := h.Pull()
		sorted[i] = max
	}

	return sorted
}

func heapifyUp(slice []int, from int) []int {
	p := parent(from)
	_ = p
	for slice[parent(from)] < slice[from] {
		slice = swap(slice, parent(from), from)
		from = parent(from)
	}

	return slice
}

func heapifyDown(slice []int) []int {
	tmp := make([]int, len(slice))

	return tmp
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1

}

func right(i int) int {
	return left(i) + 1
}

func swap(src []int, i, j int) []int {
	tmp := make([]int, len(src))
	copy(tmp, src)

	tmp[i], tmp[j] = tmp[j], tmp[i]

	return tmp
}
