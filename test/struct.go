package main

import "fmt"

type test struct{}

type person struct {
	Name string
	Age  int
}

type newperson struct {
	Name    string
	Age     int
	Contact struct { // 匿名结构的应用，确实非常的方便
		Phone, City string
	}
}

// 潜入结构体
type human struct{ Sex int }
type person2 struct {
	human
	Name string
	Age  int
}

func main() {
	// 组合结构体的使用
	a2 := person2{Name: "joe", Age: 19, human: human{Sex: 9}}
	a2.Name = "joe222"
	a2.Age = 18
	a2.human.Sex = 100
	fmt.Println(a2)

	a := test{}
	fmt.Println("struct test.", a)
	b := person{}
	b.Name = "joe"
	b.Age = 19
	fmt.Println(b)

	c := person{Name: "joe2", Age: 19}
	fmt.Println(c)

	A(c)
	fmt.Println(c)

	B(&c)
	fmt.Println(c)

	// 初始化时，就取地址，但是对属性的修改，并不需要取地址
	d := &person{
		Name: "joe3",
		Age:  20, // 换行了，就需要在末尾带上逗号
	}
	B(d)
	fmt.Println(d)

	// 还是可以使用d.Name来操作
	d.Name = "ok"
	fmt.Println(d)

	// 定义一个匿名结构
	e := struct {
		Name string
		Age  int
	}{
		Name: "joe",
		Age:  19, // 注意末尾的,是必须的
	}
	fmt.Println(e)

	// 匿名字段内的值，不能直接初始化，需要使用.格式
	f := newperson{Name: "joe", Age: 19}
	f.Contact.City = "beijing"
	f.Contact.Phone = "111111"
	var g newperson = f // 不会影响f的值，这里是值拷贝
	g.Name = "ok..."
	fmt.Println(f, g)

}

// 值拷贝
func A(per person) {
	per.Age = 13
	fmt.Println("A", per)
}

// 引用拷贝
func B(per *person) {
	per.Age = 14
	fmt.Println("B", per)
}
