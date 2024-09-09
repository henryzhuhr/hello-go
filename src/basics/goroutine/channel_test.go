package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	intCh := make(chan int)
	intCh <- 42
	v := <-intCh
	fmt.Println(v) // Output: 42
	close(intCh)
}

// 从关闭的通道读取数据
func TestReadFromClosedChannel(t *testing.T) {
	intCh := make(chan int, 1)
	intCh <- 1 // 写入数据
	close(intCh)
	n := <-intCh // 读取数据，管道关闭后可以继续读取
	intCh <- 1   // panic: send on closed channel
	fmt.Println(n)
}

// 无缓冲通道的阻塞测试
func TestUnbufferedChannelBlock(t *testing.T) {
	ch := make(chan int) // 创建无缓冲管道
	defer close(ch)
	fmt.Println("write") // 可以输出
	ch <- 123            // 写入数据
	fmt.Println("read")  // 写入时，由于没有接收者，会在此阻塞，所以不会打印
	n := <-ch            // 读取数据
	fmt.Println(n)
}

// 在协程中读取无缓冲通道测试
func TestReadUnbufferedChannelInCoroutine(t *testing.T) {
	ch := make(chan int) // 创建无缓冲管道
	defer close(ch)
	go func() {
		n := <-ch // 协程 读取数据
		fmt.Println(n)
	}()
	ch <- 6 // 写入数据时会阻塞，直到协程准备好后可以读取，因此不会产生死锁
}

// 在协程中写入无缓冲通道测试
func TestWriteUnbufferedChannelInCoroutine(t *testing.T) {
	ch := make(chan int) // 创建无缓冲管道
	defer close(ch)
	go func() { ch <- 6 }() // 协程 写入数据
	n := <-ch               // 读取数据会阻塞，直到协程写入数据后可以读取，因此不会产生死锁
	fmt.Println(n)
}

// 在协程中同时读写无缓冲通道测试
func TestReadWriteUnbufferedChannelInCoroutine(t *testing.T) {
	ch := make(chan int) // 创建无缓冲管道
	defer close(ch)
	go func() { ch <- 6 }() // 协程 写入数据
	go func() {
		n := <-ch // 协程 读取数据
		fmt.Println(n)
	}()
	time.Sleep(1 * time.Millisecond) // 等待协程执行完毕
}

// 读取空的有缓冲通道阻塞测试
func TestReadEmptyBufferedChannelBlock(t *testing.T) {
	ch := make(chan int, 1) // 创建有缓冲管道
	defer close(ch)
	fmt.Println("read")
	n := <-ch              // 试图读取数据，但是管道为空，会在此阻塞
	fmt.Println("get:", n) // 由于上一步阻塞，所以不会打印
}

// 写入满的有缓冲通道阻塞测试
func TestWriteFullBufferedChannelBlock(t *testing.T) {
	ch := make(chan int, 1) // 创建有缓冲管道
	defer close(ch)
	ch <- 6 // 写入数据
	fmt.Println("write:", 6)
	ch <- 10                  // 写入数据，由于管道已满，会在此阻塞
	fmt.Println("write:", 10) // 由于上一步阻塞，所以不会打印
}

// 在协程中读取有缓冲通道测试
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
