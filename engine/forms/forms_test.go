package forms

import (
	"bytes"
	"strings"
	"testing"
)

func TestForms(t *testing.T) {
	Reset()

	txt := `
	тест тест
	теста тест тесто
	тесту тест тесто
	тестом тест тесто
	тесте тест тесто
	театр театр
	театра театр
	театру театр
	театром театр
	театре театр
	`

	Load(strings.NewReader(txt))

	tHas := func(f string, wait bool) {
		if Has(f) != wait {
			t.Fatalf("Has failed for: %s", f)
		}
	}

	AddWord("огонь огня огню огонь огнем огне")

	tHas("теста", true)
	tHas("театр", true)
	tHas("огню", true)
	tHas(")))!231", false)

	buf := bytes.NewBuffer(nil)

	Save(buf)

	if buf.String() != `огне огонь
огнем огонь
огню огонь
огня огонь
огонь огонь
театр театр
театра театр
театре театр
театром театр
театру театр
тест тест
теста тест тесто
тесте тест тесто
тестом тест тесто
тесту тест тесто
` {
		t.Fatal("Save failed")
	}
}
