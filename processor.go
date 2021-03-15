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

// text processor type
type processor struct {
	buf           buffer.Buffer
	stat          stat.Stat
	localKeywords map[string][]string
}

// make new text processor
func New() Processor {
	p := &processor{}
	p.Reset()
	return p
}

// process input text via io.Reader
func (p *processor) AddText(in io.Reader) {

	tokens.Process(in, func(t string) {

		p.buf.Push(t)

		if p.buf.Full() {
			p.tact()
		}

	})

	for !p.buf.Empty() {
		p.tact()
	}

	p.stat.EndTact()
}

// search keywords from word stream
func (p *processor) search(cur string, deep int) (string, int) {
	if deep > p.buf.Len() {
		return "", 0
	}

	str := p.buf.Get(deep - 1)

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

		if sr, ss := p.search(val, deep+1); ss > 0 {
			cmpPhrase(ss, sr)
			return true
		}

		if size <= deep {

			if ok := keywords.Is(val); ok {
				cmpPhrase(deep, val)
			}

			if deep == 1 && p.buf.Len() > 1 {

				isName := false

				if dicts.InDict("m_names", val) {
					isName = true
					wn := p.buf.Get(1)
					forms.Each(wn, func(fn string) bool {
						if dicts.InDict("m_lastnames", fn) {
							res = val + "_" + fn
							if _, has := p.localKeywords[res]; !has {
								p.localKeywords[res] = []string{res, fn, val}
							}
							cmpPhrase(2, res)
							return false
						}
						return true
					})
				}

				if dicts.InDict("w_names", val) {
					isName = true
					wn := p.buf.Get(1)
					forms.Each(wn, func(fn string) bool {
						if dicts.InDict("w_lastnames", fn) {
							res = val + "_" + fn
							if _, has := p.localKeywords[res]; !has {
								p.localKeywords[res] = []string{res, fn, val}
							}
							cmpPhrase(2, res)
							return false
						}
						return true
					})
				}

				if isName {
					wn := p.buf.Get(1)
					if dicts.InDict("roman", wn) {
						res = val + "_" + wn
						if _, has := p.localKeywords[res]; !has {
							p.localKeywords[res] = []string{res}
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

// one buffer process tact
func (p *processor) tact() {
	if eos[p.buf.Get(0)] {
		p.stat.EndTact()
		p.buf.Shift(1)
		return
	}

	if res, num := p.search("", 1); num > 0 {
		if res != "" {
			if list, has := p.localKeywords[res]; has {
				for _, v := range list {
					p.stat.AddKey(v)
				}
			} else {
				list, _ := keywords.Get(res)
				for _, v := range list {
					p.stat.AddKey(v)
				}
			}
		}
		p.buf.Shift(num)
		return
	}

	p.buf.Shift(1)
}

// fetch results
func (p *processor) FetchResult(fn EachResultFunc) {
	p.stat.Result(fn)
}

func (p *processor) Reset() {
	p.stat = stat.New()
	p.buf = buffer.New(bufferSize)
	p.localKeywords = make(map[string][]string)
}
