package breadthfirstsearch

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() *T {
	last := &q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return last
}
