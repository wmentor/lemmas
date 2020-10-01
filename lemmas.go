package lemmas

import (
	"os"
	"strings"

	"github.com/wmentor/lemmas/storage"
)

var (
	filename string
)

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

func ProcForm(form string) string {

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

				storage.EachCurBase(suf, func(f string) bool {
					if has {
						maker.WriteRune(' ')
					}
					maker.WriteString(pref)
					maker.WriteString(f)
					has = true
					return true
				})

				return maker.String()
			}
		}
	}

	return form
}
