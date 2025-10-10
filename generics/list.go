package main

import "fmt"

type I interface {
	Error() string
}
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	head := &List[int]{}
	err := describe(head)
	if err != nil {
		fmt.Println(err)
	}
	for i, node := 0, head; i < 10; i++ {
		newNode := &List[int]{val: i}
		node.next = newNode
		node = newNode
	}
	err = describe(head)
	if err != nil {
		fmt.Println()
	}
}

func describe[T any](l *List[T]) error {
	if l == nil || l.next == nil {
		return l
	}
	for node := l.next; node != nil; node = node.next {
		fmt.Printf("%v\n", node.val)
	}
	return nil
}

func (l *List[T]) Error() string {
	return fmt.Sprintf("Error: %s", "the head or next node is empty")
}
