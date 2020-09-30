package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/juangm/go-exercises/go-quiz/quiz"
)

func main() {
	csvFilename := flag.String("csv", "problem.csv", "a csv file in teh format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		quiz.ExitQuiz(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		quiz.ExitQuiz("Failed to parse the provided CSV file")
	}
	problems := quiz.ParseLines(lines)
	correct := quiz.StartQuiz(problems, timeLimit)
	fmt.Printf("\nYou score %d out of %d.\n", correct, len(problems))
	os.Exit(0)
}
