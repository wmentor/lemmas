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
	forms     map[string]string = map[string]string{}
	fixed     map[string]string = map[string]string{}
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

	if strings.HasPrefix(f, "по-") {
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

	if _, has := fixed[f]; has {
		return true
	}

	for i, _ := range f {
		if _, has := forms[f[i:]]; has {
			return true
		}
	}

	if idx := strings.IndexByte(f, '-'); idx > -1 {
		if Has(f[:idx]) && Has(f[idx+1:]) {
			return true
		}
	}

	if hasSpecial(f) {
		return true
	}

	return false
}

func EachBase(f string, fn func(string) bool) {

	if str, has := forms[f]; has {
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

func add(hash map[string]string, cur string, base string) {
	cur = strings.ToLower(cur)
	base = strings.ToLower(base)

	if ds, has := hash[cur]; has {

		for _, v := range strings.Fields(ds) {
			if v == base {
				return
			}
		}

		hash[cur] = ds + " " + base

	} else {
		hash[cur] = base
	}
}

func Add(cur string, base string) {
	add(forms, cur, base)
}

func AddFixed(cur string, base string) {
	add(fixed, cur, base)
}

func Reset() {
	forms = make(map[string]string)
	fixed = make(map[string]string)
}

func load(dest map[string]string, in io.Reader) {
	br := bufio.NewReader(in)

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		if list := strings.Fields(str); len(list) > 1 {
			f := list[0]
			for _, base := range list[1:] {
				add(dest, f, base)
			}
		}
	}
}

func LoadForms(in io.Reader) {
	load(forms, in)
}

func LoadFixed(in io.Reader) {
	load(fixed, in)
}

func save(src map[string]string, out io.Writer) {
	keys := make([]string, 0, len(forms))

	for f := range src {
		keys = append(keys, f)
	}

	sort.Sort(sort.StringSlice(keys))

	for _, f := range keys {
		fmt.Fprintf(out, "%s %s\n", f, src[f])
	}
}

func SaveForms(out io.Writer) {
	save(forms, out)
}

func SaveFixed(out io.Writer) {
	save(fixed, out)
}
