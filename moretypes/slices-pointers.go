package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	// 切片不存储数据，只是在描述一个数组
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	fmt.Printf("names[0]=%v,names[1]=%v,names[2]=%v\n", &names[0], &names[1], &names[2])
	fmt.Printf("a[0]=%v,a[1]=%v,b[0]=%v,b[1]=%v\n", &a[0], &a[1], &b[0], &b[1])
}
