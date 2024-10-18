package main

import "fmt"

func main() {
	// Pt. 1
	greetings := []string{"Hello", "Hola", "こんにちは", "ສະບາຍດີ", "Вітаю"}
	g1 := greetings[:2]
	g2 := greetings[2:]
	g3 := greetings[3:]
	fmt.Println(g1)
	fmt.Println(g2)
	fmt.Println(g3)

	// Pt. 2
	message := "Hi 👩 and 👨"
	runes := []rune(message)
	fmt.Println(string(runes[3]))

	// Pt. 3
	type Employee struct {
		firstName string
		lastName string
		id int
	}

	e1 := Employee {
		"Dan",
		"Smith",
		1,
	}
	e2 := Employee{
		firstName: "Alex",
		lastName: "Hewitt",
		id: 2,
	}
	var e3 Employee
	e3.firstName = "Jabba"
	e3.lastName = "The Hutt"
	e3.id = 3

	fmt.Println(e1, e2, e3)

}
