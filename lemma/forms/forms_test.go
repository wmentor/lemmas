package forms

import (
	"strings"
	"testing"
)

func TestForms(t *testing.T) {

	tA := func(src, srcF, base, baseF string, wait string) {
		Add(src, srcF, base, baseF)
		if src != "" && !Has(src) {
			t.Fatalf("Add failed for: %s", src)
		}
		if list, _ := Get(src); strings.Join(list, sep) != wait {
			t.Fatalf("Get failed for: %s", src)
		}
	}

	tA("t1", "1", "t2", "2", "1|t2|2")
	tA("t1", "", "t3", "", "1|t2|2|0|t3|0")
	tA("", "1", "2", "", "")

	if _, ok := Get("12312313"); ok {
		t.Fatal("Get unknown value failed")
	}

	Reset()

	if _, ok := Get("t1"); ok {
		t.Fatal("Get unknown value failed")
	}
}
