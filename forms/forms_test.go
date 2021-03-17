package forms

import (
	"strings"
	"testing"
)

func TestForms(t *testing.T) {
	tE := func(src, wait string) {
		var list []string
		Each(src, func(str string) bool {
			list = append(list, str)
			return true
		})
		if strings.Join(list, " ") != wait {
			t.Fatalf("Each failed for: %s", src)
		}
	}

	tE("формы", "формы форма")
	tE("форм", "форм формы")
	tE("форма", "форма")
	tE("#михаил", "#михаил михаил")
}
