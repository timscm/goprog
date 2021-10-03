package client

import (
	"goprog/ch15/series" // GOPATH: src/之后的路径
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacSeries(5))
	t.Log(series.Square(5))
}
