package main

import "fmt"

// Go 语言不是一种 “传统” 的面向对象编程语言：它里面没有类和继承的概念。
// 但是 Go 语言里有非常灵活的 接口 概念，通过它可以实现很多面向对象的特性
type Book struct {
	title  string  // 小写字母开头的字段，表示这个字段是私有的，外部包不可以访问
	Price  float32 // 大写字母开头的字段，表示这个字段是公有的，外部包可以访问
	author string  `json:"title"` // 使用反引号定义 tag，可以在运行时通过反射的机制读取出来
}

// Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个person的构造函数。
// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func newBook(title string, price float32, author string) *Book {
	return &Book{
		title:  title,
		Price:  price,
		author: author,
	}
}

// (b *Book) 是接收者，表示这个方法是 Book 类型的方法
// 使用 * 作为 Book 的接收者，表示这个方法是 Book 的指针类型方法，只有指针类型的接收者才能修改结构体的值
func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func main() {
	var book Book
	newBook := newBook("Go Programming", 100, "George")
	book.title = newBook.title
	// book := Book{title: "Go Programming"}
	fmt.Println(book.GetTitle())
	book.SetTitle("Go Programming Language")
	fmt.Println(book.GetTitle())
}
