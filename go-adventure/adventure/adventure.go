package adventure

import (
	"encoding/json"
	"io"
)

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
