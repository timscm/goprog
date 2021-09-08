// 单行注释
/*
多行注释
*/

package main // 定义名为 main 的包

// 导入依赖库
// import "fmt" 单行导入
// 合并导入
import (
	// 使用 . 表示别名
	// . "fmt"
	// 使用 std 表示别名
	// std "fmt"
	"fmt"

	"rsc.io/quote"
	// 使用 _ 匿名，只执行包内的 init 函数
	// _ "fmt"
)

// 定义函数
func main() {
	// 函数调用
	fmt.Println(quote.Hello())
	fmt.Println("test import .")
	fmt.Printf("test import %v\n", 12)

	// 定义变量
	var a int = 1
	var e, f int          // 同类型定义，编译器默认分配零值
	value := 1            // 简写格式
	var b, c, d = 1, 2, 3 // 编译器自动推导变量类型，多重赋值功能
	fmt.Println(a, value, b, c, d, e, f)

	// 定义常量
	const g, h int = 1, 2
	fmt.Println(g, h)
	// 常量不可修改
	// g = 2
	// 常量的值是需要在编译阶段就可知的
	const i int = len("123")
	fmt.Println(i)

	// _ 匿名占位符
	_, j := 1, 2
	fmt.Println(j)

	// 运算符: +, -, *, /, %,
	fmt.Println(1+2, 1-2, 2*3, 2.0/3, 1/3.0, 5%3)

	// 位操作：&, |, ~, ^
	//    6: 0110
	//   13: 1101
	// 6&13: 0100 同1为1，其他为0
	// 6|13: 1111 含1为1，全0为0
	// 6^13: 1011 不同为1，相同为0
	fmt.Printf("%b\n", 6)
	fmt.Printf("%b\n", 13)
	fmt.Printf("%b\n", 6&13)
	fmt.Printf("%b\n", 6^13)

	// 无法取得常量i的地址
	// var ptr1, ptr2 *int = &i, &j
	var ptr1, ptr2 *int = &b, &j
	fmt.Println(ptr1, ptr1, ptr2, *ptr2)
	// 常量定义但未使用，编译不报错
	const x = 3

	// 变量定义但未使用，编译会报错
	// y := 3

	const Pi = 3.14159
	const Ln2 = 0.693147180559945309417232121458176568075500134360255254120680009
	const Log2E = 1 / Ln2 // this is a precise reciprocal
	const Billion = 1e9   // float constant
	const hardEight = (1 << 100) >> 97

	const beef, two, k = "eat", 2, "veg"
	// const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
	// 定义枚举值
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	// const (
	// 	Monday, Tuesday, Wednesday = 1, 2, 3
	// 	Thursday, Friday, Saturday = 4, 5, 6
	// )

	// 常量用作枚举
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	// itoa，const中的每一行都让itoa+1
	const (
		a1 = 3        // itoa=0
		b1 = iota + 1 // a1 让itoa=0, 此处的itoa=1, b1=1+1
		b2            // 继承了上一个表达式的值，b2 = itoa+1, b2=2+1
		c1 = iota     // itoa=3
	)
	fmt.Println(a1, b1, b2, c1)

	// 使用类型作为枚举常量
	type Color int

	const (
		RED    Color = iota // 0
		ORANGE              // 1
		YELLOW              // 2
		GREEN               // ..
		BLUE
		INDIGO
		VIOLET // 6
	)
	fmt.Println(RED, ORANGE, YELLOW, GREEN, BLUE, INDIGO, VIOLET)
}
