package datastruct_test

import (
	"testing"
)

/*
切片 slice
*/
func TestSlice(t *testing.T) {
	// 切片的长度是动态的，所以声明时只需要指定切片中的元素类型

	// 1. 使用下标创建切片，是最原始也最接近汇编语言的方式，它是所有方法中最为底层的一种
	// 编译器会将 arr[0:3] 或者 slice[0:3] 等语句转换成 OpSliceMake 操作
	arraytmp := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := arraytmp[0:3]
	t.Log(slice1)

	// 2. 通过字面量创建切片
	slice2 := []int{1, 2, 3, 4, 5} // 这不同于数组的声明方式，数组的声明需要指定长度，或者 ...
	t.Log(slice2)

	// 3. 使用 make 创建切片，make 是一个内建函数，它的主要作用是创建切片、哈希表和 Channel 等内建的数据结构
	// 运行期间创建，返回的是类型的引用
	// 切片逃逸（过大）时，会在堆上分配内存，否则在栈上分配内存
	slice3 := make([]int, 5, 10) // make([]T, size, cap)  size 是切片的长度，cap 是切片的容量
	t.Log(slice3)

	// 使用 len() 和 cap() 获取切片的长度和容量
	t.Log(len(slice3), cap(slice3))

	// 4. 使用 append() 函数向切片中添加元素
	slice3 = append(slice3, 1, 2, 3, 4, 5)
	t.Log(slice3)
}
