# sync.Pool

## sync.Pool 对象获取

* Processor:
  * 私有对象（一个对象）：协程安全
  * 共享池：协程不安全

* 尝试从私有对象获取
* 私有对象不存在，尝试从当前 Processor 的共享池获取
* 如果当前 Processor 共享池也是空的，那么就尝试去其他 Processor 的共享池获取
* 如果所有子池都是空的，最后就用用户指定的 New 函数产生一个新的对象返回

```go
// 使用 sync.Pool
pool := &sync.Pool{
  New: func() interface{} {
    return 0
  },
}

arry := pool.Get().(int)
pool.Put(10)
```

## sync.Pool 对象放回

* 如果私有对象不存在则保存为私有对象
* 如果私有对象存在，放入当前 Processor 子池的共享池中

## sync.Pool 对象的生命周期

* GC 会清除 sync.Pool 缓存的对象
* 对象的缓存有效期为下一次 GC 之前

## sync.Pool 总结

* 适合于通过复用，降低复杂对象的创建和GC代价
* 协程安全，会有锁的开销
* 生命周期受GC影响，不适合于做连接池等，需自己管理生命周期的资源的池化
