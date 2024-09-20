package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
	// contact contactInfo (variableName variableType) is an embedded struct
	// Could also leave it as just contactInfo with no named variable, then when initializing
	// You would construct it as contactInfo: contactInfo{blah}
}

func main() {
	p := person{
		firstName: "Alex",
		lastName:  "Hewitt",
		contactInfo: contactInfo{
			email:   "test@mail.com",
			zipCode: 55555,
		},
	}
	pPtr := &p
	pPtr.updateName("Jim")
	p.print()
}

func(p person) print() {
	fmt.Printf("%+v", p)
}

func (pPtr *person) updateName(newFirstName string) {
	(*pPtr).firstName = newFirstName
}