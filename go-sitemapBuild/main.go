package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/juangm/go-exercises/go-htmlParser/link"
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
	links, err := link.Parse(resp.Body)
	if err != nil {
		fmt.Printf("Some error happened when parsing the links!")
	}

	// 3. Build proper urls with our links
	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()

	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		default:
			fmt.Println("Not taking the URL: ", l.Href)
		}
	}
	// 4. Filter out any links with a diff domain.
	// 5. Find all pages (BFS)
	// 6. Print out XML
}
