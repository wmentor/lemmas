package engine

import (
	"bufio"
	"io"
	"strings"

	"github.com/wmentor/lemmas/engine/storage"
)

func LoadWords(in io.Reader) {

	br := bufio.NewReader(in)

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		if str = strings.TrimSpace(str); len(str) > 0 {
			storage.WordAdd(str)
		}
	}
}
