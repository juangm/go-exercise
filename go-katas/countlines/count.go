package count

import (
	"bufio"
	"log"
	"os"
)

// lines function - count the lines of a text file
func lines(path string) int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	// defer keyword - to ensure operations always happen
	defer f.Close()

	sc := bufio.NewScanner(f)
	var lines int
	for sc.Scan() {
		lines++
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
