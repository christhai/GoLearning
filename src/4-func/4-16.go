package main

import "fmt"

//func main() {
//	fmt.Printf("hello, world\n")
//}
// 使用指针作为函数的参数

func main() {
	a := 31

	fmt.Println("before", a)
	fmt.Printf("%T \n", a)
	fmt.Printf("%x \n", &a)
	b := &a
	change(b)
	fmt.Printf("%x \n", *b)
	fmt.Println("after", *b)
}

func change(val *int) {
	*val = 15
}
