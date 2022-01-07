package some_case

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestRuntime(t *testing.T) {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下， 再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}
func TestRuntime2(t *testing.T) {
	go func() {
		 defer fmt.Println("A.defer")
		 func() {
			 defer fmt.Println("B.defer")
			 // 结束协程
			 runtime.Goexit()
			 defer fmt.Println("C.defer")
			 fmt.Println("B")
			 }()
		 fmt.Println("A")
		 }()
	 time.Sleep(5*time.Second)
}