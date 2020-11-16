package storage

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

var (
	kws map[string]string
)

func init() {
	kws = make(map[string]string)
}

func KeywordsLoad(in io.Reader) {
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

	kws = res
}

func KeywordsSave(out io.Writer) {

	buf := bytes.NewBuffer(nil)

	for ps, pd := range kws {
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

func keywordAdd(src string, parent string) {

	if src = strings.TrimSpace(src); src != "" {
		parent = strings.TrimSpace(parent)

		curData := kws[src]

		if len(curData) > 0 {

			for _, v := range strings.Split(curData, "|") {
				if v == parent {
					return
				}
			}

			kws[src] = curData + "|" + parent

		} else {
			kws[src] = parent
		}
	}
}

func KeywordAdd(txt string) {

	list := strings.Split(strings.ReplaceAll(strings.ToLower(txt), "ё", "е"), ">")

	prev := ""

	for _, v := range list {
		cur := strings.TrimSpace(v)
		if cur != "" {
			keywordAdd(cur, cur)
			if prev != "" {
				keywordAdd(prev, cur)
			}
			prev = cur
		}
	}
}

func keywordFirstIterator(kw string, fn func(string) bool) {
	if data, has := kws[kw]; has {

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

func KeywordChain(kw string, fn func(string) bool) {

	if _, has := kws[kw]; !has {
		return
	}

	known := map[string]bool{kw: true}
	list := []string{kw}

	for {

		if len(list) == 0 {
			return
		}

		cf := list[0]
		list = list[1:]

		if !fn(cf) {
			return
		}

		keywordFirstIterator(cf, func(p string) bool {
			if !known[p] {
				list = append(list, p)
				known[p] = true
			}
			return true
		})
	}

}
