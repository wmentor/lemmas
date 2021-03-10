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

	tF("w_names", "елена", true)
	tF("w_names", "елена123", false)
	tF("m_names", "михаил", true)
	tF("m_names", "___", false)
	tF("countries", "россия", true)
	tF("___", "россия", false)
}
