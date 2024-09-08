---
outline: [3,6]
date: 2024-09-03
---

## 错误处理

### Go 的错误处理

Go 语言中的错误处理机制与其他语言有很大的不同。Go 语言中的错误处理机制是通过 `panic` 和 `recover` 来实现的。

Go 的错误主要有以下几种类型：
- `error`: 部分流程出错，需要处理
- `panic`: 很严重的问题，程序应该在处理完问题后立即退出
- `fatal`: 非常致命的问题，程序应该立即退出

> 准确的来说，Go 并没有异常，更多的是通过错误来体现，同样的，Go 中也并没有 `try-catch-finally` 这种语句，Go 创始人希望能够将错误可控，他们不希望干什么事情都需要嵌套一堆 `try-catch` ，所以大多数情况会将其作为函数的返回值来返回

这里有两篇Go团队关于错误处理的文章，感兴趣可以看看

### 错误

错误(Error)是程序运行时的一种错误，是指在该出现问题的地方出现问题，是预料之内的，比如：文件不存在，网络连接失败等，它不会导致程序的异常终止。

#### Go 内建 error

Go 内建一个 [`error`](https://pkg.go.dev/builtin#error) 接口类型作为 Go 的错误标准处理方式，`error` 接口只有一个方法 `Error()`，返回一个字符串，表示错误信息。这个接口定义如下：

```go
// src/builtin/builtin.go
type error interface {
    Error() string
}
```


接口实现示例在 [`src/errors/errors.go`](https://go.dev/src/errors/errors.go) ([Github](https://github.com/golang/go/blob/master/src/errors/errors.go)) 中：

```go
package errors

// 实现 error 接口的 New 函数
func New(text string) error {
    return &errorString{text}
}

// 错误信息结构体
type errorString struct {
    s string
}

// 实现 errorString 的 Error 方法，返回错误信息字符串
func (e *errorString) Error() string {
    return e.s
}
```

#### 错误返回

```go
func doSomething() (result int, err error) { 
    // 某个操作返回值，和错误
    result, err := doSomethingElse()
    // 如果错误不为 nil 说明错误了
    if err != nil {
        return nil, err
    }
    // 如果没有错误，返回结果
    return result, nil
}
```

### 异常

**异常**(Exception)是程序运行时的一种错误，是指在不该出现问题的地方出现问题，是预料之外的，比如空指针引用，下标越界，向空 map 添加键值等，它会导致程序的异常终止。

- 人为制造被自动触发的异常，比如：数组越界，向空 map 添加键值对等。
- 手工触发异常并终止异常，比如：连接数据库失败主动抛出异常，程序终止。
  
Go 语言中没有异常，但是有 `panic` 和 `recover` 机制来处理错误。

#### panic

panic 一词意为“恐慌”，在 Go 语言中，`panic` 用于引发恐慌，导致程序中断执行。`panic` 可以在任何地方引发，但是如果 `panic` 没有被捕获，程序就会崩溃。

例如向 `nil` 的 map 添加元素，会引发 panic：
```go
var dict map[string]int
dict["a"] = 1
// panic: assignment to entry in nil map [recovered]
//     panic: assignment to entry in nil map
```

> [!TIP]
> 只要任一协程发生 `panic` ，如果不将其捕获的话，整个程序都会崩溃

#### panic 创建

当连接数据库失败时，后续操作没有意义，可以使用 `panic` 主动触发异常，程序终止。
```go
func initDataBase(host string, port int) {
	if len(host) == 0 || port == 0 {
		panic("Error connection params")
	}
	// ...其他的逻辑
}
func TestConnectDBPanic(t *testing.T) {
	initDataBase("", 0)
}
// panic: Error connection params [recovered]
// 	panic: Error connection params
```

#### panic 的善后

程序在 `panic` 退出之前会进行一些清理工作，比如关闭文件描述符、释放资源等。Go 语言中提供了 `defer` 机制，可以在函数退出时执行一些清理工作。

```go
func TestDeferPanic(t *testing.T) {
    defer fmt.Println("A")
    fmt.Println("B")
    panic("panic")
    defer fmt.Println("C")
}
// B
// A
// panic: panic
```

可以看出，`defer` 语句在 `panic` 之前执行，`panic` 之后的 `defer` 语句不会执行。

#### recover

当 `panic` 发生时，程序会立即终止，但是可以通过 `recover` 函数捕获 `panic`，并且恢复程序的执行，需要注意，`recover` 函数必须在 `defer` 函数中调用。

```go
func PanicRecoverFunc() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("panic 恢复")
		}
	}()
	panic("发生 panic")
}
func TestRecover(t *testing.T) {
	PanicRecoverFunc()
	fmt.Println("程序正常退出")
}
// 发生 panic
// panic 恢复
// 程序正常退出
```
函数 `PanicRecoverFunc()` 调用者不知道函数内部发生了 `panic`，但是通过 `recover` 函数可以捕获 `panic`，并且恢复程序的执行。

#### 闭包中使用 recover

`recover` 函数只能在 `defer` 函数中调用，如果在闭包中调用 `recover` 函数，只能恢复闭包内部的 `panic`，无法恢复外部函数的 `panic`。

```go
func ClosureRecoverFunc() {
	defer func() {
		func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				fmt.Println("panic 恢复")
			}
		}()
	}()
	panic("发生 panic")
}
func TestClosureRecover(t *testing.T) {
	ClosureRecoverFunc()
	fmt.Println("程序正常退出")
}
// panic: 发生 panic [recovered]
//	panic: 发生 panic
```

闭包函数可以看作调用了一个函数，panic 是**向上传递**而不是向下，自然闭包函数也就无法恢复 panic。

#### recover 总结

总的来说recover函数有几个注意点

- 必须在 `defer` 中使用
- 多次使用也只会有一个能恢复 `panic`
- 闭包 `recover` 不会恢复外部函数的任何 `panic`
- panic的参数禁止使用 `nil`



#### 异常处理的最佳实践

对于真正意外的情况，那些表示不可恢复的程序错误，不可恢复才使用 `panic`。对于其他的错误情况，我们应该是期望使用 `error` 来进行判定

Go 源码很多地方写 panic，但是工程实践业务代码不要主动写 panic，理论上 panic 只存在于 server 启动阶段，比如 config 文件解析失败，端口监听失败等等，所有业务逻辑禁止主动 panic，所有异步的 goroutine 都要用 recover 去兜底处理。




### Go 的错误处理设计理念

- 理解了错误和异常的真正含义，我们就能理解 Go 的错误和异常处理的设计意图。传统的 `try{}catch{}`结构，很容易让开发人员把错误和异常混为一谈，甚至把业务错误处理的一部分当做异常来处理，于是你会在程序中看到一大堆的 catch。

- Go 开发团队认为错误应该明确地当成业务的一部分，任何可以预见的问题都需要做错误处理，于是在Go代码中，任何调用者在接收函数返回值的同时也需要对错误进行处理，以防遗漏任何运行时可能的错误

- 异常则是意料之外的，甚至你认为在编码中不可能发生的，Go 遇到异常会自动触发 panic（恐慌）， 触发 panic 程序会自动退出。除了程序自动触发异常，一些你认为不可允许的情况你也可以手动触发异常

- 另外，在Go 中除了触发异常，还可以终止异常并可选的对异常进行错误处理，也就是说，错误和异常是可以相互转换的



### Go 处理错误的三种方式

#### 经典 Go逻辑

直观的返回 error
```go
// ZooTour struct
type ZooTour interface {
	Enter() error
	VisitPanda(panda *Panda) error
	Leave() error
}

// 分步处理,每个步骤可以针对具体返回结果进行处理
func Tour(t ZooTour1, panda *Panda) error {
	if err := t.Enter(); err != nil {
		return errors.WithMessage(err, "Enter failed.")
	}

	if err := t.VisitPanda(); err != nil {
		return errors.WithMessage(err, "VisitPanda failed.")
	}
	// ...
	return nil
}
```

#### 屏蔽过程中的error 的处理

将 error 保存到对象内部，处理逻辑交给每个方法，本质上仍是顺序执行。标准库的 `bufio`、`database/sql` 包中的 `Rows` 等都是这样实现的，有兴趣可以去看下源码

```go
// ZooTour struct
type ZooTour interface {
	Enter() error
	VisitPanda(panda *Panda) error
	Leave() error
	Err() error
}

func Tour(t ZooTour1, panda *Panda) error {
	t.Enter()
    t.VisitPanda(panda)
    t.Leave()
    // 集中编写业务逻辑代码，最后统一处理错误
    if err := t.Err(); err != nil {
        return errors.WithMessage(err, "Tour failed.")
    }
    return nil
}
```


#### 利用函数式编程†延迟运行

分离关注点 -遍历访问用数据结构定义运行顺序，根据场景选择，如顺序、逆序、二叉树树遍历等。

运行逻辑将代码的控制流逻辑抽离，灵活调整。

```go
type Walker interface {
	Next MyFunc
}	
type SliceWalker struct {
	index int
	funs []MyFunc
}

func NewEnterFunc() MyFunc {
	return func(t ZooTour) error {
		return t.Enter ()
	}
}
func BreakOnError(t ZooTour, walker Walker) error {
	for {
        f := walker.Next()
		if f == nil { break }
	}
	if err := f(t); err := nil {
		// 遇到错误 break 或者 continue 继续执行 
	}
}
```

#### 三种方式对比

上面这三个例子，是 Go项目处理错误使用频率最高的三种方式，也可以应用在 error 以外的处理逻辑。

- case 1：如果业务逻辑不是很清楚，比较推荐 case1

- case 2：代码很少去改动，类似标准库，可以使用 case2

- case 3：比较复杂的场景，复杂到抽象成一种设计模式


### 参考资料

[1][2] 是Go团队关于错误处理的文章

- [Effective Go: Errors](https://go.dev/doc/effective_go#errors)

- [1] [Errors are values - Rob Pike](https://go.dev/blog/errors-are-values)

- [2] [Error handling and Go - Andrew Gerrand](https://go.dev/blog/error-handling-and-go)

- [3] [Golang深入浅出之-Go语言 defer、panic、recover：异常处理机制-腾讯云开发者社区-腾讯云](https://cloud.tencent.com/developer/article/2412169)

- [4] [一套优雅的 Go 错误问题解决方案-腾讯云开发者社区-腾讯云](https://cloud.tencent.com/developer/article/1885094)

- [Go 语言的错误处理机制是一个优秀的设计吗？ - 腾讯技术工程的回答 - 知乎](https://www.zhihu.com/question/27158146/answer/2474680292)

- [Go 错误处理：100+ 提案全部被拒绝，为何现阶段仍用 if err != nil？ - 陈煎鱼的文章 - 知乎](https://zhuanlan.zhihu.com/p/615731644)
