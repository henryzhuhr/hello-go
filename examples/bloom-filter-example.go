package main

import (
	"fmt"

	bloomfilter "github.com/henryzhuhr/hello-go/internal/algorithm/bloom-filter"
)

func main() {
	// 预期插入 1000 个元素，误判率为 0.01 (1%)
	// 使用 string 类型，并提供转换函数
	bf := bloomfilter.NewBloomFilter(1000, 0.01, func(s string) []byte {
		return []byte(s)
	})

	// 添加元素
	bf.Add("hello")
	bf.Add("world")
	bf.Add("golang")

	// 检查存在的元素
	fmt.Printf("Contains 'hello': %v\n", bf.Contains("hello"))
	fmt.Printf("Contains 'world': %v\n", bf.Contains("world"))

	// 检查不存在的元素
	fmt.Printf("Contains 'python': %v\n", bf.Contains("python"))
	fmt.Printf("Contains 'java': %v\n", bf.Contains("java"))
}
