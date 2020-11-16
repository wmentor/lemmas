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
	formsData map[string]string
)

func init() {
	formsData = make(map[string]string)
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
