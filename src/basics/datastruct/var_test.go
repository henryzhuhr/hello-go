package datastruct_test

import (
	"testing"
)

var ( // 定义全局变量与一般程序语言差不多
	gval int = 100
)

/*
定义变量
*/
func TestVar(t *testing.T) {
	var name string   //  var 变量名 变量类型
	short_var := 10   // 简短声明，只能在函数内部使用，不能用于全局变量
	const pi = 3.1415 // 常量
	const (
		n1 = iota // 0 , iota 是一个可以被编译器修改的常量，每次 const 出现时，iota 的值会重置为 0
		n2        // 1
		_         // 2 , 使用 _ 跳过某些值
		n4        // 3
	)
	const ( // 通过 iota 实现枚举
		_  = iota             // 0
		KB = 1 << (10 * iota) // 1 << (10 * 1)
		MB = 1 << (10 * iota) // 1 << (10 * 2)
		GB = 1 << (10 * iota) // 1 << (10 * 3)
		TB = 1 << (10 * iota) // 1 << (10 * 4)
		PB = 1 << (10 * iota) // 1 << (10 * 5)
	)
	t.Log(name, short_var, pi, n1, n2, n4, KB, MB, GB, TB, PB)
}
