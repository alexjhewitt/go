package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	printArea()
}

func (s square) printArea() {
	fmt.Println(s.sideLength * s.sideLength)
}

func (t triangle) printArea() {
	fmt.Println(.5 * t.base * t.height)
}
func main() {
	t := triangle{base: 5, height: 5}
	s := square{sideLength: 10}
	s.printArea()
	t.printArea()
}
