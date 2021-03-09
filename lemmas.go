package lemmas

import (
	"io"

	"github.com/wmentor/lemmas/buffer"
	"github.com/wmentor/lemmas/forms"
	"github.com/wmentor/lemmas/keywords"
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

	var search func(cur string, deep int) (string, int)

	search = func(cur string, deep int) (string, int) {

		if deep > buf.Len() {
			return "", 0
		}

		str := buf.Get(deep - 1)

		if deep > 1 {
			cur += "_"
		}

		res := ""
		size := 0

		cmpPhrase := func(cs int, cv string) {
			if cs > size {
				res = cv
				size = cs
			} else if cs == size && cv != res {
				res = "" // indeterminacy
			}
		}

		forms.Each(str, func(f string) bool {
			val := cur + f

			if sr, ss := search(val, deep+1); ss > 0 {
				cmpPhrase(ss, sr)
				return true
			}

			if size <= deep {
				if ok := keywords.Is(val); ok {
					cmpPhrase(deep, val)
				}
			}

			return true
		})

		return res, size
	}

	tact := func() {

		if eos[buf.Get(0)] {
			st.EndTact()
			buf.Shift(1)
			return
		}

		if res, num := search("", 1); num > 0 {
			if res != "" {
				st.AddKey(res)
			}
			buf.Shift(num)
			return
		}

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
