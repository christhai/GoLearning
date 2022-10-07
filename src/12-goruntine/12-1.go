package main

import (
	"fmt"
)

func main() {
	go hello()
	fmt.Println("this is main function")
}

func hello() {
	fmt.Println("hello goroutine")
}
