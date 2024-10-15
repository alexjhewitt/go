package main

import "fmt"

func main() {
	const value = 5
	i := 20
	f := float64(i)
	fmt.Printf("%d %f\n", i, f)
	i = value
	f = value
	fmt.Printf("%d %f\n", i, f)

}
