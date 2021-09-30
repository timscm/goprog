package string_test

import (
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)

	s = "hello"
	// s[1] = '2'  编译错误 strings are immutable
	t.Log(len(s))

	// 重新赋值，是创建新的字符串, “严”
	s = "\xE4\xB8\xA5"
	t.Log(s)

	s = "中"
	t.Log(len(s)) // 输出byte数: 3

	// Unicode: code point 字符集
	// UTF8: 字符集的实现
	// 字符		“中”
	// Unicode	0x4E2D
	// UTF-8    0xE4B8AD
	// string/[]byte [0xE4, 0xB8, 0xAD]

	// %x 输出16进制, Sizeof(c[0])=4, len(c)=1
	c := []rune(s)
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x, len:%d", c[0], len(c))
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		// %c, %x, %[1]=c
		t.Logf("%[1]c %[1]x", c)
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	// 编译报错，返回两个值
	// t.Log(10 + strconv.Atoi(s))
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
