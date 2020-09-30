package link

import (
	"strings"
	"testing"
)

func compareLinks(a, b []Link) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.Href != b[i].Href {
			return false
		}
		if v.Text != b[i].Text {
			return false
		}
	}
	return true
}

func parseFile(pathfile string) ([]Link, error) {
	htmlfile, err := GetHTMLFile(pathfile)
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(htmlfile)
	links, err := Parse(r)
	if err != nil {
		return nil, err
	}
	return links, nil
}

func TestParser1(t *testing.T) {
	got, err := parseFile("testdata/test1.htm")
	if err != nil {
		t.Fatalf("Error when parsing the string: %d", err)
	}
	want := []Link{
		Link{
			Href: "/other-page",
			Text: "A link to another page",
		},
	}
	if !compareLinks(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestParser2(t *testing.T) {
	got, err := parseFile("testdata/test2.htm")
	if err != nil {
		t.Fatalf("Error when parsing the string: %d", err)
	}
	want := []Link{
		Link{
			Href: "https://www.twitter.com/joncalhoun",
			Text: "Check me out on twitter",
		},
		Link{
			Href: "https://github.com/gophercises",
			Text: "Gophercises is on Github!",
		},
	}
	if !compareLinks(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestParser3(t *testing.T) {
	got, err := parseFile("testdata/test3.htm")
	if err != nil {
		t.Fatalf("Error when parsing the string: %d", err)
	}
	want := []Link{
		Link{
			Href: "#",
			Text: "Login",
		},
		Link{
			Href: "/lost",
			Text: "Lost? Need help?",
		},
		Link{
			Href: "https://twitter.com/marcusolsson",
			Text: "@marcusolsson",
		},
	}
	if !compareLinks(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestParser4(t *testing.T) {
	got, err := parseFile("testdata/test4.htm")
	if err != nil {
		t.Fatalf("Error when parsing the string: %d", err)
	}
	want := []Link{
		Link{
			Href: "/dog-cat",
			Text: "dog cat",
		},
	}
	if !compareLinks(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
