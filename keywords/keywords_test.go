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

	tG("формы", []string{"форма"})
	tG("тесты", []string{"тест", "тестирование"})
	tG("_____", nil)
	tG("boltdb", []string{"boltdb", "встраиваемые_хранилища", "базы_данных", "databases"})
}
