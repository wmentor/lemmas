package engine

import (
	"io"

	"github.com/wmentor/lemmas/engine/forms"
)

func LoadWords(in io.Reader) {

	writeAccess(func() {

		forms.Load(in)
		needSave = true

	})
}
