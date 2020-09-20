package storage

import (
	"bufio"
	"io"
)

var (
	data map[string]string = map[string]string{}
)

func Load(in io.Reader) {

	br := bufio.NewReader(in)
	res := map[string]string{}

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}
	}

	data = res
}
