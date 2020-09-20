package log

import (
	"testing"
)

func TestLog(t *testing.T) {

	str := ""

	SetLogger(func(lvl string, msg string) {
		str = lvl + "|" + msg
	})

	tF := func(wait string, lvl string, format string, args ...interface{}) {
		str = ""
		Log(lvl, format, args...)
		if str != wait {
			t.Fatalf("expect return=%s wait=%s", str, wait)
		}
	}

	tF("ERROR|Hello, Mike", "ERROR", "Hello, %s", "Mike")
}
