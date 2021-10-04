package sharemem_test

import (
	"sync"
	"testing"
	"time"
)

// 对counter的增加，不是线程安全的程序
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	// 1. 创建Mutex对象
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			// 协程退出时，自动释放锁
			defer func() {
				mut.Unlock()
			}()
			// 先加锁再计数
			mut.Lock()
			counter++
		}()
	}
	// 等待协程结束
	time.Sleep(1 * time.Second)
	// 准确的5000
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	// 1. 创建Mutex对象
	var mut sync.Mutex
	var wg sync.WaitGroup

	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1) // 增加一项
		go func() {
			// 协程退出时，自动释放锁
			defer func() {
				mut.Unlock()
			}()
			// 先加锁再计数
			mut.Lock()
			counter++
			wg.Done() // 完成了一项
		}()
	}

	// WaitGroup 所有的Add值Done，才会退出
	wg.Wait()

	// 准确的5000
	t.Logf("counter = %d", counter)
}
