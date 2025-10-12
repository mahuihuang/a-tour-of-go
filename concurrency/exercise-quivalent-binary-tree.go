package main

import (
	"fmt"
	"sort"
	"time"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	len := 10
	buffer := make(chan int, len)
	traverseTree(t, buffer)
	tree := readChannel(buffer)
	for _, v := range tree {
		ch <- v
	}
}

// 读取 channel 数据并排序
func readChannel(c chan int) []int {
	len := 10
	s := make([]int, len)
	for i := 0; i < len; i++ {
		s[i] = <-c
	}
	sort.Ints(s)
	return s
}

// traverseTree 遍历二叉树
func traverseTree(t *tree.Tree, c chan int) {
	if t == nil {
		return
	}
	c <- t.Value
	traverseTree(t.Left, c)
	traverseTree(t.Right, c)
}
func Same(t1, t2 *tree.Tree) bool {
	t1Chan := make(chan int)
	t2Chan := make(chan int)

	go Walk(t1, t1Chan)
	go Walk(t2, t2Chan)
	for i := 0; i < 10; i++ {
		if <-t1Chan != <-t2Chan {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)

	go Walk(tree.New(1), ch)
	time.Sleep(time.Second)
	fmt.Println("tree.New(1) is tree.New(1) quivalent?", Same(tree.New(1), tree.New(1)))
	fmt.Println("tree.New(1) is tree.New(2) quivalent?", Same(tree.New(1), tree.New(2)))
}
