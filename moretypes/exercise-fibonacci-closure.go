package main

import (
	"fmt"
)

func fibonacci() func() int {
	preSum, sum := 0, 1

	return func() int {
		result := preSum
		preSum, sum = sum, result+sum
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
