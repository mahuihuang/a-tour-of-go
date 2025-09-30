# 提问

## 如何写一个 for 语句，初始化语句的变量的作用于范围？

for 关键字定义循环语句:

- 初始化语句：可定义变量，定义的变量仅在循环内有效。
- 条件语句：满足条件语句循环继续，否则中断。
- 后置语句：每一次循环完成之后执行的语句。

```go
for <init statement>;<condition statement>;<post statement> {
}
```

## 如何写没有初始化语句和后置语句的 for？

for 可以没有初始化，后置语句

```go
for ;<condition statement>; {
}
```

## 如何编写一个类似于 C 语言中的 While 语句？

```go
for <condition statement> {
}
```

## 如何写一个永久循环？

```go
for {
}
```

## if 语句都有哪些写法？

1. 无初始化语句

    ```go
    if <condition> {
    }
    ```

2. 有初始化语句

    ```go
    if <init statement>;<condition> {

    }
    ```

## 如何写 if else 语句？

```go
if <condition> {
} else {
}
```

## switch 语句的结构？

## switch 语句的匹配顺序？

## swich 与 if-else  的差别？

## 如何使用 defer？ 何时触发 defer？

## defer 都有哪些适用场景？

## 当一个函数内调用了多次 defer,应该按照什么顺序执行？

