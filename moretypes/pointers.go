package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // 指针指向 i 变量的内存地址
	fmt.Println(*p) // 通过指针读取变量的值
	*p = 21         // 通过指针修改变量的值
	fmt.Println(*p)
	fmt.Println(i)

	p = &j //将指针变量存储内存地址指向 j
	*p = *p / 37
	fmt.Println(j)
}
