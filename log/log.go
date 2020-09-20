package log

import (
	"fmt"
)

// logger func type
type LogFunc func(level string, msg string)

var (
	lf LogFunc
)

func SetLogger(fn LogFunc) {
	lf = fn
}

func Log(level string, format string, args ...interface{}) {
	if lf != nil {
		lf(level, fmt.Sprintf(format, args...))
	}
}
