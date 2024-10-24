package main

import (
	"fmt"
	"os"
)

func fileLen(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}
	defer f.Close()
	data, _ := os.ReadFile(fileName)
	fmt.Println(len(data))
	return len(data), nil
}

// func main() {
// 	fileLen("divByZero.go")
// }
