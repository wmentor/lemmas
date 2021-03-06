package counter

import (
	"sort"
	"strings"
	"testing"
)

func TestMapCnt(t *testing.T) {

	m := New()

	tW := func(key string, list []string) {
		m.Inc(key)
		strs := []string{}
		for _, v := range m.Keys() {
			strs = append(strs, v.Name)
		}
		if strings.Join(strs, " ") != strings.Join(list, " ") {
			t.Fatalf("failed res=%v wait=%s", strs, list)
		}
	}

	tW("1", []string{"1"})
	tW("1", []string{"1"})
	tW("0", []string{"1", "0"})
	tW("2", []string{"1", "0", "2"})
	tW("2", []string{"1", "2", "0"})
	tW("2", []string{"2", "1", "0"})
	tW("5", []string{"2", "1", "0", "5"})

	list := m.KeyNames()
	sort.Strings(list)
	if strings.Join(list, "") != "0125" {
		t.Fatal("KeyNames failed")
	}

	m.Reset()
	tW("5", []string{"5"})
}
