// Go 接口最佳实践
// 1. 倾向于使用小的接口定义，很多接口只包含一个方法
//    type Reader interface{
// 		Read(p []byte) (n int, err error)
//    }
//    type Writer interface {
//      Write(p []byte) (n int, err error)
//    }
// 2. 较大的接口定义，可以由多个小接口定义组合而成
//    type ReadWriter interface {
//	    Reader
//		Writer
//    }
// 3. 只依赖于必要功能的最小接口
//    func StoreData(reader Reader) error {
//      ...
//    }
package empty_interface_test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	// p.(int): type assertion(类型断言)
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	fmt.Println("string", s)
	// 	return
	// }
	// fmt.Println("Unknown Type")

	// typeswitch guard(类型区别)
	switch v := p.(type) {
	case int:
		fmt.Println("Is Integer", v)
	case string:
		fmt.Println("Is String", v)
	default:
		fmt.Println("Is Unknown Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("is string")
}
