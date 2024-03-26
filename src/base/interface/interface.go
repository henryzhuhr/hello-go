package main

import "fmt"

type person interface { // 接口不能成员变量
	getName() string
	getAge() int
}


// struct 实现接口
// Go 语言中接口的实现都是隐式的，我们只需要实现 getName() string 和 getAge() int 方法就实现了 person 接口。
type student struct {
	name string
	age  int
}

func (s student) getName() string {
	return s.name
}

func (s student) getAge() int {
	return s.age
}

func main() {
	var p person = student{"张三", 20}
	fmt.Println(p.getName())
	fmt.Println(p.getAge())
}
