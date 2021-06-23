package dicts

import (
	"testing"
)

func TestLoader(t *testing.T) {

	tF := func(dict string, name string, wait bool) {
		if InDict(dict, name) != wait {
			t.Fatalf("InDict(%v,%v) != %v", dict, name, wait)
		}
	}

	tF("wnames", "елена", true)
	tF("wnames", "елена123", false)
	tF("mnames", "михаил", true)
	tF("mnames", "___", false)
	tF("countries", "россия", true)
	tF("___", "россия", false)
}
