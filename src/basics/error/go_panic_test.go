package go_error

import (
	"fmt"
	"testing"
)

func TestPanicTest(t *testing.T) {
	var dict map[string]int
	dict["a"] = 1
}

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
