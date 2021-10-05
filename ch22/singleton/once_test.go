package once_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var once sync.Once
var singleInstance *Singleton

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Singleton obj.")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestOnce(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
