package storage

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/wmentor/log"
	"github.com/wmentor/tokens"
)

var (
	forms      map[string]string = map[string]string{}
	basicForms map[string]string = map[string]string{}
)

func LoadBasic(in io.Reader) {

	br := bufio.NewReader(in)
	res := map[string]string{}

	var list []string = make([]string, 0, 10)

	fn := func(t string) {
		list = append(list, t)
	}

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		list = list[:0]

		tokens.Process(strings.NewReader(str), fn)

		if len(list) > 1 {
			res[list[0]] = strings.Join(list[1:], " ")
		} else if len(list) == 1 {
			log.Errorf("invalid string: %s", strings.TrimSpace(str))
		}
	}

	basicForms = res
}

func GetBasicForms(form string) (string, bool) {
	str, ok := basicForms[form]
	return str, ok
}

func LoadBasicFile(filename string) error {
	if rh, e := os.Open(filename); e == nil {
		defer rh.Close()
		LoadBasic(rh)
	} else {
		return e
	}
	return nil
}

func LoadLemmas(in io.Reader) {

	br := bufio.NewReader(in)
	res := map[string]string{}

	var list []string = make([]string, 0, 10)

	fn := func(t string) {
		list = append(list, t)
	}

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		list = list[:0]

		tokens.Process(strings.NewReader(str), fn)

		if len(list) > 1 {
			res[list[0]] = strings.Join(list[1:], " ")
		} else if len(list) == 1 {
			log.Errorf("invalid string: %s", strings.TrimSpace(str))
		}
	}

	forms = res
}

func LoadLemmasFile(filename string) error {
	if rh, e := os.Open(filename); e == nil {
		defer rh.Close()
		LoadLemmas(rh)
	} else {
		return e
	}
	return nil
}

func SaveLemmas(out io.Writer) {
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	list := make([]string, 0, len(forms))

	for k := range forms {
		list = append(list, k)
	}

	sort.Sort(sort.StringSlice(list))

	for i, k := range list {
		v := forms[k]

		bw.WriteString(k)
		bw.WriteRune(' ')
		bw.WriteString(v)
		bw.WriteRune('\n')

		if (i+1)%100 == 0 {
			bw.Flush()
		}
	}
}

func SaveLemmasFile(filename string) error {
	if wh, err := os.Create(filename); err == nil {
		defer wh.Close()
		SaveLemmas(wh)
	} else {
		return err
	}
	return nil
}

func Has(form string) bool {
	_, has := forms[form]
	return has
}

func GetRaw(form string) string {
	return forms[form]
}

func EachBase(form string, callback func(string) bool) {

	if val, has := forms[form]; has {
		for {
			idx := strings.Index(val, " ")
			if idx < 0 {
				callback(val)
				return
			}
			if !callback(val[:idx]) {
				return
			}
			val = val[idx+1:]
		}
	} else if _, err := strconv.ParseInt(form, 10, 64); err == nil {
		callback(form)
	}
}

func Set(form string, bases ...string) {
	if len(bases) == 0 {
		delete(forms, form)
	} else {
		forms[form] = strings.Join(bases, " ")
	}
}

func Append(form string, base string) {
	val := forms[form]
	var list []string
	tokens.Process(strings.NewReader(val), func(t string) {
		list = append(list, t)
	})

	for _, v := range list {
		if v == base {
			return
		}
	}

	list = append(list, base)
	forms[form] = strings.Join(list, " ")
}
