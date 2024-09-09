package main

import (
	"fmt"
	"time"
	// "sync"
	// "time"
)

// func main() {
// 	fmt.Println("start")
// 	for i := 0; i < 10; i++ {
// 		go fmt.Println(i)
// 	}
// 	time.Sleep(10 * time.Microsecond) // [!code ++]
// 	fmt.Println("end")
// }

// func main() {
// 	var wg sync.WaitGroup
// 	fmt.Println("start")
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			fmt.Println(i)
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// 	fmt.Println("end")
// }

func main() {
	ch := make(chan int) // 创建无缓冲管道
	defer close(ch)
	go func() {
		ch <- 6 // 写入数据
	}()
	go func() {
		n := <-ch // 读取数据
		fmt.Println(n)
	}()
	time.Sleep(1 * time.Second)
}
