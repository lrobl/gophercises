package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

type quiz struct {
	filename      string
	timeLimit     int
	questionCount int
	correctCount  int
}

func main() {
	// parse args
	filePtr := flag.String("file", "problems.csv", "Path to the quiz csv file.")
	limitPtr := flag.Int("timeLimit", 30, "Quiz time limit in second.")

	flag.Parse()

	// create the quiz
	q := quiz{
		filename:  *filePtr,
		timeLimit: *limitPtr,
	}

	// proctor the quiz
	proctor(q)

}

func proctor(q quiz) {
	// Open the file and handle any error
	file, err := os.Open(q.filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	// Create a reader opject from the open file
	r := csv.NewReader(file)

	// loop through each row in the csv
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		isCorrect := ask(record[0], record[1])
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
