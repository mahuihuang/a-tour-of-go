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

channel 和 mutex 都可以做为互斥锁。为了更好的理解与使用，需要知道他们的明确用途。

- channel 用于 goroutine 之前的数据传输
- mutex 保证同一时刻数据只能被单个 goroutine 操作

## 为什么多个 goroutine 写入 map 需要互斥锁，而访问 int32 不需要？

map 是一种复合类型的数据结构，有多个指针。int32 的底层指向了一个连续地址。

## 多 goroutine 不使用互斥锁 访问 int64 会出现什么问题？

发生数据竞态，最终的结果与期望的结果不一致，如果是一个自增的函数，可能会导致结果偏小。
