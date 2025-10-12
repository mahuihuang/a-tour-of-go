# 提问

## 写一个函数的泛型？

`comparable` 是泛型的约束条件之一，即便使用泛型也尽可能缩小参数的可变范围

```go
func f[T comparable](t T) T {
    return t
}
```

## 写一个结构体的泛型？

```go
type S[T comparable]struct {
   value T
}
```

## 泛型对比 `any`/`interface{}` 的优势？

1. 减少代码的重复编写，例如将 map 的所有值相加，map 的类型是 `map[string]int` 或 `map[int]int`。

2. 缩小不确定范围，能用泛型就不要用
