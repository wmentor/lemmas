package stat

import (
	"strings"
	"testing"
)

func TestStat(t *testing.T) {

	st := New()

	tF := func(in []string, wait []string) {
		for _, key := range in {
			st.AddKey(key)
		}

		var list []string

		for _, key := range st.Result() {
			list = append(list, key.Name)
		}

		if strings.Join(list, " ") != strings.Join(wait, " ") {
			t.Fatalf("Expect: %v", wait)
		}
	}

	tF([]string{"1", "2", "3", "2"}, []string{"1", "2", "3"})
	tF([]string{"2", "3", "4"}, []string{"2", "3", "1", "4"})
	tF([]string{"0", "12", "3", "4"}, []string{"3", "2", "4", "0", "1", "12"})
}
