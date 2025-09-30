# 问题

## 如何定义包名？

|参考： [Package names](https://go.dev/blog/package-names)

`pakcage <package_name>` 的格式定义包名。包名应当使用小写字母；合理的缩写，不产生歧义。例如 `fmt` 为 formatted 的缩写。

## fmt.Println() 第一个参数的末尾是否需要保留空格？

不需要，Println 拼接参数时会自动添加空格。Println 是 Print Line 的缩写。

## 如何格式化输出变量？

调用 fmt 包内的函数，不同的函数可满足不同的格式化需求。

## 如何判断包内的定义的函数，类型等能否被引用？

被导出的函数、常量、变量、类型以大写字母开头

## 如何定义函数？主函数和普通函数

`func` 关键字定义函数，函数名后面紧跟参数，括号内先定义参数的变量名，再是参数的类型；最后是返回值的类型。主函数的名称为 `main` 没有入参和返回值。

普通函数

```go
func functionname(arg type)(type)
{ 
    return var
}
```

主函数

```go
func main()
{
}
```

## 当函数的相邻参数是相同类型，应该如定义函数参数？

仅保留最后一个参数的类型。例如 `func(x string, y string)` 简写后 `func(x string, y string)`

## 返回单个值与多个返回值的写法差异？可参考引包写法

返回单个值，可不用使用括号包裹返回类型。

```go
func add(x, y int) int {
    return x + y
}
```

返回多个返回值，使用括号包裹返回值类型，使用逗号分割

```go
func swap(x, y string) (string, string) {
    return y, x
}
```

## 函数返回参数定义了变量名称有何作用，例如 `func sum() (x, y int)` ？不推荐使用

返回值参数定义了变量名时，在函数内部可省略定义对应变量名的步骤。

## 如何定义变量？基础变量的默认值是什么？

var 关键字定义变量，表达式为 `var variable type`，例如 `var c string`。字符串，布尔值，整型的默认值依次为 "",false,0。

## 根据根据函数作用域范围可分为哪两类变量？

函数外全局变量，函数内为局部变量。

## 如何使用一个 var 关键词定义多个同类型变量？

句式为 `var variable1, variable1 type`，例如 `var c, python string`。

## 如何定义有初始值的变量？

句式为 `var variable type = value` 例如定义单个变量 `var x int = 1`

## `:=` 简短变量表达式有什么作用？使用的时候需要注意什么？

1. 用于声明一个有初始化值的变量，并自动进行类型推导。例如 `sum := 0`。

2. `:=` 仅能在函数内使用

## 基础变量类型 byte，rune 实际的类型是什么？ 都有哪些使用场景

byte 是 uint8 别称，表示一个字节，可用于存储一个 ASCII 字符。rune 是 int32 的别称，可用于存储一个 unicode 编码。byte 与 rune 都是语义化的称呼，增加了可读性。

## 如何打印变量的类型？

使用 fmt 的 Printf 方法，指定格式化类型参数为 `%T`,例如

```go
fmt.Printf("%T\n", v)
```

## 如何打印变量值的默认类型值？

使用 fmt 的 Printf 方法，且指定格式化类型参数为 `%v`,例如

```go
fmt.Printf("%v\n", v)
```

## 如何转换类型？

使用 `T(v)` 表达式转换，T 表示转换后的类型，v 表示被转换的对象。

## 如何定义常量？能否使用 `:=` 语法？

使用 `const` 关键字定义变量，虽然不使用 `:=` ，但依然可以进行类型推导。例如：

```go
const Pi = 3.14
const Truth bool = true
```

## 当函数的参数为 float64 类型，能否传入 int 类型？

可以，高精度能够兼容低精度的变量。
