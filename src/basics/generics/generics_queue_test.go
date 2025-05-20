package generics_test

import "testing"

type Queue[T any] struct {
	items []T
}

// NewQueue 创建一个新的队列
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue 将元素添加到队列的末尾
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue 从队列的前面移除并返回一个元素
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Peek 返回队列中的第一个元素但不移除它
func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	return q.items[0], true
}

// IsEmpty 检查队列是否为空
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size 返回队列中的元素数量
func (q *Queue[T]) Size() int {
	return len(q.items)
}

func TestQueue(t *testing.T) {
	q := NewQueue[int]()

	if !q.IsEmpty() {
		t.Errorf("队列初始化后应为空")
	}

	if size := q.Size(); size != 0 {
		t.Errorf("队列初始化大小应为 0，实际为 %d", size)
	}

	// 测试入队和出队
	values := []int{1, 2, 3}
	for _, v := range values {
		q.Enqueue(v)
	}

	if q.IsEmpty() {
		t.Errorf("队列添加元素后不应为空")
	}

	if size := q.Size(); size != len(values) {
		t.Errorf("队列大小应为 %d，实际为 %d", len(values), size)
	}

	expected := []int{1, 2, 3}
	for i, exp := range expected {
		item, ok := q.Dequeue()
		if !ok {
			t.Errorf("队列非空时 Dequeue 应该成功，失败于索引 %d", i)
		}
		if item != exp {
			t.Errorf("Dequeue 返回的值错误，期望 %v，实际 %v", exp, item)
		}
	}

	if !q.IsEmpty() {
		t.Errorf("所有元素出队后队列应该为空")
	}

	// 测试空队列出队
	if _, ok := q.Dequeue(); ok {
		t.Errorf("空队列出队应该返回 false")
	}

	// 测试 Peek
	for _, v := range values {
		q.Enqueue(v)
	}
	if peek, ok := q.Peek(); !ok || peek != values[0] {
		t.Errorf("Peek 返回的值错误，期望 %v，实际 %v", values[0], peek)
	}

	// 确保 Peek 不移除元素
	if size := q.Size(); size != len(values) {
		t.Errorf("Peek 后队列大小应保持不变，期望 %d，实际 %d", len(values), size)
	}
}
