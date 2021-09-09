package main

import "fmt"

func main() {
	fmt.Println("a")
	// 后定义的先调用，先输出c, 后输出b
	defer fmt.Println("b")

	// 出现了panic错误，此时输出a,-1, b
	// 即上面defer在main的panic发出之前，执行了。
	// for i := -1; i < 3; i++ {
	// 	fmt.Println(1 / i)
	// }
	defer fmt.Println("c")

	// 后定义的先调用，输出：2,1,0
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	// 可以调用函数，但是输出的值为：3，3，3
	// 原因：i 在函数内是一个引用值，在执行时才会访问变量i
	// 而 defer 是要在函数 main 退出后才会执行，此时i已经
	// 在 main 函数中被赋值为 3了。
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	fmt.Println("exit main.")
}
