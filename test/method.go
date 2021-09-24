package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
	// name string: 内部包不受影响，但影响外部包的引用
}

// 扩展地层类型的方法
type TZ int

func (tz *TZ) Increment(num int) {
	*tz += TZ(num)
}

func main() {
	fmt.Println("hello")

	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
	fmt.Println(b.Name)

	var c TZ
	c.Print()       // method value
	(*TZ).Print(&c) // method expression
	c.Increment(100)
	fmt.Println(c)
}

func (a TZ) Print() {
	fmt.Println("TZ")
}

// 方法是和类型绑定在一起的
// 不存在方法名重载
func (a *A) Print() {
	a.Name = "aa"
	fmt.Println("A", a.Name)
}

func (b B) Print() {
	b.Name = "bb"
	fmt.Println("B", b.Name)
}
