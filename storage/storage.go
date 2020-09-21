package storage

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/wmentor/lemmas/log"
	"github.com/wmentor/tokens"
)

var (
	data map[string]string = map[string]string{}
)

func Load(in io.Reader) {

	br := bufio.NewReader(in)
	res := map[string]string{}

	var list []string = make([]string, 0, 10)

	fn := func(t string) {
		list = append(list, t)
	}

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		list = list[:0]

		tokens.Process(strings.NewReader(str), fn)

		if len(list) > 1 {
			res[list[0]] = strings.Join(list[1:], " ")
		} else if len(list) == 1 {
			log.Log("ERROR", "invalid string: %s", strings.TrimSpace(str))
		}
	}

	data = res
}

func Save(out io.Writer) {
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	list := make([]string, 0, len(data))

	for k := range data {
		list = append(list, k)
	}

	sort.Sort(sort.StringSlice(list))

	for _, k := range list {
		v := data[k]
		fmt.Fprintf(bw, "%s %s\n", k, v)
	}
}
