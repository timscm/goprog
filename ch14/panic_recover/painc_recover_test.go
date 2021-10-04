package panicrecover_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		fmt.Println("Finally!")
	}()

	// 最常见的“错误恢复”
	// 当心！ recover 成为恶魔
	// 1. 形成僵尸服务进程，导致 health check 失效
	// 2. "Let it Crash!" 往往是我们恢复不确定错误的最好方法
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()

	fmt.Println("Start")
	// os.Exit 退出时不会调用 defer 指定的函数
	// os.Exit 退出时不会输出当前调用栈信息
	// os.Exit(-1)
	panic(errors.New("something wrong."))
}
