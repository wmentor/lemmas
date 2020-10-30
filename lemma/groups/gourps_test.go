package groups

import (
	"strings"
	"testing"

	"github.com/wmentor/lemmas/lemma/forms"
)

func TestGroups(t *testing.T) {

	data := `
  noun.sg.mr человек человека человеку человека человеком человеке
  noun.mg люди людей людям людей людьми людях
  `

	Load(strings.NewReader(data))

	if !forms.Has("людей") {
		t.Fatal("word not found")
	}
}
