package meta

import (
	"bytes"
	"strings"
	"testing"
)

func TestMeta(t *testing.T) {

	Reset()

	txt := `
    #россия Россия
    #михаил Михаил
    #игра игра
    `

	Load(strings.NewReader(txt))

	if Get("#игра") != "игра" {
		t.Fatal("Get failed")
	}

	buf := bytes.NewBuffer(nil)
	Save(buf)

	if buf.String() != `#игра игра
#михаил Михаил
#россия Россия
` {
		t.Fatal("Save failed")
	}
}
