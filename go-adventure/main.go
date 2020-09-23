package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/juangm/go-exercises/go-adventure/adventure"
)

func main() {
	port := flag.Int("port", 3000, "The port to start the web application.")
	filename := flag.String("file", "gopher.json", "The JSON file with the adventure story.")
	flag.Parse()
	fmt.Printf("Using the story in %s. \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		// Not a good idea...
		panic(err)
	}
	story, err := adventure.JSONStory(f)
	if err != nil {
		panic(err)
	}

	h := adventure.NewHandler(story)
	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
