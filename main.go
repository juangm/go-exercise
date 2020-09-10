package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Structure for problems
type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problem.csv", "a csv file in teh format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			fmt.Printf("Correct!\n")
			correct++
		}
	}
	fmt.Printf("You score %d out of %d.\n", correct, len(problems))
	os.Exit(0)
}

// Function to parse lines from csv to problems
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			// Help to avoid problem with spaces in csv answers column
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// Helper function to exit the program with error message
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
