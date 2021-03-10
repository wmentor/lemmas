package dicts

import (
	"testing"
)

func TestDicts(t *testing.T) {
	var d Dict = mdict(map[string]bool{"1": true, "3": true})
	if d.Has("4") || !d.Has("1") || !d.Has("3") {
		t.Fatal("failed")
	}
}
