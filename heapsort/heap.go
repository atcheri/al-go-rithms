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
	h.slice = heapify(newSlice, len(newSlice)-1)

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
	h.slice = heapify(h.slice[1:], len(h.slice[1:])-1)

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

func heapify(slice []int, from int) []int {
	for slice[parent(from)] < slice[from] {
		slice = swap(slice, parent(from), from)
		from = parent(from)
	}

	return slice
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
