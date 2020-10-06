package lemmas

import (
	"io"
	"os"
	"strings"

	"github.com/wmentor/lemmas/storage"
	"github.com/wmentor/tokens"
)

var (
	filename string
	signs    map[string]bool
)

func init() {
	signs = map[string]bool{
		".":  true,
		",":  true,
		"?":  true,
		"!":  true,
		";":  true,
		":":  true,
		"\"": true,
		"'":  true,
		"-":  true,
	}
}

type LemmaFunc func(string)

func Open(src string) error {

	if src == "" {
		if filename = os.Getenv("WMENTOR_LEMMAS_DB"); filename == "" {
			filename = "lemmas.db"
		}
	} else {
		filename = src
	}

	return storage.LoadFile(filename)
}

func Save() error {
	return storage.SaveFile(filename)
}

func AddForm(form string, bases ...string) {
	for _, base := range bases {
		storage.Append(form, base)
	}
}

func DelForm(form string) {
	storage.Set(form)
}

func ProcessForm(form string) string {

	if signs[form] {
		return form
	}

	for pos, _ := range form {
		if pos == 0 {
			if v := storage.GetRaw(form); v != "" {
				return v
			}
		} else {
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
