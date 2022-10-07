package main

import "fmt"

//func main() {
//	fmt.Printf("hello, world\n")
//}
// 使用指针作为函数的参数

func main() {
	countryMap := make(map[string]string)
	countryMap["China"] = "Beijing"
	countryMap["zzz"] = "ddd"
	countryMap["xccc"] = "eeee"
	countryMap["vvvv"] = "ddsa"
	value, ok := countryMap["China"]
	fmt.Printf("%q \n", value)
	fmt.Printf("%T, %v \n", ok, ok)
	if ok {
		fmt.Println("Captical:", value)
	} else {
		fmt.Println("not Fund")
	}
}
