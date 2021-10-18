package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// 存在一个默认匹配的路由规则：
// 注册：/time, 则只能匹配到 /time, /time/ 匹配到了 /
// 注册：/time/, 则能匹配 /time, /time/, /time/123
func main() {
	// http://127.0.0.1:8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// 注册： http://127.0.0.1:8080/time  --> 固定匹配到 /time
	// 注册： http://127.0.0.1:8080/time/ --> 可以匹配到 /time/, /time/1
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		w.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8080", nil)
}
