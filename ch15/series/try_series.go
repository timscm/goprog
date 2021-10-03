package series

import "fmt"

func init() {
	fmt.Println("Init1 in series")
}

func init() {
	fmt.Println("Init2 in series")
}

func GetFibonacSeries(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

// 外面的包引用，需要首字母大写
func Square(i int) int {
	return i * i
}
