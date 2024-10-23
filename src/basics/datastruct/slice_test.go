package datastruct_test

import (
	"log"
	"testing"
)

// 基本切片创建
func TestBaseSlice(t *testing.T) {
	var slice1 []int // 定义切片，nil切片
	t.Log(slice1, len(slice1), cap(slice1))
	// 初始化切片
	slice1 = []int{1, 2, 3}
	t.Log(slice1, len(slice1), cap(slice1))

	// 使用make创建切片
	slice2 := make([]int, 3, 5) // 创建长度为3，容量为5的切片
	t.Log(slice2, len(slice2), cap(slice2))
}

// 空切片
func TestEmptySlice(t *testing.T) {
	var slice = []int{}
	t.Log(slice, len(slice), cap(slice), slice == nil)

	// 给切片赋值
	slice[0] = 1 // panic: runtime error: index out of range [0] with length 0
}

// nil切片
func TestNilSlice(t *testing.T) {
	var slice []int
	t.Log(slice, len(slice), cap(slice), slice == nil)

	// 给切片赋值
	slice[0] = 1 // panic: runtime error: index out of range [0] with length 0
}

// 切片的扩容
func TestFuncWithSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	t.Log(s, len(s), cap(s))

	// append() 函数向切片追加元素
	s = append(s, 6)
	t.Log(s, len(s), cap(s))
}

// 切片的 Append() 方法
func TestAppend(t *testing.T) {
	slice := []int{1}
	elem1, elem2 := 2, 3
	anotherSlice := []int{4, 5}
	slice = append(slice, elem1, elem2)    // 添加元素
	slice = append(slice, anotherSlice...) // 添加切片
	log.Println(slice)
}

// 切片的 Append() 方法，对于字符串
func TestAppendString(t *testing.T) {
	slice := append([]byte("hello "), "world"...) // 添加字符串
	log.Println(slice)
}

// 向切片插入元素
func TestInsertToSlice(t *testing.T) {
	slice := []int{2, 3}
	// 在头部插入元素
	slice = append([]int{0, 1}, slice...)
	// 在尾部插入元素
	slice = append(slice, 4, 5)
	log.Println(slice)

	// 在中间插入元素
	index, value := 2, 10
	slice = append(slice[:index], append([]int{value}, slice[index:]...)...)
	log.Println(slice)
}

func modify(slice []int) { slice[0] = 10 }

func TestSliceAsFuncParam(t *testing.T) {
	slice := []int{1, 2, 3}
	modify(slice)
	log.Println(slice)
}
