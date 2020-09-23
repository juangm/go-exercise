package adventure

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandler))
}

var tpl *template.Template

var defaultHandler = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
    {{range .Options}}
    <li><a href="{{.Chapter}}">{{.Text}}</li>
    {{end}}
    </ul>
  </body>
</html>`

type handler struct {
	s Story
}

// NewHandler will construct and http.Handler that will render the story provided
func NewHandler(s Story) http.Handler {
	h := handler{s}
	return h
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

// JSONStory parse JSON file into a map
func JSONStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

// Story type
type Story map[string]Chapter

// Chapter of the story
type Chapter struct {
	Title      string    `json:"title"`
	Paragraphs []string  `json:"paragrahs"`
	Options    []Options `json:"options"`
}

// Options available in the story
type Options struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
