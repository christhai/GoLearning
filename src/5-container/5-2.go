package main

import "fmt"

func main() {
	a := [4]float64{67.7, 89.8, 21, 78}
	b := [...]int{2, 3, 5}
	fmt.Printf("len(array A) %d, len(array B) %d \n", len(a), len(b))
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], "\t")
	}
	fmt.Println()

	for _, value := range b {
		fmt.Print(value, "\t")
	}
}
