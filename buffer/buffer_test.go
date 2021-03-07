package buffer

import (
	"strings"
	"testing"
)

func TestBuffer(t *testing.T) {

	buf := New(5)

	tPush := func(val string, wait []string) {
		buf.Push(val)
		if buf.Len() != len(wait) {
			t.Fatal("invalid buffer sLen")
		}
		if buf.Len() == 0 && !buf.Empty() || buf.Len() > 0 && buf.Empty() {
			t.Fatal("Empty faild")
		}
		if buf.Len() == 5 && !buf.Full() || buf.Len() < 5 && buf.Full() {
			t.Fatal("Full failed")
		}
		var list []string
		for i := 0; i < buf.Len(); i++ {
			list = append(list, buf.Get(i))
		}
		if strings.Join(list, ":") != strings.Join(wait, ":") {
			t.Fatal("invalid result")
		}
	}

	tShift := func(n int, wait []string) {
		buf.Shift(n)
		if buf.Len() != len(wait) {
			t.Fatal("invalid buffer sLen")
		}
		if buf.Len() == 0 && !buf.Empty() || buf.Len() > 0 && buf.Empty() {
			t.Fatal("Empty faild")
		}
		if buf.Len() == 5 && !buf.Full() || buf.Len() < 5 && buf.Full() {
			t.Fatal("Full failed")
		}
		var list []string
		for i := 0; i < buf.Len(); i++ {
			list = append(list, buf.Get(i))
		}
		if strings.Join(list, ":") != strings.Join(wait, ":") {
			t.Fatal("invalid result")
		}
	}

	tPush("1", []string{"1"})
	tPush("2", []string{"1", "2"})
	tPush("3", []string{"1", "2", "3"})
	tPush("4", []string{"1", "2", "3", "4"})
	tPush("5", []string{"1", "2", "3", "4", "5"})
	tPush("6", []string{"2", "3", "4", "5", "6"})
	tPush("7", []string{"3", "4", "5", "6", "7"})
	tPush("8", []string{"4", "5", "6", "7", "8"})
	tShift(1, []string{"5", "6", "7", "8"})
	tShift(2, []string{"7", "8"})
	tPush("9", []string{"7", "8", "9"})
	tShift(3, []string{})
	tPush("7", []string{"7"})
	tPush("5", []string{"7", "5"})
	tPush("3", []string{"7", "5", "3"})
	tShift(2, []string{"3"})
}
