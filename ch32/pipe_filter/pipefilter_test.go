/*
Pipe-Filter 模式
1. 非常适合与数据处理及数据分析系统
2. Filter封装数据处理的能力
3. 松耦合：Filter只跟数据（格式）耦合
4. Pipe用于连接Filter传递传递数据或者在异步处理过程中缓冲数据流，
   进程内不同调用时，pipe演变成为数据在方法调用间传递

Filter和组合模式
              __________________________________
              |                                 |
Pump ----pipe---> Filter ----pipe----> Filter ----pipe---->Sink
              |_________________________________|

示例：
"1,2,3" ->SplitFilter -> ["1", "2", "3"] ->ToIntFilter -> [1,2,3] ->SumFilter -> 6

*/
package pipefilter

import "testing"

func TestStraightPipline(t *testing.T) {
	spliterFilter := NewSplitFilter(",")
	toIntFilter := NewToIntFilter()
	sumFilter := NewSumFilter()

	sp := NewStraightPipeline("p1", spliterFilter, toIntFilter, sumFilter)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
	if ret != 6 {
		t.Fatalf("the expected is 6, but the actual is %d", ret)
	}
}
