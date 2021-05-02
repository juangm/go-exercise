package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/juangm/go-exercises/go-htmlParser/link"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for.")
	flag.Parse()

	pages := get(*urlFlag)
	for _, page := range pages {
		fmt.Println(page)
	}

	// 5. Find all pages (BFS)
	// 6. Print out XML
}

func hrefs(body io.Reader, base string) []string {
	links, err := link.Parse(body)
	if err != nil {
		fmt.Printf("Some error happened when parsing the links!")
	}
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		default:
			fmt.Println("Not taking the URL: ", l.Href)
		}
	}
	return ret
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 3. Build proper urls with our links
	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	return filter(base, hrefs(resp.Body, base))
}

func filter(base string, links []string) []string {
	var ret []string
	for _, link := range links {
		// Only links with same base URL
		if strings.HasPrefix(link, base) {
			ret = append(ret, link)
		}
	}
	return ret
}
