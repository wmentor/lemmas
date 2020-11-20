package meta

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

func Reset() {
	data = map[string]string{}
}

func add(meta string, tag string) {
	ds := data[meta]
	if ds != "" {
		for _, v := range strings.Fields(ds) {
			if v == tag {
				return
			}
		}
		data[meta] = ds + " " + tag
	} else {
		data[meta] = tag
	}
}

func Add(str string) {
	list := strings.Fields(strings.ToLower(str))

	src := strings.Builder{}
	var tags []string

	for _, w := range list {
		if w != "" {
			if w[0] == '=' {
				tags = append(tags, w)
			} else {
				if src.Len() > 0 {
					src.WriteRune(' ')
				}
				src.WriteString(w)
			}
		}
	}

	if src.Len() > 0 {
		s := src.String()
		for _, t := range tags {
			add(s, t)
		}
	}
}

func Has(meta string) bool {
	_, has := data[meta]
	return has
}

func Load(in io.Reader) {
	br := bufio.NewReader(in)

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		Add(str)
	}
}

func Save(out io.Writer) {
	keys := make([]string, 0, len(data))

	for k := range data {
		keys = append(keys, k)
	}

	sort.Sort(sort.StringSlice(keys))

	for _, k := range keys {
		fmt.Fprintf(out, "%s %s\n", k, data[k])
	}
}

func EachMeta(item string, fn func(string) bool) {
	for _, base := range strings.Fields(data[item]) {
		if !fn(base) {
			return
		}
	}
}

func ItemAndMeta(item string, fn func(string) bool) {
	if fn(item) {
		EachMeta(item, func(meta string) bool {
			if meta != item {
				return fn(meta)
			}
			return true
		})
	}
}
