package main

import (
	"fmt"
	"time"
)

func main() {

	// 1. non-cached
	var ch1 chan int
	ch1 = make(chan int)
	fmt.Println("non-cached channel", len(ch1), cap(ch1))
	go func() {
		data := <-ch1
		fmt.Println("get data from ch1", data)
	}()
	ch1 <- 100
	time.Sleep(time.Second)
	fmt.Println(" ok ", "main over ...")
	// 2. non-cached
	ch2 := make(chan string)
	go sendData(ch2)
	for data := range ch2 {
		fmt.Println("\treading data from ch2", data)
	}
	fmt.Println("main over")

	//3.cached channel
	ch3 := make(chan string, 6)
	go sendData(ch3)
	for data := range ch3 {
		fmt.Println("\treading data", data)
	}

	fmt.Println("main over")
}

func sendData(ch chan string) {
	for i := 1; i <= 3; i++ {
		ch <- fmt.Sprintf("data%d", i)
		fmt.Println("put data to the channel", i)
	}
	defer close(ch)
}
