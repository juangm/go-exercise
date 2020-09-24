package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

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

	// Create custom template
	tpl := template.Must(template.New("").Parse(storyTmpl))

	h := adventure.NewHandler(story, adventure.WithTemplate(tpl), adventure.WithPathFunc(pathFn))
	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

// Custom pathFn
func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	// Remove the '/' from the path
	return path[len("/story/"):]
}

// Custom template
var storyTmpl = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Choose Your Own Adventure</title>
  </head>
  <section class="page">
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}} {{if .Options}}
    <ul>
      {{range .Options}}
      <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
    </ul>
    {{else}}
    <h3>The End</h3>
    {{end}}
  </section>
  <style>
    body {
      font-family: helvetica, arial;
    }
    h1 {
      text-align: center;
      position: relative;
    }
    .page {
      width: 80%;
      max-width: 500px;
      margin: auto;
      margin-top: 40px;
      margin-bottom: 40px;
      padding: 80px;
      background: #fffcf6;
      border: 1px solid #eee;
      box-shadow: 0 10px 6px -6px #777;
    }
    ul {
      border-top: 1px dotted #ccc;
      padding: 10px 0 0 0;
      -webkit-padding-start: 0;
    }
    li {
      padding-top: 10px;
    }
    a,
    a:visited {
      text-decoration: none;
      color: #6295b5;
    }
    a:active,
    a:hover {
      color: #7792a2;
    }
    p {
      text-indent: 1em;
    }
  </style>
</html>`
