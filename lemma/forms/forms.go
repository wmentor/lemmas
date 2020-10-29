package forms

import (
	"strings"
)

const (
	sep string = "|"
)

var (
	data map[string]string = map[string]string{}
)

func Reset() {
	data = make(map[string]string)
}

func Add(src string, srcFlags, base string, baseFlags string) {

	if src == "" || base == "" {
		return
	}

	maker := strings.Builder{}

	if str, has := data[src]; has {
		maker.WriteString(str)
		maker.WriteString(sep)
	}

	if srcFlags != "" {
		maker.WriteString(srcFlags)
	} else {
		maker.WriteRune('0')
	}

	maker.WriteString(sep)
	maker.WriteString(base)
	maker.WriteString(sep)

	if baseFlags != "" {
		maker.WriteString(baseFlags)
	} else {
		maker.WriteRune('0')
	}

	data[src] = maker.String()
}

func Get(src string) ([]string, bool) {
	if res, ok := data[src]; ok {
		return strings.Split(res, sep), true
	}
	return nil, false
}

func Has(src string) bool {
	_, has := data[src]
	return has
}
