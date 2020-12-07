package forms

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

var (
	data      map[string]string = map[string]string{}
	signs     map[string]bool
	ruLetters string = "йцукенгшщзхъфывапролджэёячсмитьбю"
	enLetters string = "qwertyuiopasdfghjklzxcvbnm"
	skLetters string = "-" + ruLetters + enLetters
)

func init() {
	signs = map[string]bool{}

	for _, s := range strings.Fields(". , ! ? - : ; ( ) [ ] { } \" ' + / & % « » < > =") {
		signs[s] = true
	}
}

func hasSpecial(f string) bool {
	if signs[f] {
		return true
	}

	if _, err := strconv.ParseInt(f, 10, 64); err == nil {
		return true
	}

	if strings.IndexAny(f, ".:/_@#'") > -1 {
		return true
	}

	if strings.IndexAny(f, "0123456789") > -1 && strings.IndexAny(f, skLetters) > -1 {
		return true
	}

	return false
}

func Has(f string) bool {
	if _, has := data[f]; has {
		return has
	}

	if hasSpecial(f) {
		return true
	}

	return false
}

func EachBase(f string, fn func(string) bool) {

	if str, has := data[f]; has {
		for _, base := range strings.Fields(str) {
			if !fn(base) {
				return
			}
		}
	} else if hasSpecial(f) {
		fn(f)
	}
}

func CurBase(f string, fn func(string) bool) {
	if fn(f) {
		EachBase(f, func(v string) bool {
			if v != f {
				return fn(v)
			}
			return true
		})
	}
}

func Add(cur string, base string) {
	cur = strings.ToLower(cur)
	base = strings.ToLower(base)

	if ds, has := data[cur]; has {

		for _, v := range strings.Fields(ds) {
			if v == base {
				return
			}
		}

		data[cur] = ds + " " + base

	} else {
		data[cur] = base
	}
}

func Reset() {
	data = make(map[string]string)
}

func Load(in io.Reader) {
	br := bufio.NewReader(in)

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		if list := strings.Fields(str); len(list) > 1 {
			f := list[0]
			for _, base := range list[1:] {
				Add(f, base)
			}
		}
	}
}

func Save(out io.Writer) {
	keys := make([]string, 0, len(data))

	for f := range data {
		keys = append(keys, f)
	}

	sort.Sort(sort.StringSlice(keys))

	for _, f := range keys {
		fmt.Fprintf(out, "%s %s\n", f, data[f])
	}
}
