package storage

import (
	"bufio"
	"io"
	"strings"

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
			res[list[0]] = strings.Join(list[1:], "|")
		}
	}

	data = res
}
