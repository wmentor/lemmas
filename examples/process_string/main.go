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

	processor.FetchResult(func(keyword string, weight float64) {
		fmt.Println(keyword, weight)
	})
}
