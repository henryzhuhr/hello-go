package main

import "fmt"

const LIM = 41

var fibs [LIM]uint64

type MyStruct struct {
	i int
}

// 函数的参数传递
func myFunction1(i int, arr [2]int) {
	fmt.Printf("in myfunciton1 - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
}

func myFunction2(
	ptr1 MyStruct, // 传递 结构体 时	   会拷贝结构体中的全部内容；
	ptr2 *MyStruct, // 传递 结构体指针 时	会拷贝结构体指针；
) {
	ptr1.i++
	ptr2.i++
	fmt.Printf("in myfunciton2 - ptr1=(%v, %p) ptr2=(%v, %p)\n", ptr1, &ptr1, ptr2, &ptr2)
}

// 函数的参数类型、返回值类型，多返回值需要给出类型
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// 命名（多）返回值
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3 // 需要显式返回
	return // 不需要显式返回
}

// 函数的可变参数
func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func main() {
	// 1. 传值
	// 传值是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数
	i := 30
	arr := [2]int{66, 77}
	fmt.Printf("before calling - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
	myFunction1(i, arr)
	fmt.Printf("after  calling - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
	fmt.Println()

	// 结构体和指针
	// 传递 结构体 时	   会拷贝结构体中的全部内容；
	// 传递 结构体指针 时	会拷贝结构体指针；
	ptr1 := MyStruct{i: 30}
	ptr2 := &MyStruct{i: 40}
	fmt.Printf("before calling - ptr1=(%v, %p) ptr2=(%v, %p)\n", ptr1, &ptr1, ptr2, &ptr2)
	myFunction2(ptr1, ptr2)
	fmt.Printf("after  calling - ptr1=(%v, %p) ptr2=(%v, %p)\n", ptr1, &ptr1, ptr2, &ptr2)

	numx2, numx3 := getX2AndX3(10)
	fmt.Println(numx2, numx3)
	numx2, numx3 = getX2AndX3_2(10)
	fmt.Println(numx2, numx3)

	myfunc(1, 2, 3, 4, 5) // 传递任意个数的 int 值

	// 闭包
	// 闭包是指有权访问另一个函数作用域中的变量的函数，创建闭包的最常见的方式就是在一个函数内部创建另一个函数，通过另一个函数访问这个函数的局部变量
	// 闭包是由函数及其相关引用环境组合而成的实体（即：闭包=函数+引用环境）
	// 闭包是引用了自由变量的函数
	fplus := func(x, y int) int {
		return x + y
	}
	fmt.Println(fplus(3, 4))

	// 通过内存缓存来提升性能
	fibonacci(i)
}

func fibonacci(n int) (res uint64) {
	// 内存缓存与动态规划 https://learnku.com/go/t/38628
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
