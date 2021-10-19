package main

import (
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

/*
sudo apt install graphviz
go get github.com/uber/go-torch
wget https://raw.githubusercontent.com/brendangregg/FlameGraph/master/flamegraph.pl
chmod +x flamegraph.pl
mv flamegraph.pl golabs/bin/

一、通过文件方式输出Profile

* 灵活性高，适用于特定代码段的分析
* 通过手动调用 runtime/pprof 的API
* API 相关文档 https://studygolang.com/static/pkgdoc/pkg/runtime_pprof.htm
* go tool pprof [binary] [binary.prof]

*/

const (
	row int = 1000
	col int = 1000
)

func fillMatrix(m *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(1000)
		}
	}
}

func calculate(m *[row][col]int) {
	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += m[i][j]
		}
	}
}

/*
Go 支持的多种Profile
go help testflag
https://golang.org/src/runtime/pprof/pprof.go

go build prof.go
./prof
ls *.conf
#-rw-rw-r-- 1 tim tim     615 Oct 19 22:46 cpu.prof
#-rw-rw-r-- 1 tim tim     570 Oct 19 22:46 goroutine.prof
#-rw-rw-r-- 1 tim tim     814 Oct 19 22:46 mem.prof
go tool pprof prof cpu.prof
(pprof) top
(pprof) list fillMatrix
(pprof) list fillMatrix
(pprof) exit
ls profile001.svg
# 浏览器打开该svg文件

go-torch cpu.prof
#INFO[22:53:11] Run pprof command: go tool pprof -raw -seconds 30 cpu.prof
#INFO[22:53:11] Writing svg to torch.svg
# 浏览器打开该 torch.svg 文件进行查看

*/
func main() {
	// 创建输出文件
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}

	// 获取系统信息
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	// 主逻辑区，进行一些简单的代码运算
	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	f1, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}

	// runtime.GC() // GC, 获取最新的数据信息
	if err := pprof.WriteHeapProfile(f1); err != nil { // 写入内存信息
		log.Fatal("could not write memory profile: ", err)
	}
	f1.Close()

	f2, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("colud not create goroutine profile: ", err)
	}
	if gProf := pprof.Lookup("goroutine"); gProf == nil {
		log.Fatal("could not write goroutine profile:", err)
	} else {
		gProf.WriteTo(f2, 0)
	}
	f2.Close()
}
