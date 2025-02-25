package main

import (
	"fmt"
)

func PrintStar(x int) {
	for i := 0; i < x; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := x - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	x := 5
	PrintStar(x)
}
