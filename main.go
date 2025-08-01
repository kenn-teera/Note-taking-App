package main

import (
	"fmt"

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

func main() {
	var mds = `
	# header
	Sample text.
	[link](http://example.com)
`
	//r := routers.NewRouter()
	md := []byte(mds)
	html := convert.MdToHTML(md)
	fmt.Printf("--- Markdown:\n%s\n\n--- HTML:\n%s\n", md, html)
}
