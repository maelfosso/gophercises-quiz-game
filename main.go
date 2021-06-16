package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func loadQuizzesFile() ([][]string, error) {
	file, err := os.Open("./problems.csv")
	if err != nil {
		return nil, fmt.Errorf("error encountered when opening quizzes file, %v", err)
	}

	csvReader := csv.NewReader(file)

	return csvReader.ReadAll()
}

func main() {
	// Load the CSV file
	dataFile, err := loadQuizzesFile()
	if err != nil {
		log.Fatalf("error occured when loading quizzes file, %v", err)
	}

	// Parse its content

	// Map into struct the read content

	// Loop through all the quizzes for the user to solve them
}
