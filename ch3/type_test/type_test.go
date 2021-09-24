package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var (
		a int = 1
		b int64
		c MyInt
	)
	b = int64(a)
	t.Log(a, b)

	c = MyInt(a)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	// if s == nil {
	// }
}
