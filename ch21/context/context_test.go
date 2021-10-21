package context_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
 * 根 Context : 通过 context.Background() 创建
 * 子 Context : context.WithCancel(parentContext) 创建
 * ctx, cancel := context.WithCancel(context.Background())
 * 当前 Context 被取消时，基于他的子 context 都会被取消
 * 接收取消通知 <-ctx.Done()
 */

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			wg.Done()
		}(i, ctx)
		fmt.Println(i, "Cancelled")
	}
	cancel()
	// time.Sleep(time.Second * 1)
	wg.Wait()
}
