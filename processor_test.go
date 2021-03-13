package lemmas

import (
	"strings"
	"testing"
)

func TestProcessor(t *testing.T) {

	tTP := func(src string, wait []string) {
		p := New()
		res := p.TextProc(strings.NewReader(src))
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

	tTP("тест", []string{"тест", "тестирование"})
	tTP("текст . тест . тест", []string{"тест", "тестирование", "текст"})
	tTP("создать экспертную систему", []string{"информационные технологии", "экспертные системы"})
	tTP("о тесте Петра Смирнова, ", []string{"петр", "петр смирнов", "смирнов", "тест", "тестирование"})
	tTP("вижу Ольгу Петрову, ", []string{"ольга", "ольга петрова", "петрова"})
	tTP("о петре VIII, ", []string{"петр viii"})

}
