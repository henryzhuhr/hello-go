package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func NewTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
		if i >= 5 {
			break
		}
	}
}

func Hello() {
	fmt.Println("Step 3")
}
func TestGoroutine(t *testing.T) {
	fmt.Println("STATR")
	go fmt.Println("Step 1")
	go fmt.Println("Step 2")
	go Hello()
	go func() {
		fmt.Println("Step 4")
	}()
	time.Sleep(5 * time.Microsecond)
	fmt.Println("END")

	// go make([]int, 10)	// go discards result of make([]int, 10) (value of type []int)
}
