package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// t.Log(a == c) 编译错误，长度不同不能比较
	t.Log(a == d)
}

// 按位置0， &^
// 1 &^ 0 -- 1   右边为0，左边的位保留原样
// 1 &^ 1 -- 0   右边为1，左边的位置为0
// 0 &^ 1 -- 0   右边为1，左边的位置为0
// 0 &^ 0 -- 0   右边为0，左边的位保留原样
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
