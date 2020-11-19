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

func Set(meta string, base string) {
	data[meta] = base
}

func Get(meta string) string {
	return data[meta]
}

func Load(in io.Reader) {
	br := bufio.NewReader(in)

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		if list := strings.Fields(str); len(list) > 1 {
			data[list[0]] = strings.Join(list[1:], " ")
		}
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
