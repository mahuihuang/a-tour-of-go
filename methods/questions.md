# 提问

## 函数与方法的差别？

方法比函数多了一个接收器

## 指针接收器的使用场景？

方法内修改接收器的值

## 非指针字面量能否调用指针字面量接收器的方法？

可以，调用时会自动取字面量的地址

## 如何把一个指针接收器的方法修改为函数？

方法

```go
func (p *T) f() {}
```

函数

```go
func f(p *T) {}
```

## 什么是 interface？如何定义

interface 是方法的签名集合，绑定方法必须使用结构体

```go
type i interface {
    f() float64
}
```

## interface 的值是什么？

可理解为接口的值是接受对象的值和对象类型，可使用 `fmt.Printf("v = %v t =%T\n")`

## 如何判断接口的值是否为空？

接口

## 空接口 `interface {}` 有什么用法？

接受任意类型的值，在不确定类型时可用。

## 什么是断言？如何通过断言获取字面量的值？

获取变量的底层值。`i.(string)`

## 如何通过断言获取变量的类型？配合 switch 验证？

```go
switch i.(type) {
    case string:
    case int:
    default:
}
```

## 如何自定义错误输出格式？理解 `stronv.Atoi` 错误处理过程。

`error` 是一个 interface,实现了 `Error` 方法即可赋值给 `error`

## 温习 io.Reader
