package lemmas

import (
	"io"
	"strings"

	"github.com/wmentor/lemmas/buffer"
	"github.com/wmentor/lemmas/stat"
	"github.com/wmentor/tokens"
)

type Keyword = stat.StatRecord //nolint

const (
	bufferSize int = 5
)

var (
	eos map[string]bool = map[string]bool{".": true, "?": true, "!": true, "…": true}
)

func TextProc(in io.Reader) []*Keyword {

	st := stat.New()
	buf := buffer.New(bufferSize)

	maker := strings.Builder{}

	tact := func() {

		if eos[buf.Get(0)] {
			st.EndTact()
			buf.Shift(1)
			return
		}

		maker.Reset()

		buf.Shift(1)
	}

	tokens.Process(in, func(t string) {

		buf.Push(t)

		if buf.Full() {
			tact()
		}

	})

	for !buf.Empty() {
		tact()
	}

	return st.Result()
}
