package storage

import (
	"strings"
	"testing"
)

func TestKeywords(t *testing.T) {

	KeywordAdd("пушкин > писатель > человек")
	KeywordAdd("орел > птица")
	KeywordAdd("орел > город")

	tKC := func(kw string, chain string) {

		var list []string

		KeywordChain(kw, func(cur string) bool {
			list = append(list, cur)
			return true
		})

		res := strings.Join(list, ">")
		if res != chain {
			t.Fatalf("KeywordChain failed for=%s expect=%s result=%s", kw, chain, res)
		}

	}

	tKC("пушкин", "пушкин>писатель>человек")
	tKC("123123123", "")
	tKC("писатель", "писатель>человек")
	tKC("орел", "орел>птица>город")
}
