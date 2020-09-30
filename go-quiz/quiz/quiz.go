package quiz

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Problem struc
type Problem struct {
	q string
	a string
}

// StartQuiz starts the quiz with problems provided
func StartQuiz(problems []Problem, timeLimit *int) int {
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			return correct
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}
	return correct
}

// ParseLines parses a slice of strings into a slice of problems struct.
func ParseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			q: line[0],
			// Help to avoid problem with spaces in csv answers column
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// ExitQuiz exits the program with error message
func ExitQuiz(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
