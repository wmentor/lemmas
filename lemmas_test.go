package lemmas

import (
	"testing"
)

func TestLemmas(t *testing.T) {
	if Open("") != nil {
		t.Fatal("Open failed")
	}

	tPF := func(src string, wait string) {
		if res := ProcessForm(src); res != wait {
			t.Fatalf("ProcessForm(%s) != %s", src, wait)
		}
	}

	tPF(".", ".")
	tPF(",", ",")
	tPF("тигру", "тигр")
	tPF("тигров", "тигры")
	tPF("кибертигром", "кибертигр")
	tPF("тигром-12", "тигр-12")
	tPF("летчиком-испытателем", "летчик-испытатель")
	tPF("летчик-испытатель", "летчик-испытатель")
	tPF("налетчику", "налетчик")
}
