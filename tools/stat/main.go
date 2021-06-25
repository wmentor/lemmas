package main

import (
	"fmt"

	"github.com/wmentor/lemmas/dicts"
	"github.com/wmentor/lemmas/forms"
	"github.com/wmentor/lemmas/keywords"
)

func main() {

	fmt.Printf("total forms: %d\n", forms.Size())
	fmt.Printf("total words in dicts: %d\n", dicts.Size())
	fmt.Printf("total keywords and templates: %d\n", keywords.Size())
}
