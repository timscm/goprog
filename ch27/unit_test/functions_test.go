// go test -v -cover
//
// 提供测试用的断言
// go get -u github.com/stretchr/testify/assert
package testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expect := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expect[i] {
			t.Errorf("input is %d, the expected is %d, the actual %d",
				inputs[i], expect[i], ret)
		}
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End") // 会输出该行
}

func TestFatalInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Fatal")
	fmt.Println("End") // 不会输出该行
}

func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expect := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expect[i], ret)
	}
}
