package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for.")
	flag.Parse()

	fmt.Println(*urlFlag)

	// 1. GET the webpage
	// 2. Parse the links on the page (using previous package)
	// 3. Build proper urls with our links
	// 4. Filter out any links with a diff domain.
	// 5. Find all pages (BFS)
	// 6. Print out XML
}
