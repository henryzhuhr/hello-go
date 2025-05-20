package gogenerics

type GQueue[T any] struct {
	items []T
}

// NewGQueue 创建一个新的队列
func NewGQueue[T any]() *GQueue[T] {
	return &GQueue[T]{}
}

// Enqueue 将元素添加到队列的末尾
func (q *GQueue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue 从队列的前面移除并返回一个元素
func (q *GQueue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Peek 返回队列中的第一个元素但不移除它
func (q *GQueue[T]) Peek() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	return q.items[0], true
}

// IsEmpty 检查队列是否为空
func (q *GQueue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size 返回队列中的元素数量
func (q *GQueue[T]) Size() int {
	return len(q.items)
}
