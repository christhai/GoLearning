package main

import (
	"fmt"
)

func main() {
	var ch1 chan int
	ch1 = make(chan int)

	go func() {
		ch1 <- 100
		ch1 <- 200
		close(ch1)
		ch1 <- 10
	}()
	data, ok := <-ch1
	fmt.Println("main read data", data, ok)
	data, ok = <-ch1
	fmt.Println("main read data", data, ok)
	data, ok = <-ch1
	fmt.Println("main read data", data, ok)
	data, ok = <-ch1
	fmt.Println("main read data", data, ok)
	data, ok = <-ch1
	fmt.Println("main read data", data, ok)
}
