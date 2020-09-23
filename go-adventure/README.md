# GO - CHOOSE YOUR OWN ADVENTURE

![Status](https://img.shields.io/badge/Status-InProgress-orange)
![Bonus](https://img.shields.io/badge/Bonus-OnHold-blue)

## Description

- Exercise #3 from [gophercises](https://gophercises.com/)
- Implement a go program to recreate the experience of [Choose Your Own Adventure](https://en.wikipedia.org/wiki/Choose_Your_Own_Adventure)
- Implement a web application where each page will be a portion of the story
- Stories will be provided via a JSON file - [example](https://github.com/gophercises/cyoa/blob/master/gopher.json)
- Stories can be cyclical
- All stories will have a story arc named `intro` (starting point)
- Matt Holt's JSON-to-Go is a really handy tool when working with [JSON in Go!](https://mholt.github.io/json-to-go/)

## Work To Do

- [ ] Use the `html/template` package to create your HTML pages
- [ ] Create `http.Handler` to handle web requests (instead of a handler function)
- [ ] Use the `encoding/json` package to decode the JSON file

## Bonus points

- [ ] Create a command line version
- [ ] How you would alter your program in order to support stories starting form a story-defined arc?
- [ ] How would you refactor/redesign the program or restructure the JSON?
