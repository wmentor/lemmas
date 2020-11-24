package engine

import (
	"bufio"
	"io"
	"strings"

	"github.com/wmentor/lemmas/engine/meta"
)

func LoadMeta(in io.Reader) {

	writeAccess(func() {

		br := bufio.NewReader(in)

		for {
			str, err := br.ReadString('\n')
			if err != nil && str == "" {
				break
			}

			if str = strings.TrimSpace(str); len(str) > 0 {
				meta.Add(str)
			}
		}

		needSave = true
	})
}
