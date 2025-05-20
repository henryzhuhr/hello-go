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

type GList[T any] struct {
	head *GListNode[T]
	size int
}

// ListIterator 迭代器
// type ListIterator[T any] struct {
// 	current *ListNode[T]
// }

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
