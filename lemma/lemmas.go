package lemma

import (
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/wmentor/lemmas/lemma/storage"
	"github.com/wmentor/log"
	"github.com/wmentor/tokens"
)

var (
	dir     string
	signs   map[string]bool
	tinyMap map[string]string
)

func init() {
	signs = map[string]bool{
		"":   true,
		".":  true,
		",":  true,
		"?":  true,
		"!":  true,
		";":  true,
		":":  true,
		"\"": true,
		"'":  true,
		"-":  true,
		"«":  true,
		"»":  true,
		"[":  true,
		"]":  true,
		"(":  true,
		")":  true,
		"{":  true,
		"}":  true,
		"%":  true,
	}

	tinyMap = map[string]string{}
}

type LemmaFunc func(string)

func Open(src string) {

	if src == "" {
		if dir = os.Getenv("LEMMAS_DATA"); dir == "" {
			if dir = os.Getenv("GOPATH"); dir != "" {
				dir = dir + "/src/github.com/wmentor/lemmas/data"
			} else {
				dir = "./data"
			}
		}
	} else {
		dir = src
	}

	if err := storage.LoadLemmasFile(dir + "/lemmas.db"); err != nil {
		log.Errorf("load %s/lemmas.db failed: %s", dir, err.Error())
	}

	if err := storage.LoadBasicFile(dir + "/basic.db"); err != nil {
		log.Errorf("load %s/basic.db failed: %s", dir, err.Error())
	}
}

func Save() error {
	return storage.SaveLemmasFile(dir + "/lemmas.db")
}

func AddForm(form string, bases ...string) {
	for _, base := range bases {
		storage.Append(form, base)
	}
}

func DelForm(form string) {
	storage.Set(form)
}

func CanProcess(form string) bool {

	if signs[form] {
		return true
	}

	if _, ok := storage.GetBasicForms(form); ok {
		return true
	}

	if strings.IndexAny(form, ".:/#_@") >= 0 {
		return true
	}

	return canProcess(form)
}

func canProcess(form string) bool {

	if form == "" {
		return true
	}

	if storage.Has(form) {
		return true
	}

	if _, err := strconv.ParseInt(form, 10, 64); err == nil {
		return true
	}

	for idx, run := range form {
		if idx != 0 {
			if run == '-' {
				return canProcess(form[:idx]) && canProcess(form[idx+1:])
			} else {
				suf := form[idx:]
				if storage.Has(suf) {
					return true
				}
			}
		}
	}

	return false
}

func EachBase(form string, fn func(string) bool) {
	res := ProcessForm(form)
	for _, v := range strings.Split(res, " ") {
		if !fn(v) {
			return
		}
	}
}

func CurEachBase(form string, fn func(string) bool) {
	if !fn(form) {
		return
	}

	res := ProcessForm(form)
	for _, v := range strings.Split(res, " ") {
		if v != form && !fn(v) {
			return
		}
	}
}

func ProcessForm(form string) string {

	if signs[form] {
		return form
	}

	if res, ok := storage.GetBasicForms(form); ok {
		return res
	}

	return processForm(form)
}

func processForm(form string) string {

	if form == "" {
		return ""
	}

	if v := storage.GetRaw(form); v != "" {
		return v
	}

	for pos, run := range form {
		if pos > 0 {

			if run == '-' {
				f1 := processForm(form[:pos])
				f2 := processForm(form[pos+1:])

				fl1 := strings.Split(f1, " ")
				fl2 := strings.Split(f2, " ")

				maker := strings.Builder{}
				has := false

				for _, w1 := range fl1 {
					for _, w2 := range fl2 {
						if has {
							maker.WriteRune(' ')
						}
						has = true
						maker.WriteString(w1)
						maker.WriteRune('-')
						maker.WriteString(w2)
					}
				}

				return maker.String()
			}

			suf := form[pos:]
			if storage.Has(suf) {
				pref := form[:pos]

				maker := strings.Builder{}
				has := false

				storage.EachBase(suf, func(f string) bool {
					if has {
						maker.WriteRune(' ')
					}
					maker.WriteString(pref)
					maker.WriteString(f)
					has = true
					return true
				})

				if has {
					return maker.String()
				}
			}
		}
	}

	return form
}

func Process(in io.Reader, fn LemmaFunc) {
	tokens.Process(in, func(t string) {
		fn(ProcessForm(t))
	})
}

func Stream(in io.Reader) <-chan string {
	out := make(chan string, 2048)

	go func() {
		defer close(out)
		Process(in, func(w string) {
			out <- w
		})
	}()

	return out
}
