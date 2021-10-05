package cancelbyclose_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	var wg sync.WaitGroup
	cancelChan := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
			wg.Done()
		}(i, cancelChan)
	}
	// cancel_1(cancelChan)
	cancel_2(cancelChan)
	wg.Wait()
}
