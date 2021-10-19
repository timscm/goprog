package main

/*

二、通过 HTTP 方式输出 Profile

* 简单，适合于持续性运行的应用
* 在应用程序中导入 import _ "net/http/pprof", 并启动 http server 即可
* http://<host>:<port>/debug/pprof
* go tool pprof http://<host>:<port>/debug/pprof/profile?seconds=10 （默认值30秒）
* go-torch -seconds 10 http://<host>:<port>/debug/pprof/profile

*/

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func GetFibonacciSerie(n int) []int {
	ret := make([]int, 2, n)
	ret[0] = 1
	ret[1] = 1
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func createFBS(w http.ResponseWriter, r *http.Request) {
	var fbs []int
	for i := 0; i < 1000000; i++ {
		fbs = GetFibonacciSerie(50)
	}
	w.Write([]byte(fmt.Sprintf("%v", fbs)))
}

/*
http://127.0.0.1:8081/
http://127.0.0.1:8081/fb
http://127.0.0.1:8081/debug/pprof/

go tool pprof http://127.0.0.1:8081/debug/pprof/profile
(pprof) top
(pprof) top -cum

go-torch http://127.0.0.1:8081/debug/pprof/profile
#INFO[23:12:07] Run pprof command: go tool pprof -raw -seconds 30 http://127.0.0.1:8081/debug/pprof/profile
#INFO[23:12:37] Writing svg to torch.svg
# 浏览器打开 torch.svg 文件

*/
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/fb", createFBS)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
