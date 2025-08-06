package main

import (
	"fmt"
	"os"

	"github.com/kenn-teera/note-taking-app/convert"
)

/*
Scope

features:
	Users can upload markdown files
	Check the grammar
	Save the note
	Render it in HTML <Use Package gomarkdown>
*/

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := os.ReadFile("markdown/markdown-1.md")
	check(err)
	fmt.Print(string(dat))

	// 	var mds = `
	// 	# header
	// 	Sample text.
	// 	[link](http://example.com)
	// `
	//r := routers.NewRouter()
	md := []byte(dat)
	html := convert.MdToHTML(md)
	fmt.Printf("--- Markdown:\n%s\n\n--- HTML:\n%s\n", md, html)
}
