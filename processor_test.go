package lemmas

import (
	"strings"
	"testing"
)

func TestProcessor(t *testing.T) {

	tTP := func(src string, wait []string) {
		p := New()
		p.AddText(strings.NewReader(src))
		var res []string
		p.FetchResult(func(kw string, w float64) {
			res = append(res, kw)
		})
		if len(res) != len(wait) {
			t.Fatalf("TextProc failed for: %s return: %v", src, res)
		}
		if strings.Join(res, " ") != strings.Join(wait, " ") {
			t.Fatalf("TextProc failed for: %s return: %v", src, res)
		}
	}

	tTH := func(src string, wait []string) {
		p := New()
		p.AddHTML(strings.NewReader(src))
		var res []string
		p.FetchResult(func(kw string, w float64) {
			res = append(res, kw)
		})
		if len(res) != len(wait) {
			t.Fatalf("TextProc failed for: %s return: %v", src, res)
		}
		if strings.Join(res, " ") != strings.Join(wait, " ") {
			t.Fatalf("TextProc failed for: %s return: %v", src, res)
		}
	}

	tTP("тест", []string{"тест", "тестирование"})
	tTP("текст . тест . тест", []string{"тест", "тестирование", "текст"})
	tTP("создать экспертную систему", []string{"информационные технологии", "экспертные системы"})
	tTP("о тесте Петра Смирнова, ", []string{"петр", "петр смирнов", "смирнов", "тест", "тестирование"})
	tTP("вижу Ольгу Петрову, ", []string{"ольга", "ольга петрова", "петрова"})
	tTP("о петре VIII, ", []string{"петр viii"})

	tTH("<html><body><p>о петре&nbsp;I</p></body></html>", []string{"петр i"})
}
