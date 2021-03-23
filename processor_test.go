package lemmas

import (
	"strings"
	"testing"
)

func TestProcessor(t *testing.T) {

	tTP := func(src string, wait []string, tokCnt int64) {
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
		if p.Tokens() != tokCnt {
			t.Fatalf("Tokens return wrong number %v expect %v", p.Tokens(), tokCnt)
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

	tTP("тест", []string{"тест", "тестирование"}, 1)
	tTP("текст . тест . тест", []string{"тест", "тестирование", "текст"}, 5)
	tTP("создать экспертную систему", []string{"информационные технологии", "экспертные системы"}, 3)
	tTP("о тесте Петра Смирнова, ", []string{"петр", "петр смирнов", "смирнов", "тест", "тестирование"}, 5)
	tTP("вижу Ольгу Петрову, ", []string{"ольга", "ольга петрова", "петрова"}, 4)
	tTP("о петре VIII, ", []string{"петр viii"}, 4)
	tTP("разные вакцины", []string{"вакцины", "лекарства", "медицина"}, 2)
	tTP("#футбол", []string{"спорт", "футбол"}, 1)
	tTP("он как сыр в масле", []string{}, 5)
	tTP("без сыра", []string{"еда", "сыр"}, 2)
	tTP("владимиром путиным", []string{"политика", "путин", "россия"}, 2)

	tTH("<html><body><p>о петре&nbsp;I</p></body></html>", []string{"петр i"})
}
