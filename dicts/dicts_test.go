package dicts

import (
	"testing"
)

func TestDict(t *testing.T) {

	d := dict(map[string]int{"test": 1, "dict": 2, "word": 3})

	tF := func(word string, wait int) {
		if d.Get(word) != wait {
			t.Fatalf("Dict.Get failed for: %s", word)
		}
	}

	tF("test", 1)
	tF("dict", 2)
	tF("word", 3)
	tF("someword", 0)

	tD := func(dict string, word string, wait int) {
		d, err := GetDict(dict)
		if err != nil {
			if wait != 0 {
				t.Fatalf("GetDict failed for: %s", dict)
			}
			return
		}

		if d.Get(word) != wait {
			t.Fatalf("dict.Get failed for %v %v", dict, word)
		}
	}

	tD("1231231", "1231", 0)
	tD("names", "михаил", 1)
}
