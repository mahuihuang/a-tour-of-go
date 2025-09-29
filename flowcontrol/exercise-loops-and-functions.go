package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("第 %d 迭代 z = %v\n", i, z)
		//1e-15=0.000000000000001
		if z*z-x <= 1e-15 {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(4))
}
