package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	go sendData(ch1)
	//Method 1. receive data
	//for {
	//	data := <-ch1
	//	if data == "" {
	//		break
	//	}
	//	fmt.Println("method 1:Receive data from chan", data)
	//}
	//Method 2. receive data
	//for {
	//	data, ok := <-ch1
	//	fmt.Println(ok)
	//	if !ok {
	//		break
	//	}
	//	fmt.Println("method 2: Receive data from chan ", data)
	//}

	//Method 3. for ...range receive data
	for value := range ch1 {
		fmt.Println("method 3: Receive data from chan \n", value)
	}
}

func sendData(ch1 chan string) {
	defer close(ch1)
	for i := 0; i < 3; i++ {
		ch1 <- fmt.Sprint("sending data", i)
	}
	fmt.Println("Finish ... sending data")
}
