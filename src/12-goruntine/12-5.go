package main

import (
	"fmt"
	"time"
)

func main() {
	go printNum()
	go printLetter()
	time.Sleep(250 * time.Second)
	fmt.Println("\n main over...")
}

func printNum() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Print("%d", i)
	}
}

func printLetter() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Print("%c", i)
	}
}
