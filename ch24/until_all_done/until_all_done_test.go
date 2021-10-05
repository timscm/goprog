package untilalldone_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(i int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("the result is from %d", i)
}

func AllResponse() string {
	numOfRunner := 10
	// 使用 buffer channel
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
