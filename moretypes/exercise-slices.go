package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		// 不使用嵌套循环初始化切片，试试会发生什么错误？
		s[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			s[y][x] = uint8(x * y)
		}
	}
	return s
}

func main() {
	pic.Show(Pic)
}
