package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建一个 sync.Map 实例
	// sync.Map 是线程安全的 map，适用于多个 goroutine 并发读写的场景
	var sm sync.Map

	// 1. Store: 存储键值对
	fmt.Println("=== Store 操作 ===")
	sm.Store("name", "张三")
	sm.Store("age", 25)
	sm.Store("city", "北京")
	fmt.Println("已存储: name=张三, age=25, city=北京")

	// 2. Load: 读取值
	fmt.Println("\n=== Load 操作 ===")
	if value, ok := sm.Load("name"); ok {
		fmt.Printf("name = %v\n", value)
	}
	if value, ok := sm.Load("age"); ok {
		fmt.Printf("age = %v\n", value)
	}
	// 读取不存在的键
	if value, ok := sm.Load("email"); ok {
		fmt.Printf("email = %v\n", value)
	} else {
		fmt.Println("email 不存在")
	}

	// 3. LoadOrStore: 如果键存在则返回现有值，否则存储新值
	fmt.Println("\n=== LoadOrStore 操作 ===")
	// 键已存在，返回现有值
	if actual, loaded := sm.LoadOrStore("name", "李四"); loaded {
		fmt.Printf("name 已存在，值为: %v\n", actual)
	}
	// 键不存在，存储新值
	if actual, loaded := sm.LoadOrStore("email", "zhangsan@example.com"); !loaded {
		fmt.Printf("email 不存在，已存储新值: %v\n", actual)
	}

	// 4. LoadAndDelete: 读取并删除
	fmt.Println("\n=== LoadAndDelete 操作 ===")
	if value, loaded := sm.LoadAndDelete("age"); loaded {
		fmt.Printf("已删除 age，其值为: %v\n", value)
	}
	// 验证是否已删除
	if _, ok := sm.Load("age"); !ok {
		fmt.Println("age 已被删除")
	}

	// 5. Delete: 删除键值对
	fmt.Println("\n=== Delete 操作 ===")
	sm.Delete("city")
	fmt.Println("已删除 city")

	// 6. Range: 遍历所有键值对
	fmt.Println("\n=== Range 遍历 ===")
	sm.Store("country", "中国")
	sm.Store("language", "Go")
	fmt.Println("当前所有键值对:")
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v = %v\n", key, value)
		return true // 返回 true 继续遍历，返回 false 停止遍历
	})

	// 7. 并发场景示例
	fmt.Println("\n=== 并发操作示例 ===")
	var wg sync.WaitGroup
	var concurrentMap sync.Map

	// 启动多个 goroutine 并发写入
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", id)
			value := fmt.Sprintf("value_%d", id)
			concurrentMap.Store(key, value)
			fmt.Printf("Goroutine %d: 存储 %s = %s\n", id, key, value)
		}(i)
	}

	// 启动多个 goroutine 并发读取
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", id)
			if value, ok := concurrentMap.Load(key); ok {
				fmt.Printf("Goroutine Reader %d: 读取 %s = %v\n", id, key, value)
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("\n=== 并发操作完成 ===")
	fmt.Println("最终存储的数据:")
	concurrentMap.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v = %v\n", key, value)
		return true
	})
}
