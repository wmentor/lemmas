package lemmas

import (
	"strconv"
	"strings"

	"github.com/wmentor/lemmas/keywords"
)

type state struct {
	Size   int
	Result string
	Vars   []string
	proc   *processor
}

func newState(p *processor) *state {
	return &state{
		Size:   0,
		Result: "",
		Vars:   make([]string, 0, bufferSize),
		proc:   p,
	}
}

func (st *state) Len() int {
	return st.Size
}

func (st *state) PushPhrase(size int, value string, vars []string) {
	if size > st.Size {
		st.Result = value
		st.Size = size
		st.Vars = st.Vars[:size]
		copy(st.Vars, vars)
	} else if st.Size == size && st.Result != value {
		dataRes := st.proc.getKeywordData(st.Result)
		dataCV := st.proc.getKeywordData(value)
		if strings.Join(dataRes, ";") != strings.Join(dataCV, ";") {
			st.Result = "" // indeterminacy
		}
		st.Vars = st.Vars[:0]
	}
}

func (st *state) Reset() {
	st.Size = 0
	st.Result = ""
	st.Vars = st.Vars[:0]
}

func (st *state) Keywords() []string {
	if st.Result == "" {
		return nil
	}

	list, has := keywords.Get(st.Result)
	if !has {
		return nil
	}

	result := make([]string, 0, len(list))
	for _, v := range list {
		if strings.Contains(v, "$") {
			cur := v
			for i, val := range st.Vars {
				src := "$" + strconv.Itoa(i+1)
				cur = strings.ReplaceAll(cur, src, val)
			}
			if !strings.Contains(cur, "$") {
				result = append(result, cur)
			}
		} else {
			result = append(result, v)
		}
	}

	return result
}
