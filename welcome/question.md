# 提问

## 最简的 golang 程序结构是什么？

`package` 关键字定义包名；`func` 定义函数，`main` 函数必须存在。

```go
package main

func main() {
}
```

## 如何运行 golang 程序？

1. 自动编译运行，运行后自动删除编译文件

    ```bash
    go run hello.go
    ```

2. 先编译，再运行

    ```bash
    # 编译
    go build hello.go
    # 运行
    ./hello
    ```

## 引入多个包的方式写法有哪些？

1. 一次引用 import 一个包

    ```go
    import "fmt"
    import "time"
    ```

2. 一次引用多个包，使用括号包含被引用的包名列表

    ```go
    import (
        "fmt"
        "time"
    )
    ```
