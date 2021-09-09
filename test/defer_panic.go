package main

import "fmt"

func main() {
	// 输出:
	// func A.
	// panic: Panic in B.
	// goroutine ...
	// 未输出：func C.

	// B 调整增加 defer 后的输出
	// func A.
	// Recover in B.
	// func C.
	A()
	B()
	C()
}

func A() {
	fmt.Println("func A.")
}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B.")
		}
	}()
	panic("Panic in B.")
}

func C() {
	fmt.Println("func C.")
}
