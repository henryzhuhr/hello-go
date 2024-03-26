package main

import (
	"fmt"

	"github.com/henryzhuhr/hello-go/router"
)

func main() {
	r := router.Router()

	r.Run(":9999") // 指定启动端口号
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[ERROR]", err)
		}
	}()
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}
