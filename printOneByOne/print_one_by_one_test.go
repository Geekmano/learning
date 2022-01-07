package printOneByOne

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

var wg sync.WaitGroup
func TestPrint(t *testing.T) {
	n:=4
	ch1:=make(chan int)
	ch2:=make(chan int)
	ch3:=make(chan int)
	go first(ch1,ch2,n)
	go second(ch2,ch3,n)
	go third(ch3,ch1,n)
	ch1<-1
	wg.Wait()

}
func first(ch1,ch2 chan int, n int) {
	wg.Add(1)
	defer wg.Done()
	for i:=0;i<n;i++{
		<-ch1
		fmt.Println("first")

		ch2<-1
	}

}
func second(ch2,ch3 chan int,n int) {
	wg.Add(1)
	defer wg.Done()
	for i:=0;i<n;i++{
		<-ch2
		fmt.Println("second")
		ch3<-1
	}

}
func third(ch3,ch1 chan int,n int) {
	wg.Add(1)
	defer wg.Done()
	for i:=0;i<n;i++{
		<-ch3
		fmt.Println("third")
		if i!=n-1{
			ch1<-1
		}

	}

}
func TestPrint2(t *testing.T) {
	runtime.GOMAXPROCS(1)
	//n:=4;
	var wg sync.WaitGroup
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			fmt.Println("first")
			runtime.Gosched()
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			fmt.Println("second")
			runtime.Gosched()
		}
	}(&wg)
	wg.Wait()

}