package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

// Duck Type
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	// p: 类型 + 数据, p是一个接口变量
	// 类型 -> type GoProgrammer struct {}
	// 数据 -> &GoProgrammer{}
	var p Programmer = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
