package datastruct_test

import (
	"slices"
	"testing"
)

func TestArray(t *testing.T) {
	array1 := [3]int{}             // 定义数组
	array2 := [3]int{1, 2, 3}      // 初始化数组
	array3 := [...]int{1, 2, 3}    // 初始化数组，不指定长度
	array4 := [5]int{1: 10, 3: 30} // 指定下标初始化

	t.Log(array1, array2, array3, array4)

	// 遍历数组
	for i := 0; i < len(array4); i++ {
		t.Log(array4[i], ", ")
	}
	for j, v := range array4 {
		t.Log(j, ":", v, ", ")
	}
	t.Log()
}

// 数组在切割后，就会变为切片类型
func TestSplitArray(t *testing.T) {
	nums := [5]int{1, 2, 3, 4, 5}
	t.Logf("Type: %T\n", nums)
	t.Logf("Type: %T\n", nums[2:4])
}

// 数组转化为切片类型
func TestArrayToSlice(t *testing.T) {
	nums := [5]int{1, 2, 3, 4, 5}
	t.Logf("nums: %v\n", nums)
	// 不带参数进行切片，可以将数组转化为切片
	sli1 := nums[:]

	// 修改切片的值，会影响到原数组
	// 这是因为切片是对数组的引用，是一个指针
	sli1[0] = 0
	t.Logf("nums: %v\n", nums)
	t.Logf("sli1: %v\n", sli1)

	// 如果要对转换后的切片进行修改，建议使用下面这种方式进行转换
	nums2 := [5]int{1, 2, 3, 4, 5}
	sli2 := slices.Clone(nums2[:])
	sli2[0] = 0
	t.Logf("nums2: %v\n", nums2)
	t.Logf(" sli2: %v\n", sli2)
}
