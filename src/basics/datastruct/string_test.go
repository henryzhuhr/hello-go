package datastruct_test

import (
	"testing"
)

/*
字符串 string
*/
func TestString(t *testing.T) {
	// 字符串是一种值类型，且值不可变，即创建某个文本后你无法再次修改这个文本的内容，而是新建了一个文本

	// 1. 使用双引号创建字符串
	str1 := "hello world"
	t.Log(str1)

	// 2. 使用反引号创建多行字符串
	str2 := `hello
	world`
	t.Log(str2)

	// 类型转换
	// 当我们使用 Go 语言解析和序列化 JSON 等数据格式时，经常需要将数据在 string 和 []byte 之间来回转换
	// 当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。
	// Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾
	// 1. string 转 []byte
	str3 := "hello world"
	// bytes := []byte(str3) // 和下面的等价
	bytes := []rune(str3)
	t.Log(bytes)

	// 2. []byte 转 string
	bytes[0] = 'H'
	str4 := string(bytes)
	t.Log(str4)
}
