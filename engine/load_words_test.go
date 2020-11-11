package engine

import (
	"strings"
	"testing"

)

func TestLoadWords(t *testing.T) {

	data := `
test:en.noun.sg tests:en.noun.mg
тест:ru.noun.sg.mr.ip.vp теста:ru.noun.sg.mr.rp тесту:ru.noun.sg.mr.dp тестом:ru.noun.mr.sg.tp тесте:ru.noun.mr.sg.pp
тесты:ru.noun.mg.ip.vp тестов:ru.noun.mg.rp тестам:ru.noun.mg.dp тестами:ru.noun.mg.tp тестах:ru.noun.mg.pp
`

	LoadWords(strings.NewReader(data))

}
