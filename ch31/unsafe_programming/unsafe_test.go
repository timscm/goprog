package unsafeprogramming_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type Customer struct {
	Name string
	Age  int
}

func TestUnsafe(t *testing.T) {
	i := 10
	// !danger
	f := *(*float64)(unsafe.Pointer(&i))
	t.Log(f)
}

type MyInt int

func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	//合理的转换
	var b []MyInt = *(*[]MyInt)(unsafe.Pointer(&a))
	t.Logf("%T, %v\n", b, b)
}

func TestAtomic(t *testing.T) {
	var shareBufferPtr unsafe.Pointer
	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufferPtr, unsafe.Pointer(&data))
	}
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufferPtr)
		fmt.Println(data, *(*[]int)(data))
	}

	var wg sync.WaitGroup
	writeDataFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
