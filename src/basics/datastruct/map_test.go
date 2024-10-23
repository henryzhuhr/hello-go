package datastruct_test

import (
	"testing"
)

// 创建map
func TestCreateMap(t *testing.T) {
	mp1 := map[string]int{"a": 0, "b": 1, "c": 2, "d": 3}
	mp2 := make(map[string]int, 10)

	t.Logf("map: %v", mp1)
	t.Logf("map: %v", mp2)
}

func TestMapOperation(t *testing.T) {
	mp := map[string]int{"a": 0, "b": 1}

	// 获取元素
	t.Logf("map: %v", mp)
	t.Logf("map['a']: %v", mp["a"])
	t.Logf("map['c']: %v", mp["c"])

	// 添加元素`	`
	mp["c"] = 4
	t.Logf("map: %v", mp)

	// 删除元素
	delete(mp, "c")
	t.Logf("map: %v", mp)

	// 判断元素是否存在
	if val, exist := mp["c"]; exist {
		t.Logf("map['c']: %v", val)
	} else {
		t.Logf("map['c'] not exist")
	}

	clear(mp)
}
