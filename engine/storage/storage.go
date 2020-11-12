package storage

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"sync"

	"github.com/wmentor/lemmas/engine/forms"
	"github.com/wmentor/lemmas/engine/words"
)

type FindResult struct {
	Form  string
	Words []*words.Word
}

var (
	mt sync.RWMutex

	formsData   map[string]string
	parentsData map[string]string
)

func init() {
	formsData = make(map[string]string)
	parentsData = make(map[string]string)
}

func formAdd(f *forms.Form, baseForm *forms.Form) {

	curData := formsData[f.Name]

	maker := strings.Builder{}

	if len(curData) > 0 {
		maker.WriteString(curData)
		maker.WriteRune('|')
	}

	maker.WriteString(baseForm.String())
	maker.WriteRune(' ')
	maker.WriteString(f.String())

	formsData[f.Name] = maker.String()
}

func parentAdd(src string, parent string) {

	if src = strings.TrimSpace(src); src != "" {
		parent = strings.TrimSpace(parent)

		curData := parentsData[src]

		if len(curData) > 0 {

			for _, v := range strings.Split(curData, "|") {
				if v == parent {
					return
				}
			}

			parentsData[src] = curData + "|" + parent

		} else {
			parentsData[src] = parent
		}
	}
}

func ParentAdd(txt string) {

	list := strings.Split(strings.ToLower(txt), ">")

	if len(list) > 2 {

		src := ""

		for _, v := range list {

			val := strings.TrimSpace(v)
			if val == "" {
				continue
			}

			if src != "" {
				parentAdd(src, val)
			}

			src = val
		}
	}
}

func WordAdd(wstr string) bool {

	if w := words.New(wstr); w != nil {

		mt.Lock()
		defer mt.Unlock()

		for _, f := range w.Forms {
			formAdd(f, w.Forms[0])
		}

		return true
	}

	return false
}

func FormsLoad(in io.Reader) {
	br := bufio.NewReader(in)

	res := make(map[string]string)

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		if str = strings.TrimSpace(str); str != "" {
			if idx := strings.IndexByte(str, '|'); idx > 0 {
				res[str[:idx]] = str[idx+1:]
			}
		}
	}

	mt.Lock()
	defer mt.Unlock()

	formsData = res
}

func ParentsLoad(in io.Reader) {
	br := bufio.NewReader(in)

	res := make(map[string]string)

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		if str = strings.TrimSpace(str); str != "" {
			if idx := strings.IndexByte(str, '|'); idx > 0 {
				res[str[:idx]] = str[idx+1:]
			}
		}
	}

	mt.Lock()
	defer mt.Unlock()

	parentsData = res
}

func FormsSave(out io.Writer) {

	mt.RLock()
	mt.RUnlock()

	buf := bytes.NewBuffer(nil)

	for f, data := range formsData {
		if data != "" {
			buf.WriteString(f)
			buf.WriteByte('|')
			buf.WriteString(data)
			buf.WriteByte('\n')
			out.Write(buf.Bytes())
			buf.Reset()
		}
	}
}

func ParentsSave(out io.Writer) {

	mt.RLock()
	mt.RUnlock()

	buf := bytes.NewBuffer(nil)

	for ps, pd := range parentsData {
		if pd != "" {
			buf.WriteString(ps)
			buf.WriteByte('|')
			buf.WriteString(pd)
			buf.WriteByte('\n')
			out.Write(buf.Bytes())
			buf.Reset()
		}
	}
}

func FormHas(f string) bool {
	mt.RLock()
	defer mt.RUnlock()

	_, has := formsData[f]
	return has
}

func EachWord(f string, fn func(w *words.Word) bool) {

	mt.RLock()
	defer mt.RUnlock()

	if data, has := formsData[f]; has {
		for {
			if idx := strings.IndexByte(data, '|'); idx > 0 {
				if w := words.New(data[:idx]); w != nil {
					if !fn(w) {
						break
					}
				}
				data = data[idx+1:]
			} else {
				if w := words.New(data); w != nil {
					fn(w)
				}
				break
			}
		}
	}
}
