package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个无缓冲的channel
	ch := make(chan int)

	go func() {
		defer close(ch)
		ch <- 1 // 发送数据
	}()

	num := <-ch // 接收数据
	fmt.Println(num)

	chWithBuffer := make(chan int, 3) // 创建一个容量为3的有缓冲区的channel

	go func() {

		defer fmt.Println("子goroutine结束")

		for i := 0; i < 3; i++ {
			chWithBuffer <- i
			fmt.Println("子goroutine正在写入=", i, " len=", len(chWithBuffer), " cap=", cap(chWithBuffer))
			//time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-chWithBuffer
		fmt.Println("子goroutine正在读取=", num, " len=", len(chWithBuffer), " cap=", cap(chWithBuffer))
	}
	defer close(chWithBuffer)
	fmt.Println("main finished")

}
