package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Quiz struct {
	question string
	answer   string
}

type Quizzes = []Quiz

func loadQuizzesFile() ([][]string, error) {
	file, err := os.Open("./problems.csv")
	if err != nil {
		return nil, fmt.Errorf("error encountered when opening quizzes file, %v", err)
	}

	csvReader := csv.NewReader(file)

	return csvReader.ReadAll()
}

func parseDataQuizzes(dataQuizzes [][]string) Quizzes {
	var quizzes Quizzes

	for _, data := range dataQuizzes {
		quiz := Quiz{
			question: data[0],
			answer:   data[1],
		}

		quizzes = append(quizzes, quiz)
	}

	return quizzes
}

func main() {
	// Load the CSV file
	dataFile, err := loadQuizzesFile()
	if err != nil {
		log.Fatalf("error occured when loading quizzes file, %v", err)
	}

	// Parse its content
	parseDataQuizzes(dataFile)

	// Map into struct the read content

	// Loop through all the quizzes for the user to solve them
}
