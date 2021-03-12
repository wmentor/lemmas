package lemmas

import (
	"io"

	"github.com/wmentor/lemmas/buffer"
	"github.com/wmentor/lemmas/dicts"
	"github.com/wmentor/lemmas/forms"
	"github.com/wmentor/lemmas/keywords"
	"github.com/wmentor/lemmas/stat"
	"github.com/wmentor/tokens"
)

type Keyword = stat.Record //nolint

const (
	bufferSize int = 5
)

var (
	eos map[string]bool = map[string]bool{".": true, "?": true, "!": true, "…": true}
)

func TextProc(in io.Reader) []Keyword {

	st := stat.New()
	buf := buffer.New(bufferSize)

	localKeywords := make(map[string][]string)

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

				if deep == 1 && buf.Len() > 1 {

					isName := false

					if dicts.InDict("m_names", val) {
						isName = true
						wn := buf.Get(1)
						forms.Each(wn, func(fn string) bool {
							if dicts.InDict("m_lastnames", fn) {
								res = val + "_" + fn
								if _, has := localKeywords[res]; !has {
									localKeywords[res] = []string{res, fn, val}
								}
								cmpPhrase(2, res)
								return false
							}
							return true
						})
					}

					if dicts.InDict("w_names", val) {
						isName = true
						wn := buf.Get(1)
						forms.Each(wn, func(fn string) bool {
							if dicts.InDict("w_lastnames", fn) {
								res = val + "_" + fn
								if _, has := localKeywords[res]; !has {
									localKeywords[res] = []string{res, fn, val}
								}
								cmpPhrase(2, res)
								return false
							}
							return true
						})
					}

					if isName {
						wn := buf.Get(1)
						if dicts.InDict("roman", wn) {
							res = val + "_" + wn
							if _, has := localKeywords[res]; !has {
								localKeywords[res] = []string{res}
							}
							cmpPhrase(2, res)
						}
					}
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
				if list, has := localKeywords[res]; has {
					for _, v := range list {
						st.AddKey(v)
					}
				} else {
					list, _ := keywords.Get(res)
					for _, v := range list {
						st.AddKey(v)
					}
				}
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
