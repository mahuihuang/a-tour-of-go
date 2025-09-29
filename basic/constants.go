package main

import "fmt"

const Pi = 3.14

func main() {

	// 可自动推导类型，但不能使用 :=
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
