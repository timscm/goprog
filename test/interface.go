package main

import "fmt"

type USB interface {
	Name() string
	// 可以将该方法放在嵌入接口内：Connect()
	Connector
}

type Connector interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connect:", pc.name)
}

// 使用空接口，类似顶层
// func Disconnect(usb USB) {
func Disconnect(usb interface{}) {
	// 判断类型， ok pattern
	// if pc, ok := usb.(PhoneConnector); ok {
	// 	fmt.Println("Disconnected:", pc.name)
	// 	return
	// }
	// fmt.Println("Unknown device.")

	// type switch 方式替换上面的模式
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown device.")
	}
}

func main() {
	fmt.Println("Hello")
	var a USB

	// ERROR：a.name 在USB接口内是不存在的属性
	// a = PhoneConnector{}
	// a.name = "PhoneConnector"

	a = PhoneConnector{"PhoneConnector"}

	a.Connect()
	Disconnect(a)

	// 父子接口的转换问题
	var c Connector
	c = Connector(a)
	c.Connect()

	// var d USB = USB(c) 报错：不能将子接口转换为父接口
}
