package generics_test

import (
	"testing"
)

// ListNode 节点
type ListNode[T any] struct {
	data T
	next *ListNode[T]
}

type List[T any] struct {
	head *ListNode[T]
	size int
}

// ListIterator 迭代器
// type ListIterator[T any] struct {
// 	current *ListNode[T]
// }

// NewList 创建一个链表
func NewList[T any]() *List[T] {
	return &List[T]{}
}

// Add 添加元素
func (l *List[T]) Add(data T) {
	node := &ListNode[T]{data: data}
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
func (l *List[T]) Remove(index int) bool {
	if index < 0 || index >= l.size {
		return false
	}
	if index == 0 {
		l.head = l.head.next
	} else {
		prevNode := l.getNode(index - 1)
		prevNode.next = prevNode.next.next
	}
	l.size--
	return true
}

// getNode 获取指定索引的节点 辅助函数
func (l *List[T]) getNode(index int) *ListNode[T] {
	cNode := l.head
	for i := 0; i < index; i++ {
		cNode = cNode.next
	}
	return cNode
}

func TestGenericsList(t *testing.T) {
	// 创建一个泛型链表实例
	list := NewList[int]()
	if list == nil {
		t.Fatalf("Failed to create a new list")
	}

	// 测试添加元素
	list.Add(10)
	list.Add(20)
	list.Add(30)

	// 验证链表大小
	if list.size != 3 {
		t.Errorf("Expected list size to be 3, got %d", list.size)
	}

	// 测试获取指定索引节点
	node := list.getNode(1)
	if node.data != 20 {
		t.Errorf("Expected node data at index 1 to be 20, got %v", node.data)
	}

	// 测试删除元素
	success := list.Remove(1)
	if !success {
		t.Errorf("Failed to remove element at index 1")
	}

	// 验证删除后的链表大小
	if list.size != 2 {
		t.Errorf("Expected list size to be 2 after removal, got %d", list.size)
	}

	// 验证删除后的节点数据
	node = list.getNode(1)
	if node.data != 30 {
		t.Errorf("Expected node data at index 1 to be 30 after removal, got %v", node.data)
	}
}
