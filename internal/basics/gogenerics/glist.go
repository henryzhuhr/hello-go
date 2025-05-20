package gogenerics

// GListNode 节点
type GListNode[T any] struct {
	data T
	next *GListNode[T]
}

// Data 返回节点的 data 值
func (n *GListNode[T]) Data() T {
	return n.data
}

// GListIterator 迭代器
// 这个迭代器的作用是方便地遍历链表中的元素。
type GListIterator[T any] struct {
	current *GListNode[T]
}

// NewGListIterator 创建一个新的迭代器
func (l *GList[T]) NewGListIterator() *GListIterator[T] {
	return &GListIterator[T]{current: l.head}
}

// HasNext 检查是否还有下一个元素
func (it *GListIterator[T]) HasNext() bool {
	return it.current != nil
}

// Next 返回当前元素并移动到下一个元素
func (it *GListIterator[T]) Next() T {
	if it.current == nil {
		panic("no more elements")
	}
	data := it.current.data
	it.current = it.current.next
	return data
}

type GList[T any] struct {
	head *GListNode[T]
	size int
}

// NewGList 创建一个链表
func NewGList[T any]() *GList[T] {
	return &GList[T]{}
}

// Add 添加元素
func (l *GList[T]) Add(data T) {
	node := &GListNode[T]{data: data}
	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	l.size++
}

// Remove 删除元素
func (l *GList[T]) Remove(index int) bool {
	if index < 0 || index >= l.size {
		return false
	}
	if index == 0 {
		l.head = l.head.next
	} else {
		prevNode := l.GetNode(index - 1)
		prevNode.next = prevNode.next.next
	}
	l.size--
	return true
}

// getNode 获取指定索引的节点 辅助函数
func (l *GList[T]) GetNode(index int) *GListNode[T] {
	cNode := l.head
	for i := 0; i < index; i++ {
		cNode = cNode.next
	}
	return cNode
}

// Size 获取链表大小
func (l *GList[T]) Size() int {
	return l.size
}
