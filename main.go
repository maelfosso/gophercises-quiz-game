package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Quiz struct {
	id       int
	question string
	answer   string
}

type Quizzes = []Quiz

type Turn struct {
	totalQuestions     int
	totalCorrectAnswer int
	totalWrongAnswer   int
}

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

	for index, data := range dataQuizzes {
		quiz := Quiz{
			id:       index,
			question: data[0],
			answer:   data[1],
		}

		quizzes = append(quizzes, quiz)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(quizzes), func(i, j int) {
		quizzes[i], quizzes[j] = quizzes[j], quizzes[i]
	})

	return quizzes
}

func main() {
	// Load the CSV file
	dataFile, err := loadQuizzesFile()
	if err != nil {
		log.Fatalf("error occured when loading quizzes file, %v", err)
	}

	// Parse its content
	quizzes := parseDataQuizzes(dataFile)

	fmt.Println("Quizzes loaded ...")
	fmt.Println("Please answer to these questions. Press Q to exit.")

	// Loop through all the quizzes for the user to solve them
	turn := Turn{
		totalQuestions:     0,
		totalCorrectAnswer: 0,
		totalWrongAnswer:   0,
	}

	for _, quiz := range quizzes {
		var answer string

		fmt.Printf("%s > ", quiz.question)

		_, err := fmt.Scanf("%s", &answer)
		if err != nil {
			turn.totalWrongAnswer += 1
		} else {

			if answer == "Q" {
				break
			}

			turn.totalQuestions += 1

			if answer == quiz.answer {
				turn.totalCorrectAnswer += 1
			} else {
				turn.totalWrongAnswer += 1
			}
		}

	}

	fmt.Println("\nResults:")
	fmt.Println("\tTotal Number of Questions: \t", turn.totalQuestions)
	fmt.Println("\tTotal Number of Corrent Answers: \t", turn.totalCorrectAnswer)
	fmt.Println("\tTotal Number of Wrong Answers: \t", turn.totalWrongAnswer)
}
