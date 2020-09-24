package adventure

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
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
  <section class="page">
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}} {{if .Options}}
    <ul>
      {{range .Options}}
      <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
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

type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

// HandlerOpts s
type HandlerOpts func(h *handler)

// WithTemplate s
func WithTemplate(t *template.Template) HandlerOpts {
	return func(h *handler) {
		h.t = t
	}
}

// WithPathFunc s
func WithPathFunc(fn func(r *http.Request) string) HandlerOpts {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	// Remove the '/' from the path
	return path[1:]
}

// NewHandler will construct and http.Handler that will render the story provided
func NewHandler(s Story, opts ...HandlerOpts) http.Handler {
	h := handler{s, tpl, defaultPathFn}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)
	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
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
	Paragraphs []string  `json:"story"`
	Options    []Options `json:"options"`
}

// Options available in the story
type Options struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
