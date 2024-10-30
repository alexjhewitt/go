package makePerson

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p1 := MakePerson("Alex", "Hewitt", 777)
	p2 := MakePersonPointer("Dan", "Smith", 35)

	fmt.Println(p1)
	fmt.Println(p2)
}

func MakePerson(firstName, lastName string, age int) Person {
	p := Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
	return p
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	p := &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
	return p

}
