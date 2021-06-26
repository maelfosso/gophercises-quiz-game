package main

import (
	"encoding/csv"
	"flag"
	"fmt"
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

var (
	file     = flag.String("csv", "./problems.csv", "Path to the file containing the quizzes")
	duration = flag.Int("duration", 30, "Max time to answer")
)

func loadQuizzesFile() ([][]string, error) {
	file, err := os.Open(*file)
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

type RInputTimer struct {
	timeout    bool
	input      string
	inputError bool
}

func readInput(r chan RInputTimer) {
	var input string
	_, err := fmt.Scanf("%s", &input)

	if err != nil {
		r <- RInputTimer{
			timeout:    false,
			inputError: true,
		}
	} else {
		r <- RInputTimer{
			timeout:    false,
			input:      input,
			inputError: false,
		}
	}

}

func main() {
	flag.Parse()

	var durationIsOk string

	// Load the CSV file
	dataFile, err := loadQuizzesFile()
	if err != nil {
		exit(fmt.Sprintf("error occured when loading quizzes file, %v", err))
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

	fmt.Println("The Maximum time for the quiz is ", *duration, " seconds")
	fmt.Println("Press enter to continue or Ctrl+C to stop the program and restart it with another duration")
	fmt.Scanf("%s", &durationIsOk)
	t := time.After(time.Duration(*duration) * time.Second)

problemloop:
	for _, quiz := range quizzes {

		fmt.Printf("%s > ", quiz.question)

		r := make(chan RInputTimer)
		go readInput(r)

		select {
		case v := <-r:
			if v.inputError {
				turn.totalWrongAnswer += 1
			} else {

				if v.input == "Q" {
					break problemloop
				}

				turn.totalQuestions += 1

				if v.input == quiz.answer {
					turn.totalCorrectAnswer += 1
				} else {
					turn.totalWrongAnswer += 1
				}
			}

		case <-t:
			break problemloop
		}
	}

	fmt.Println("\nResults:")
	fmt.Println("\tTotal Number of Questions: \t", turn.totalQuestions)
	fmt.Println("\tTotal Number of Corrent Answers: \t", turn.totalCorrectAnswer)
	fmt.Println("\tTotal Number of Wrong Answers: \t", turn.totalWrongAnswer)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
