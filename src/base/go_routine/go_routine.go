package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	go newTask()

	// 匿名函数
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new Goroutine : i = %d\n", i)
			time.Sleep(1 * time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
