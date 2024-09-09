---
outline: [3,6]
date: 2024-09-08
---

## 并发

并发 (Concurrency) 是指程序中有多个独立的执行路径，而并行 (Parallelism) 是指程序中有多个独立的执行路径同时执行。

Go 语言的并发模型是基于 CSP (Communicating Sequential Processes) 模型的，通过 `goroutine` 和 `channel` 实现并发。

Go 官方的 [_Effective Go: Concurrency_](https://go.dev/doc/effective_go#concurrency) 也对此进行了详细的介绍。

### 协程

Go 语言的并发模型是基于协程 (Coroutine) 的，Go 语言的协程被称为 `goroutine`。`goroutine` 是 Go 语言的一个重要特性，它是一种轻量级的线程，由 Go 语言的运行时管理。


`goroutine` 的创建非常简单，只需要在函数调用前加上 `go` 关键字即可。

```go
func hello() {
	fmt.Println("Step 3")
}
func TestGoroutine(t *testing.T) {
	fmt.Println("STATR")
    // 普通函数
	go fmt.Println("Step 1")
	go fmt.Println("Step 2")  
	go hello()
    // 匿名函数
	go func() {
		fmt.Println("Step 4")
	}()
    time.Sleep(10 * time.Microsecond)
	fmt.Println("END")
}
// STATR
// Step 4
// Step 1
// Step 2
// END
// Step 3
```

从上述的输出可以看出：
- 协程的执行顺序是不确定的
- 由于系统创建协程需要时间，导致了 `END` 甚至会在其他 `Step` 之前输出
- 不过在 `END` 之前添加了一个 `time.Sleep` ，使得 `END` 后才输出 `Step 1` ，这说明，协程的创建时间大致在微秒级别（仅在本次运行环境下），这个时间是不确定的，不同的环境下可能会有所不同，但是可以通过这个方法测试一下创建协程的大致时间。

> [!TIP]
> 需要注意的是，协程启动的函数不允许有返回值
> ```go
> go make([]int, 10)	
> // 报错: go discards result of make([]int, 10) (value of type []> int)
> ```

### 协程的退出时机

刚才是在测试函数中创建协程，因此可以完整执行，但是如果在 `main` 函数中创建协程，情况会有一点不一样，如下

```go
func main() {
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
	fmt.Println("end")
}
// start
// end
```

输出没有如何数字，这是因为 `main` 函数执行完毕后，程序就会退出，因此没有等待其他协程执行完毕，所以没有输出。

所以可以添加一个 `time.Sleep` 来等待协程执行完毕

```go
func main() {
    fmt.Println("start")
    for i := 0; i < 10; i++ {
        go fmt.Println(i)
    }
    time.Sleep(10 * time.Microsecond) // [!code ++]
    fmt.Println("end")
}
// start
// 9
// 5
// 6
// 7
// end
```
不过，可以看出也只是随机输出了部分，因为协程的执行时间不确定，所以 `time.Sleep` 不是一个可靠的方法

所以需要一些并发控制机制，以确保协程能按照预期执行。 Go 提供了多种并发控制机制（同步原语）：
- 管道 `channel`
- `WaitGroup` (`sync` 包)
- `Context` (`context` 包)

### 管道 channel

#### 管道的基本操作

管道(channel)是 Go 语言中的一种并发原语，可以用于协程间的通信和并发控制

- **管道声明**

管道通过关键字 `chan` 声明，管道的类型是管道中元素的类型。声明管道的语法如下：
```go
var ch chan Type
```

声明一个类型为 `Type` 的管道 `ch`，`Type` 可以是任意类型，包括函数类型、结构体类型、接口类型等。声明的管道还未初始化，因此其值为 `nil`。

- **管道初始化**

管道的初始化需要使用 `make` 函数，`make` 函数定义如下：
```go
func make(t Type, size ...int) Type
```
`size` 为管道的缓冲大小，如果 `size` 为 0 或者省略，则表示无缓冲管道。

创建管道的示例：
```go
ci := make(chan int)        // 无缓冲管道 int 类型
cj := make(chan int, 0)     // 无缓冲管道 / 缓冲区大小为 0
cs := make(chan int, 100)   // 有缓冲管道 缓冲区大小为 100
```

- **管道关闭**

创建管道后需要关闭管道，使用到 `close` 函数，定义如下：
```go
func close(c chan<- Type)
```

关闭管道后，不能再向管道发送数据，会导致 panic，但是可以继续从管道接收数据，如果管道中还有数据，可以继续读取，直到管道中的数据读取完毕。
```go
intCh := make(chan int, 1)
intCh <- 1      // 写入数据
close(intCh)
n := <-intCh    // 读取数据，管道关闭后可以继续读取
intCh <- 1      // panic: send on closed channel
```


通常可以使用 `defer` 关键字来延迟关闭管道，确保在函数退出时关闭管道。
```go
ch := make(chan int)
// ... 操作管道
defer close(ch)
```

#### 管道的读写

管道的读写操作需要使用 `<-` 操作符，`<-` 操作符的方向表示数据的流向：
- `ch <-` 用于发送数据到管道
- `<- ch` 用于从管道接收数据

管道的读写如下
```go
ch <- value // 发送数据
value := <-ch // 接收数据
```

对于读取操作还可以返回第二个参数，表示是否读取成功
```go
value, ok := <-ch
```
如果管道关闭了，但是管道中还有数据，那么也可以读取成功，即 `ok` 为 `true`

管道的数据流向是先进先出 (FIFO) 的，即先发送的数据先接收。


> [!TIP]
> 管道的读写需要考虑管道是否有缓冲区，因此在读写操作时需要考虑管道的状态，否则可能会导致死锁，下一节会详细介绍。

#### 管道的缓冲区

- **无缓冲管道 (unbuffered channel)** 

**无缓冲管道**缓冲区容量为 0 ，不会存储任何数据，发送和接收操作是同时进行的，因此：
- 发送者会阻塞，直到接收者准备好
- 接收者会阻塞，直到发送者准备好

例如这个例子，就会产生死锁
```go
func main() {
	ch := make(chan int) // 创建无缓冲管道
	defer close(ch)
    fmt.Println("write") // 可以输出
	ch <- 123   // 写入数据
    fmt.Println("read") // 写入时，由于没有接收者，会在此阻塞，所以不会打印
	n := <-ch   // 读取数据
	fmt.Println(n)
}
// write
```

为了解决上述问题，可以使用协程来解决

::: code-group

```go [用协程读取数据]
// 用协程读取数据
func main() {
    ch := make(chan int) // 创建无缓冲管道
    defer close(ch)
    go func() {
        n := <-ch // 协程 读取数据
        fmt.Println(n)
    }()
    ch <- 6 // 写入数据时会阻塞，直到协程准备好后可以读取，因此不会产生死锁
}
```

```go [用协程写入数据]
// 用协程写入数据
func main() {
	ch := make(chan int) // 创建无缓冲管道
	defer close(ch)
	go func() { ch <- 6 }() // 协程 写入数据
	n := <-ch   // 读取数据会阻塞，直到协程写入数据后可以读取，因此不会产生死锁
	fmt.Println(n)
}
```

```go [用两个协程来读写数据]
// 用两个协程来读写数据
func main() {
	ch := make(chan int) // 创建无缓冲管道
	defer close(ch)
	go func() { ch <- 6 }() // 协程 写入数据
	go func() {
		n := <-ch // 协程 读取数据
		fmt.Println(n)
	}()
	time.Sleep(1 * time.Second) // 等待协程执行完毕
}
```

:::

无论哪种方式，都可以解决死锁问题。因此，无缓冲管道是同步的，也被称为**同步管道**

- **有缓冲管道 (buffered channel)**

**有缓冲管道**缓冲区容量大于 0 ，可以存储一定数量的数据，因此：
- 读取空管道会阻塞
- 写入满管道会阻塞

读取空管道会阻塞，例如
```go
func main() {
	ch := make(chan int, 1) // 创建有缓冲管道
	defer close(ch)
	fmt.Println("read")
	n := <-ch              // 试图读取数据，但是管道为空，会在此阻塞
	fmt.Println("get:", n) // 由于上一步阻塞，所以不会打印
}
// read
```

写入满管道会阻塞，例如
```go
func main() {
	ch := make(chan int, 1) // 创建有缓冲管道
	defer close(ch)
	ch <- 6  // 写入数据
	fmt.Println("write:", 6)
	ch <- 10 // 写入数据，由于管道已满，会在此阻塞
	fmt.Println("write:", 10) // 由于上一步阻塞，所以不会打印
}
// write: 6
```

内置函数 `len(ch)` 返回管道中的元素个数，`cap(ch)` 返回管道的缓冲区大小，下面一个例子展示了有缓冲管道的读写操作
```go
func TestReadBufferedChannelInCoroutine(t *testing.T) {
	ch := make(chan int, 3)                      // 创建有缓冲管道
	chw, chr := make(chan bool), make(chan bool) // 用于同步的管道
	defer func() { close(ch); close(chw); close(chr) }()
    // 写协程
	go func() { 
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("write:", i, "len:", len(ch), "cap:", cap(ch))
		}
		chw <- true // 循环 写完后通知读协程
	}()
	// 读协程
	go func() {
		for i := 0; i < 5; i++ {
			n := <-ch
			fmt.Println("read :", n, "len:", len(ch), "cap:", cap(ch))
		}		
		chr <- true // 循环 读完后通知写协程
	}()
	fmt.Println("read  done", <-chr) // 阻塞，等待读协程完成
	fmt.Println("write done", <-chw) // 阻塞，等待写协程完成
}
```

输出如下
```bash
write: 0 len: 1 cap: 3
write: 1 len: 2 cap: 3
write: 2 len: 3 cap: 3   // 缓冲区已满，阻塞
read : 0 len: 3 cap: 3   // 读取数据后，缓冲区有空间
read : 1 len: 2 cap: 3
read : 2 len: 1 cap: 3
read : 3 len: 0 cap: 3
write: 3 len: 3 cap: 3   // 缓冲区有空间，继续写入
write: 4 len: 0 cap: 3
read  done {}            // 读协程完成
read : 4 len: 0 cap: 3
write done {}            // 写协程完成
```

因此，有缓冲管道是异步的，也被称为**异步管道**

缓冲区大小为1的管道，可以用来实现一个简单的**互斥锁**

#### nil 管道

特别地，`nil` 管道无论如何读写都会阻塞
```go
var ch chan int
ch <- 6 // 阻塞
<-ch    // 阻塞
```

关闭一个 `nil` 管道会导致 panic
```go
var ch chan int
close(ch) // panic: close of nil channel
```

#### 管道的并发安全

