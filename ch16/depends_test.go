// Go 未解决的依赖问题
// 1. 同一环境下，不同项目使用同一包的不同版本
// 2. 无法管理对包的特定版本的依赖
// vendor 路径
// 随着 Go 1.5 release 版本的发布，vendor 目录被添加到 GOPATH 和
// GOROOT 之外的依赖目录查找的解决方案。在 Go 1.6 之前，你需要手动
// 的设在环境变量
// 查找依赖包路径的解决方案如下：
// 1. 当前包下的 vendor 目录
// 2. 向上级目录查找，直到找到 src 下的 vendor 目录
// 3. 在 GOPATH 下面查找依赖包
// 4. 在 GOROOT 目录下查找
// godep， glide， dep 依赖包管理工具
//
// Go 1.11 使用 module 来实现模块
package depends_test

import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
