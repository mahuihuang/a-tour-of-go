package main

import (
	"fmt"
	"time"
)

type MyError struct {
	when time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it did't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
