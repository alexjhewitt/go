package main

import "fmt"

func main() {
	type person struct {
		firstName string
		lastName  string
	}
	p := person{firstName: "Alex", lastName: "Hewitt"}
	fmt.Println(p)
	p.firstName = "Blah"
	fmt.Printf("%+v", p)
}
