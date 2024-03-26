package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name string
	sex  string
}

func (h *Human) Eat() {
	fmt.Println("Human Eat")
}
func (h *Human) Walk() {
	fmt.Println("Human Walk")
}

type SuperHuman struct {
	Human     // SuperHuman 组合了 Human，SuperHuman 也拥有了 Human 的属性和方法
	level int // SuperHuman 自己的属性
}

// Eat SuperHuman 重写了 Eat 方法
func (s *SuperHuman) Eat() {
	fmt.Println("SuperHuman Eat")
}

// Fly SuperHuman 新增了 Fly 方法
func (s *SuperHuman) Fly() {
	fmt.Println("SuperHuman Fly: " + strconv.Itoa(s.level))
}

func main() {
	h := Human{"A", "B"}
	h.Eat()
	h.Walk()

	super := SuperHuman{Human{"C", "D"}, 1}
	super.Eat()
	super.Walk()
	super.Fly()
}
