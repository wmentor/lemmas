package forms

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

var (
	data map[string]string = map[string]string{}
)

func Has(f string) bool {
	_, has := data[f]
	return has
}

func EachBase(f string, fn func(string) bool) {
	for _, base := range strings.Fields(data[f]) {
		if !fn(base) {
			return
		}
	}
}

func CurBase(f string, fn func(string) bool) {
	if fn(f) {
		EachBase(f, func(v string) bool {
			if v != f {
				return fn(v)
			}
			return true
		})
	}
}

func Add(cur string, base string) {
	cur = strings.ToLower(cur)
	base = strings.ToLower(base)

	if ds, has := data[cur]; has {

		for _, v := range strings.Fields(ds) {
			if v == base {
				return
			}
		}

		data[cur] = ds + " " + base

	} else {
		data[cur] = base
	}
}

func AddWord(txt string) {
	list := strings.Fields(strings.ToLower(txt))
	if len(list) > 0 {
		base := list[0]
		for _, f := range list {
			Add(f, base)
		}
	}
}

func Reset() {
	data = make(map[string]string)
}

func Load(in io.Reader) {
	br := bufio.NewReader(in)

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		if list := strings.Fields(str); len(list) > 1 {
			f := list[0]
			for _, base := range list[1:] {
				Add(f, base)
			}
		}
	}
}

func Save(out io.Writer) {
	keys := make([]string, 0, len(data))

	for f := range data {
		keys = append(keys, f)
	}

	sort.Sort(sort.StringSlice(keys))

	for _, f := range keys {
		fmt.Fprintf(out, "%s %s\n", f, data[f])
	}
}
