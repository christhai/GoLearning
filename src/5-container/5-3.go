package main

import "fmt"

func main() {
	a := [4][2]int{{3, 4}, {5, 4}, {5, 3}, {6, 7}}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Print(a[i][j], "\t")
		}
		fmt.Println()
	}
}
