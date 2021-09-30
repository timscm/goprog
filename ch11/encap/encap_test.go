// 对行为的封装
package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

// 该方式：在实例对应方法被调用时，实例的成员会进行值复制
func (e Employee) String1() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// 避免内存拷贝，使用该方式
func (e *Employee) String2() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := &Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String1())
	t.Log(e.String2()) // 与e具有一样的指针值
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) // 返回指针
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)               // {0 Bob 20}
	t.Log(e1)              // {Mike 30}
	t.Log(e2)              // &{2 Rose 22}
	t.Logf("e is %T", e)   // encap.Employee
	t.Logf("e2 is %T", e2) // *encap.Employee
}
