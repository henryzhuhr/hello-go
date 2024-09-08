package go_struct

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) Eat()  { fmt.Println("Person  Eat") }
func (p *Person) Walk() { fmt.Println("Person  Walk") }

type Student struct {
	person Person // Student 组合了 Person，Student 也拥有了 Person 的属性和方法
	school string // Student 自己的属性
}

// Eat Student 重写了 Eat 方法
func (s *Student) Eat() { fmt.Println("Student Eat") }

// Fly Student 新增了 Study 方法
func (s *Student) Study() { fmt.Println("Student Walk") }


type Teacher struct {
	Person // 匿名字段，Teacher 继承了 Person 的属性和方法
	school string
}