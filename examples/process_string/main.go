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

	// обнуляем состояние объекта для переиспользования
	processor.Reset()

	// обработка данных в виде HTML
	html := "<html><body><p>о петре&nbsp;I</p></body></html>"

	processor.AddHTML(strings.NewReader(html))

	processor.FetchResult(func(keyphrase string, weight float64) {
		fmt.Println(keyphrase, weight)
	})

	// получаем информацию об обработанных данных
	fmt.Printf("Число токенов: %d\n", processor.Tokens())
	fmt.Printf("Время на прочтение: %d\n", processor.ReadingTime())
}
