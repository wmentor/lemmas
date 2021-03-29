# lemmas

Библиотека для анализа текстов на чистом Go

![test](https://github.com/wmentor/lemmas/workflows/test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/wmentor/lemmas/badge.svg?branch=master&v=20210323)](https://coveralls.io/github/wmentor/lemmas?branch=master)
[![https://goreportcard.com/report/github.com/wmentor/lemmas](https://goreportcard.com/badge/github.com/wmentor/lemmas)](https://goreportcard.com/report/github.com/wmentor/lemmas)
[![https://pkg.go.dev/github.com/wmentor/lemmas](https://pkg.go.dev/badge/github.com/wmentor/lemmas.svg)](https://pkg.go.dev/github.com/wmentor/lemmas)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

## Кратко

* Требуется Go >= 1.16
* Написана на чистом Go
* Лицензия MIT

## Установка

```plaintext
go get -u github.com/wmentor/lemmas
```

## Пример использования

```golang
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
  fmt.Printf("Время на прочтение: %d\n", processor.ReadingTime() )
}
```
