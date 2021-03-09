package lemmas

import (
	"strings"
	"testing"
)

func TestLemmas(t *testing.T) {

	tTP := func(src string, wait []string) {
		res := TextProc(strings.NewReader(src))
		if len(res) != len(wait) {
			t.Fatalf("TextProc failed for: %s return: %v", src, res)
		}
		list := make([]string, len(res))
		for i, v := range res {
			list[i] = v.Name
		}
		if strings.Join(list, " ") != strings.Join(wait, " ") {
			t.Fatalf("TextProc failed for: %s return: %v", src, list)
		}
	}

	tTP("тест", []string{"тест"})
}
