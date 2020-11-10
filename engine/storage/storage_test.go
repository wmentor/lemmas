package storage

import (
	"testing"
)

func TestStorage(t *testing.T) {

	TestOpen()
	defer Close()

	for i := int64(1); i < 100; i++ {
		if nextId() != i {
			t.Fatalf("NextId failed for: %d", i)
		}
	}

	wid := AddWord("test:en.noun.sg tests:en.noun.mg")
	if wid != 100 {
		t.Fatalf("AddWord failed")
	}

	if AddWord("") != 0 {
		t.Fatalf("AddWord must return 0")
	}

	w := getWord(wid)
	if w == nil {
		t.Fatalf("GetWord failed")
	}

	AddWord("тест:ru.noun.sg.mr.ip.vp теста:ru.noun.sr.mr.rp тесту:ru.noun.sg.mr.dp тестом:ru.noun.sg.mr.tp тесте:ru.noun.sg.mr.pp")

	fr := Find("тестом")
	if fr == nil {
		t.Fatalf("Find result nil for: %s", "тестом")
	}

	if len(fr.Words) != 1 {
		t.Fatalf("Find must return 1 word for: %s", "тестом")
	}

	AddWord("тесто:ru.noun.sg.sr.ip.vp теста:ru.noun.sg.sr.rp тесту:ru.noun.sg.sr.dp тестом:ru.noun.sg.sr.tp тесте:ru.noun.sg.sr.pp")

	fr = Find("тестом")
	if fr == nil {
		t.Fatalf("Find result nil for: %s", "тестом")
	}

	if len(fr.Words) != 2 {
		t.Fatalf("Find must return 2 word for: %s", "тестом")
	}

	DelWord(wid)

	if getWord(wid) != nil {
		t.Fatalf("GetWord must return nil")
	}
}
