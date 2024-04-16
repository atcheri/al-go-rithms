package depthfirstsearch

type Stack[T any] struct {
	items []T
}

func (q *Stack[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Stack[T]) Push(item T) {
	q.items = append(q.items, item)
}

func (q *Stack[T]) Pop() *T {
	first := &q.items[0]
	q.items = q.items[1:]
	return first
}
