package keywords

import (
	"strings"
	"testing"
)

func TestKeywords(t *testing.T) {

	tG := func(src string, wait []string) {
		ret, _ := Get(src)
		if strings.Join(ret, " ") != strings.Join(wait, " ") {
			t.Fatalf("Get failed src=%v ret=%v wait=%v", src, ret, wait)
		}
	}

	tIS := func(src string, has bool) {
		ret := Is(src)
		if ret != has {
			t.Fatalf("Is failed for %v", src)
		}
	}

	tG("формы", []string{"форма"})
	tG("тесты", []string{"тест", "тестирование"})
	tG("_____", nil)
	tG("boltdb", []string{"boltdb", "встраиваемые_хранилища", "базы_данных", "databases", "информационные_технологии"})
	tG("как_сыр_в_масле", []string{})

	tIS("формы", true)
	tIS("как_сыр_в_масле", true)
	tIS("boltdb", true)
}
