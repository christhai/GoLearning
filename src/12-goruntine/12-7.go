package main

import (
	"fmt"
)

func main() {
	var ch1 chan int
	ch1 = make(chan int)
	fmt.Printf("%T\n", ch1)
	ch2 := make(chan bool)
	go func() {
		data, ok := <-ch1
		if ok {
			fmt.Println("Receive data from child chan ", data)
		}
		ch2 <- true
	}()
	ch1 <- 10
	//<-ch2 //block
	fmt.Println("main over")
}
