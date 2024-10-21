package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Pt. 1 Put 100 random numbers between 0 and 100 into an int slice
	var arr = make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(101))
	}
	fmt.Println(arr)

	// Pt. 2 Loop over slice in Pt. 1 and do something based on rules in if statement
	for _, n := range arr {
		switch {
		case n%2 == 0 && n%3 == 0:
			fmt.Println("Six!")
		case n%2 == 0:
			fmt.Println("Two!")
		case n%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
