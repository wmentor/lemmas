package forms

import (
	"testing"
)

func TestForms(t *testing.T) {

	tF := func(txt string) {
		f := New(txt)
		if f == nil {
			t.Fatalf("New failed for: %s", txt)
		}

		if f.String() != txt {
			t.Fatalf("String failed for: %s", txt)
		}
	}

	tF("тест:ru.noun.sg.mr.ip")
}
