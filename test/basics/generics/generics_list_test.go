package generics_test

import (
	"testing"

	"github.com/henryzhuhr/hello-go/internal/basics/gogenerics"
)

func TestGenericsList(t *testing.T) {

	// 创建一个泛型链表实例
	list := gogenerics.NewGList[int]()
	if list == nil {
		t.Fatalf("Failed to create a new list")
	}

	// 测试添加元素
	list.Add(10)
	list.Add(20)
	list.Add(30)

	// 验证链表大小
	if list.Size() != 3 {
		t.Errorf("Expected list size to be 3, got %d", list.Size())
	}

	// 测试获取指定索引节点
	node := list.GetNode(1)
	if node.Data() != 20 {
		t.Errorf("Expected node data at index 1 to be 20, got %v", node.Data())
	}

	// 测试删除元素
	success := list.Remove(1)
	if !success {
		t.Errorf("Failed to remove element at index 1")
	}

	// 验证删除后的链表大小
	if list.Size() != 2 {
		t.Errorf("Expected list size to be 2 after removal, got %d", list.Size())
	}

	// 验证删除后的节点数据
	node = list.GetNode(1)
	if node.Data() != 30 {
		t.Errorf("Expected node data at index 1 to be 30 after removal, got %v", node.Data())
	}
}
