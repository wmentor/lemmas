package stat

import (
	"strings"

	"github.com/wmentor/lemmas/counter"
)

type EachResultFunc func(key string, frequency float64)

// text keyword statistics collector
type Stat interface {
	AddKey(string)
	EndTact()
	Result(EachResultFunc)
}

type stat struct {
	tact   int
	used   map[string]int
	curCnt []string
	allCnt counter.Counter
}

// create new Stat collector
func New() Stat {
	return &stat{
		tact:   1,
		used:   make(map[string]int),
		curCnt: make([]string, 0, 32),
		allCnt: counter.New(),
	}
}

// add key (keyword/phrase from currect text statement)
func (a *stat) AddKey(name string) {
	if a.tact != a.used[name] {
		a.curCnt = append(a.curCnt, name)
		a.used[name] = a.tact
	}
}

// close currect tact (text statement)
func (a *stat) EndTact() {
	a.tact++
	for _, k := range a.curCnt {
		a.allCnt.Inc(k)
	}
	a.curCnt = a.curCnt[:0]
}

// result text stat result
func (a *stat) Result(fn EachResultFunc) {
	a.EndTact()

	total := float64(a.allCnt.Total())

	a.allCnt.EachFreq(func(key string, val int64) {
		key = strings.ReplaceAll(key, "_", " ")
		freq := float64(val) / total
		fn(key, freq)
	})
}
