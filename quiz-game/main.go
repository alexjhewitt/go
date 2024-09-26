package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Keep track of right/wrong answers
	// correctAnswers := 0
	// wrongAnswers := 0

	// Flag creation
	fileNamePtr := flag.String("f", "problems.csv", "Name of file with quiz questions")
	flag.Parse()

	// Read in CSV and work with data
	file, err := os.Open(*fileNamePtr)
	if err != nil {
		fmt.Println("Error:", err)
	}
	r := csv.NewReader(file)
	data, err := r.ReadAll()
	for _, line := range data {
		q := line[0]
		// a := line[1]
		fmt.Printf("What is %s?\t", q)
	}

	// Ask next question

	// Customize the csv file name with flag, but default to problems.csv

	//
}
