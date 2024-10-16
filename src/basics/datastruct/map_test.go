package datastruct_test

import (
	"testing"
)

/*
哈希表 map
*/
func TestMap(t *testing.T) {
	// map 是一种无序的键值对的集合，也称为关联数组、字典或者散列表
	hash0 := make(map[string]int, 10) // make(map[K]V, cap)  cap 是 map 的容量
	hash1 := map[string]int{"a": 1, "b": 2}
	t.Log(hash0, hash1)

	// 遍历 map
	for k, v := range hash1 {
		t.Log(k, v)
	}

	// 通过 key 获取 value
	t.Log(hash1["a"])
}
