package main

import (
	"fmt"
	"strings"

	"github.com/wmentor/lemmas"
)

func main() {

	txt := "Создать экспертную систему."

	processor := lemmas.New()

	processor.AddText(strings.NewReader(txt))

	processor.FetchResult(func(keyphrase string, weight float64) {
		fmt.Println(keyphrase, weight)
	})

	// reset and reinit object
	processor.Reset()

	// process data from HTML
	html := "<html><body><p>о петре&nbsp;I</p></body></html>"

	processor.AddHTML(strings.NewReader(html))

	processor.FetchResult(func(keyphrase string, weight float64) {
		fmt.Println(keyphrase, weight)
	})
}
