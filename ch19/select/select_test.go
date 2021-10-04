// 多路选择： select，case，default
package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Microsecond * 500)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	// retStr <- chanName (chanName 写入到 retStr)
	// chanName <- retStr (retStr 写入到 chanName)
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Microsecond * 100):
		t.Error("time out")
	}
}
