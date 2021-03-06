package custom_type

import (
	"fmt"
	"testing"
	"time"
)

// 自定义类型
type IntConv func(n int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(n int) int {
	time.Sleep(time.Second * 1)
	return n
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
