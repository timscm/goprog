package main

import (
	"fmt"
)

func main() {
	f := clouse(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func clouse(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
