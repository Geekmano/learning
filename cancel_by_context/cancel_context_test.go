package cancel_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCancelContext(t *testing.T) {
	cancel, cancelFunc := context.WithCancel(context.Background())
	for i:=0;i<5;i++{
		go func(i int,ctx context.Context) {
			for {
				if isCancel(ctx){
					break
				}

			}
			fmt.Println(i,"cancelled")
		}(i,cancel)
	}
		cancelFunc()
	time.Sleep(time.Second)
}
func isCancel(context context.Context)bool{
	select {
	case <-context.Done():
		return true
	default:
		return false

	}
}