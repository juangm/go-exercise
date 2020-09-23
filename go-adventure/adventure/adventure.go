package adventure

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
