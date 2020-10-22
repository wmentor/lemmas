package lemma

import (
	"strings"
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
		var res []string
		EachBase(src, func(l string) bool {
			res = append(res, l)
			return true
		})
		if strings.Join(res, " ") != wait {
			t.Fatalf("EachBase failed for: %s", src)
		}
	}

	tCP := func(src string, wait bool) {
		if CanProcess(src) != wait {
			t.Fatalf("CanProcess(%s) != %t", src, wait)
		}
	}

	tCEB := func(src string, wait string) {
		var res []string
		CurEachBase(src, func(l string) bool {
			res = append(res, l)
			return true
		})
		if strings.Join(res, " ") != wait {
			t.Fatalf("EachBase failed for: %s", src)
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
	tPF("игры", "игра игры")
	tPF("ли", "ли")

	tCP(",", true)
	tCP("тигр", true)
	tCP("тигра", true)
	tCP("летчиком-испытателем", true)
	tCP("летчик-испытатель", true)
	tCP("тигр-", true)
	tCP("123123123", true)
	tCP("___.12312331", false)
	tCP("налетчиком", true)
	tCP("же", true)

	tCEB("налетчиком", "налетчиком налетчик")
	tCEB("летчиком-испытателем", "летчиком-испытателем летчик-испытатель")
}
