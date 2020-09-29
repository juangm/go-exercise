package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/juangm/go-exercises/go-htmlParser/link"
)

func main() {
	file := flag.String("file", "./link/testdata/test1.htm", "The Path to the HTML file to be parsed.")
	flag.Parse()

	htmlfile, err := link.GetHTMLFile(*file)
	if err != nil {
		panic(err)
	}

	r := strings.NewReader(htmlfile)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
