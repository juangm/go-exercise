# GO - SITEBUILD MAP

![Status](https://img.shields.io/badge/Status-InProgress-Orange)
![Bonus](https://img.shields.io/badge/Bonus-OnHold-blue)

## Description

- Exercise #5 from [gophercises](https://gophercises.com/)
- Sitemap -> Basically a map of all the pages within a specific domain
- Build it visiting the root page and making a list of every link (do it recursively)

## Work To Do

- [ ] Start the program with the URL provided by flag
- [ ] Output of the program should be in XML
- [ ] Be aware that links can be cyclical
- [ ] Use `net/http` to initiate GET request to each page
- [ ] Use `encoding/xml` to print out the XML output at the end
- [ ] Use `flag` to parse user provided flags

## Bonus points

- [ ] Add in `depth` flag that defines the maximum number of links to follow when building the sitemap.
