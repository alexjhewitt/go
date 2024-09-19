package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	p := person{
		firstName: "Alex",
		lastName:  "Hewitt",
		contact: contactInfo{
			email:   "test@mail.com",
			zipCode: 55555,
		},
	}
	fmt.Println(p)
}
