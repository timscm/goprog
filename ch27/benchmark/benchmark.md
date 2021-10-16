# benchmark

```go
// BenchmarkXXX
// testing.B
func BenchmarkConcatStringByAdd(b *testing.B) {
    // 与性能测试无关的代码
    b.ResetTimer()
    for i :=0; i<b.N; i++{
        // 测试代码
    }
    b.StopTimer()
    // 与性能测试无关的代码
}

// go test -bench=. -benchmem
// -bench=<相关benchmark测试>
// windows下使用go test命令行时，-bench=.应写为-bench="."
```
