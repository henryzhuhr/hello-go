package datastruct_test

import "testing"

func TestInitArray(t *testing.T) {
	var array1 [3]int        // 定义数组
	array1 = [3]int{1, 2, 3} // 初始化数组

	array2 := [3]int{}             // 定义数组
	array3 := [3]int{1, 2, 3}      // 初始化数组
	array4 := [...]int{1, 2, 3}    // 初始化数组，不指定长度
	array5 := [5]int{1: 10, 3: 30} // 指定下标初始化

	nums := new([5]int) // new创建数组指针

	t.Log(array1, array2, array3, array4, array5, *nums)
}

func TestSplitArray(t *testing.T) {
	array := [5]int{1, 2, 3, 4, 5}
	s1 := array[1:3] // 下标范围 [1, 3) [2, 3]
	s2 := array[:3]  // 下标范围 [0, 3) [0, 1, 2]
	s3 := array[1:]  // 下标范围 [1, 5) [1, 2, 3, 4]
	s4 := array[:]   // 下标范围 [0, 5) [0, 1, 2, 3, 4]

	t.Log(s1, s2, s3, s4)
	// 分割数组后，会转化为切片
	t.Logf("s1 type=%T", s1)
}

func TestArray2Slice(t *testing.T) {
	array := [5]int{1, 2, 3, 4, 5}
	t.Log("array =", array)
	s1 := array[1:3] // 下标范围 [1, 3) [2, 3]
	t.Logf("s type=%T", s1)
	s1[0] = 10 // 修改切片，会影响原数组
	t.Log("array =", array)
}
