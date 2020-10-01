package main

import (
	"flag"
	"net/http"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for.")
	flag.Parse()

	// 1. GET the webpage
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 2. Parse the links on the page (using previous package)

	// 3. Build proper urls with our links
	// 4. Filter out any links with a diff domain.
	// 5. Find all pages (BFS)
	// 6. Print out XML
}
