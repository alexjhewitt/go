package main

import "fmt"

func prefixer(s string) func(string) string {
	return func(i string) string {
		return fmt.Sprintf("%s %s", s, i)
	}

}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
