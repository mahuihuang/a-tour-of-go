# 提问

## 一句描述什么是 goroutines？

goroutine 是 golang 管理的轻量级线程

## channels 有何用法？如何创建 channel

作用，不同 goroutine 的相互通信。

无缓存 channel

```go
c := make(chan int)
```

有缓存 channel

```go
c := make(chan int, 2)
```

## goroutine 结束后没有关闭有 buffer 的 channel 会出现什么问题？

当在发送 goroutine 结束前没有使用 close 关闭 channel，会导致结束 channel 读取出现死锁。使用 `for v := range c` 持续读取 channel 内的数据，直至 channel 被正常关闭。

## select 的使用场景？

`for` 与 `select` 的组合可以实现轮训多个 channel

## 重构等价二叉树练习题（Equivalent Binary Trees）

## 有了 channel 为什么还需要 mutex ?

## 为什么多个 goroutine 访问 map 需要互斥锁，而访问 int32 不需要？

## 多 goroutine 不使用互斥锁 访问 int64 会出现什么问题？
