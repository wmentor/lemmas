package meta

import (
	"bytes"
	"strings"
	"testing"
)

func TestMeta(t *testing.T) {

	Reset()

	txt := `
	москва =город =россия =столица =москва
	россия =страна =россия
	рф =страна =россия
	российская федерация =страна =россия
	птица =птица
	сорока =птица
	москва =город
	`

	Load(strings.NewReader(txt))

	tH := func(src string, has bool) {
		if Has(src) != has {
			t.Fatalf("Has failed for: %s", src)
		}
	}

	tH("132314_23141234", false)
	tH("российская федерация", true)
	tH("москва", true)

	buf := bytes.NewBuffer(nil)
	Save(buf)

	wait := `москва =город =россия =столица =москва
птица =птица
российская федерация =страна =россия
россия =страна =россия
рф =страна =россия
сорока =птица
`
	if buf.String() != wait {
		t.Fatal("Save failed")
	}
}
