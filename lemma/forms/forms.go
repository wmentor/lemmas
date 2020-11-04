package forms

import (
	"strings"

	"github.com/wmentor/lemmas/lemma/opts"
)

type Form struct {
	Name string
	Opts opts.Opts
}

func New(txt string) *Form {

	if idx := strings.Index(txt, ":"); idx > 0 {
		return &Form{
			Name: txt[:idx],
			Opts: opts.New(txt[idx+1:]),
		}
	}

	return nil
}

func (f *Form) String() string {
	return f.Name + ":" + f.Opts.String()
}
