package engine

import (
	"io"

	"github.com/wmentor/lemmas/engine/forms"
)

func LoadWords(in io.Reader) {

	writeAccess(func() {

		forms.LoadForms(in)
		needSave = true

	})
}

func LoadFixed(in io.Reader) {

	writeAccess(func() {

		forms.LoadFixed(in)
		needSave = true

	})
}
