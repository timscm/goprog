// channel 的关闭
// 1. 向关闭的 channel 发送数据，会导致 panic
// 2. v, ok := <-ch, ok==true 表示通道可用且正常接受，false表示通道关闭
// 3. 所有的 channel 接受者都会在channel关闭时，立刻从阻塞等待中返回上述ok值为false，
// 这个广播机制常被利用，进行向多个订阅者同时发送信号
// 如：退出信号
package channelclose_test

import (
	"fmt"
	"sync"
	"testing"
)

func dataProduct(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // 关闭channel
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			// data, ok:= <- ch
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProduct(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
