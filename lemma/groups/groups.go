package groups

import (
	"bufio"
	"io"
	"strings"

	"github.com/wmentor/lemmas/lemma/flags"
	"github.com/wmentor/lemmas/lemma/forms"
	"github.com/wmentor/tokens"
)

func Load(in io.Reader) {
	forms.Reset()

	br := bufio.NewReader(in)

	var list []string = make([]string, 0, 16)

	ff := func(t string) {
		list = append(list, t)
	}

	for {

		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		tokens.Process(strings.NewReader(str), ff)

		if len(list) < 2 {
			continue
		}

		f := flags.New(list[0])

		if f&flags.F_NOUN != 0 {

		}
	}
}
