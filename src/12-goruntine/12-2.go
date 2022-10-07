package main

import (
	"fmt"
	"time"
)

func hellog() {
	fmt.Println("hello goroutine")
}

func main() {
	go hellog()
	time.Sleep(50 * time.Millisecond)
	fmt.Println("this is main function")
}
