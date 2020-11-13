package storage

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/wmentor/lemmas/engine/forms"
	"github.com/wmentor/lemmas/engine/words"
)

type FindResult struct {
	Form  string
	Words []*words.Word
}

var (
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

	if len(list) >= 2 {

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

	parentsData = res
}

func FormsSave(out io.Writer) {

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
	_, has := formsData[f]
	return has
}

func EachWord(f string, fn func(w *words.Word) bool) {

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

func FirstParents(f string, fn func(parent string) bool) {

	if data, has := parentsData[f]; has {

		for {

			if idx := strings.IndexByte(data, '|'); idx > 0 {

				if !fn(data[:idx]) {
					return
				}

				data = data[idx+1:]

			} else {
				fn(data)
				return
			}
		}
	}
}

func FullChain(f string, fn func(parent string) bool) {

	known := map[string]bool{f: true}

	list := []string{f}

	for {

		if len(list) == 0 {
			return
		}

		cf := list[0]
		list = list[1:]

		if !fn(cf) {
			return
		}

		FirstParents(cf, func(p string) bool {
			if !known[p] {
				list = append(list, p)
				known[p] = true
			}
			return true
		})
	}

}
