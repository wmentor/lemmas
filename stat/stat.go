package stat

import (
	"github.com/wmentor/lemmas/counter"
)

// text keyword statistics record
type StatRecord = counter.Key

// text keyword statistics collector
type Stat interface {
	AddKey(string)
	EndTact()
	Result() []*StatRecord
}

type stat struct {
	curCnt counter.Counter
	allCnt counter.Counter
}

// create new Stat collector
func New() Stat {
	return &stat{
		curCnt: counter.New(),
		allCnt: counter.New(),
	}
}

// add key (keyword/phrase from currect text statement)
func (a *stat) AddKey(name string) {
	if name[0] != '%' {
		a.curCnt.Inc(name)
	}
}

// close currect tact (text statement)
func (a *stat) EndTact() {
	for _, k := range a.curCnt.KeyNames() {
		a.allCnt.Inc(k)
	}
	a.curCnt.Reset()
}

// result text stat result
func (a *stat) Result() []*StatRecord {
	a.EndTact()
	return a.allCnt.Keys()
}
