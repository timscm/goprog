package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	// p *Pet  // 复合实现方法
	Pet // 匿名嵌套类型
}

// 以下两个方法，是针对复合实现方法
// 在匿名嵌套类型中，不提供
// func (d *Dog) Speak() {
// 	// d.p.Speak()
// 	fmt.Println("Wang!")
// }

// func (d *Dog) SpeakTo(host string) {
// 	// d.p.SpeakTo(host)
// 	d.Speak()
// 	fmt.Println(" ", host)
// }

// 匿名嵌套类型中声明 Speak 方法
func (d *Dog) Speak() {
	fmt.Println("Wang!")
}

// LSP？
func TestDog(t *testing.T) {
	dog := new(Dog)
	// Pet 的 SpeakTo 调用了 Speak，在Dog中声明了Speak
	// 但是这里的调用依然访问的是 Pet 提供的Speak 方法
	// 打印：... Me, 而不是 Wang! Me
	dog.SpeakTo("Me") // 在匿名嵌套类型中，感觉上继承了Pet的方法

	// 而在这里的 dog 访问Speak，则是直接访问自身的方法
	// 打印： Wang!
	dog.Speak() // 在匿名嵌套类型中，感觉上继承了Pet的方法
}
