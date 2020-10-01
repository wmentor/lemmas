package lemmas

import (
	"testing"
)

func TestLemmas(t *testing.T) {
	if Open("") != nil {
		t.Fatal("Open failed")
	}
}
