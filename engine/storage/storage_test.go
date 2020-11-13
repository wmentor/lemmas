package storage

import (
	"bytes"
	"strings"
	"testing"

	"github.com/wmentor/lemmas/engine/words"
)

func TestStorage(t *testing.T) {

	WordAdd(`test:en.noun.sg tests:en.noun.mg`)
	WordAdd(``)
	WordAdd(`   `)
	WordAdd(`тест:ru.noun.sg.mr.ip.vp теста:ru.noun.sg.mr.rp тесту:ru.noun.sg.mr.dp тестом:ru.noun.sg.mr.tp тесте:ru.noun.sg.mr.pp`)
	WordAdd(`тесто:ru.noun.sg.sr.ip.vp теста:ru.noun.sg.sr.rp тесту:ru.noun.sg.sr.dp тестом:ru.noun.sg.sr.tp тесте:ru.noun.sg.sr.pp`)

	ParentAdd("пушкин > писатель > человек")
	ParentAdd("орел > птица")
	ParentAdd("орел > город")

	tEW := func(f string, cnt int) {
		i := 0
		EachWord(f, func(w *words.Word) bool {
			i++
			return true
		})
		if i != cnt {
			t.Fatalf("EachWord failed for: %s", f)
		}
	}

	tFC := func(f string, res string) {
		var list []string
		FullChain(f, func(v string) bool {
			list = append(list, v)
			return true
		})
		if strings.Join(list, ",") != res {
			t.Fatalf("FullChain failed for: %s res: %s", f, strings.Join(list, ","))
		}
	}

	tEW("tests", 1)
	tEW("теста", 2)

	buf := bytes.NewBuffer(nil)
	FormsSave(buf)
	if buf.Len() == 0 {
		t.Fatalf("FormsSave failed")
	}

	tFC("орел", "орел,птица,город")
	tFC("писатель", "писатель,человек")
	tFC("пушкин", "пушкин,писатель,человек")
}
