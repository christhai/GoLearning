package main

import "fmt"

//func main() {
//	fmt.Printf("hello, world\n")
//}
//

func main() {
	b := 3158
	a := &b
	fmt.Println("", a)
	fmt.Println("", *a)
	*a++
	fmt.Println("", b)
}

