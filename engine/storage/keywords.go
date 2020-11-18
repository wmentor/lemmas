package storage

import (
	"bufio"
	"io"
	"strings"
)

var (
	kws map[string]bool
)

func init() {
	kws = make(map[string]bool)
}

func KeywordsLoad(in io.Reader) {
	br := bufio.NewReader(in)

	res := make(map[string]bool)

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		if str = strings.Join(strings.Fields(str), " "); str != "" {
			res[str] = true
		}
	}

	kws = res
}

func KeywordsSave(out io.Writer) {

	for kw := range kws {
		out.Write([]byte(kw + "\n"))
	}
}

func KeywordAdd(txt string) {

	txt = strings.ReplaceAll(strings.ToLower(txt), "ё", "е")

	if txt = strings.Join(strings.Fields(txt), " "); txt != "" {
		kws[txt] = true
	}
}

func IsKeyword(txt string) bool {
	return kws[txt]
}
