package lemmas

import (
	"bytes"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/wmentor/html"
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
	tokensCounter int64
	imageCounter  int64
	currentWords  []string
	replacing     []string
}

// make new text processor
func New() Processor {
	p := &processor{}
	p.Reset()
	return p
}

// calc reading time in seconds
func (p *processor) ReadingTime() int64 {
	sum := p.tokensCounter + p.imageCounter*2
	add := int64(0)
	if sum%3 != 0 {
		add = 1
	}
	res := sum/3 + add
	return res
}

// process input text via io.Reader
func (p *processor) AddText(in io.Reader) {

	st := newState(p)

	tokens.Process(in, func(t string) {

		p.tokensCounter++
		p.buf.Push(t)

		if p.buf.Full() {
			p.tact(st)
		}

	})

	for !p.buf.Empty() {
		p.tact(st)
	}

	p.stat.EndTact()
}

// return current tokens counter value
func (p *processor) Tokens() int64 {
	return p.tokensCounter
}

// process input html via io.Reader
func (p *processor) AddHTML(in io.Reader) {
	parser := html.New()

	parser.Parse(in)

	parser.EachImage(func(img string) {
		p.imageCounter++
	})

	p.AddText(bytes.NewReader(parser.Text()))
}

// process input html from web page
func (p *processor) AddURL(url string) error {

	parser := html.New()

	opts := &html.GetOpts{
		Timeout: time.Second * 10,
		Agent:   "Lemmas",
	}

	if err := parser.Get(url, opts); err != nil {
		return err
	}

	parser.EachImage(func(img string) {
		p.imageCounter++
	})

	p.AddText(bytes.NewReader(parser.Text()))

	return nil
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
			dataRes := p.getKeywordData(res)
			dataCV := p.getKeywordData(cv)
			if strings.Join(dataRes, ";") != strings.Join(dataCV, ";") {
				res = "" // indeterminacy
			}
		}
	}

	forms.Each(str, func(f string) bool {
		val := cur + f
		p.currentWords[deep-1] = f

		if sr, ss := p.search(val, deep+1); ss > 0 {
			cmpPhrase(ss, sr)
			return true
		}

		dicts.Each(f, func(f string) bool {

			val := cur + f

			if sr, ss := p.search(val, deep+1); ss > 0 {
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

		if size <= deep {

			if ok := keywords.Is(val); ok {
				cmpPhrase(deep, val)
			}

		}

		return true
	})

	return res, size
}

// one buffer process tact
func (p *processor) tact(st *state) {
	if eos[p.buf.Get(0)] {
		p.stat.EndTact()
		p.buf.Shift(1)
		return
	}

	st.Reset()

	if res, num := p.search("", 1); num > 0 {
		if res != "" {
			for _, v := range p.getKeywordData(res) {
				if idx := strings.IndexByte(v, '$'); idx >= 0 {
					for i, rep := range p.replacing[:num] {
						v = strings.ReplaceAll(v, rep, p.currentWords[i])
					}
				}
				p.stat.AddKey(v)
			}
		}
		p.buf.Shift(num)
		return
	}

	p.buf.Shift(1)
}

// get keyword data (used with local and global keywords)
func (p *processor) getKeywordData(kw string) []string {
	if list, has := p.localKeywords[kw]; has {
		return list
	}
	if list, has := keywords.Get(kw); has {
		return list
	}
	return nil
}

// fetch results
func (p *processor) FetchResult(fn EachResultFunc) {
	p.stat.Result(fn)
}

func (p *processor) Reset() {
	p.stat = stat.New()
	p.buf = buffer.New(bufferSize)
	p.tokensCounter = 0
	p.imageCounter = 0
	p.localKeywords = make(map[string][]string)
	p.currentWords = make([]string, bufferSize)
	p.replacing = make([]string, bufferSize)

	for i := range p.replacing {
		p.replacing[i] = "$" + strconv.Itoa(i)
	}
}
