# GO - URL SHORTENER

![Status](https://img.shields.io/badge/Status-InProgres-orange)
![Bonus](https://img.shields.io/badge/Bonus-OnHold-blue)

## Description

- Exercise #2 from [gophercises](https://gophercises.com/)
- create an http.Handler that will look at the path of any incoming web request and determine if it should redirect the user to a new page.
- Basically like an URL shortener.

## Work To Do

- [ ] Implement the stubbed method `MAPHandler` in `handler.go`
- [ ] Implement the stubbed method `YAMLHandler` in `handler.go`
- [ ] Use the package [yaml](gopkg.in/yaml.v2) to parse the YAML.

## Bonus points

- [ ] Update the `main/main.go` file to accept YAML files as flag.
- [ ] Build `JSONHandler` for the same purpose, but reads from JSON data.
- [ ] Build a handler that reads from a database (Up to to you to use `BoltDB`, `SQL`, etc)
