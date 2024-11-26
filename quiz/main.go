package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type quiz struct {
	filename      string
	timeLimit     int
	problems      []problem
	questionCount int
	correctCount  int
}

type problem struct {
	question string
	answer   string
}

func main() {
	// parse args
	fileName := flag.String("fileName", "problems.csv", "Path to the quiz csv file in format question, answer")
	timeLimit := flag.Int("timeLimit", 30, "Quiz time limit in second.")

	flag.Parse()

	// parse the csv into a quiz object
	q := parseQuizCSV(*fileName, *timeLimit)
	fmt.Println(q)

	// proctor the quiz
	proctor(q)

}

func parseQuizCSV(fileName string, timeLimit int) quiz {
	// Open the file and handle any error
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	// Create a reader opject from the open file
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(records)

	quizSize := len(records)

	// save the questions and answers into a 2d slice
	problems := make([]problem, quizSize)
	for i, record := range records {
		p := problem{
			question: record[0],
			answer:   strings.TrimSpace(record[1]),
		}
		problems[i] = p
	}

	q := quiz{
		filename:      fileName,
		timeLimit:     timeLimit,
		problems:      problems,
		questionCount: quizSize,
		correctCount:  0,
	}

	return q

}

func proctor(q quiz) {
	// loop through each problem in the quiz
	for _, problem := range q.problems {

		isCorrect := ask(problem.question, problem.answer)
		q.questionCount++
		if isCorrect {
			q.correctCount++
		}
		fmt.Print(q.correctCount, " out of ", q.questionCount, " correct.\n\n")
	}
}

func ask(question string, answer string) bool {
	// Ask the question
	fmt.Print(question)

	// define a string variable to hold the response
	var input string

	// get input from the user at terminal, passing an
	// address to variable we want the input stored
	// Scanf will allow us to pull the 1 word answer
	fmt.Scanf("%s\n", &input)

	// Check if the input matches the answer
	if input == answer {
		fmt.Println("Correct!")
		return true
	} else {
		fmt.Println("Incorrect!")
		return false
	}

}
