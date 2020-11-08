package words

import (
	"testing"

	"github.com/wmentor/lemmas/engine/opts"
)

func TestWords(t *testing.T) {

	tF := func(txt string, o opts.Opts) {
		if w := New(txt); w != nil {
			if w.String() != txt {
				t.Fatalf("String failed for: %s", txt)
			}
			if string(w.Bytes()) != txt {
				t.Fatalf("Bytes failed for: %s", txt)
			}
			if w.Opts != o {
				t.Fatalf("Invalid opts for: %s", txt)
			}
		} else {
			t.Fatalf("New failed for: %s", txt)
		}
	}

	tF("test:en.noun.sg tests:en.noun.mg", opts.O_EN|opts.O_NOUN)
	tF("в:ru.pret", opts.O_RU|opts.O_PRET)

	if New("") != nil {
		t.Fatal("New must failed")
	}

	if New("\t  \t   ") != nil {
		t.Fatal("New must failed")
	}
}
