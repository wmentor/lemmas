package storage

import (
	"testing"

	"github.com/wmentor/kv"
)

func TestStorage(t *testing.T) {

	kv.Open("test=1")
	defer Close()

	for i := int64(1); i < 100; i++ {
		if nextId() != i {
			t.Fatalf("NextId failed for: %d", i)
		}
	}
}
