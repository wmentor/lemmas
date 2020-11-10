package words

import (
	"bytes"
	"strings"

	"github.com/wmentor/lemmas/engine/forms"
	"github.com/wmentor/lemmas/engine/opts"
)

type Word struct {
	Id    int64
	Base  string
	Opts  opts.Opts
	Forms []*forms.Form
}

func New(str string) *Word {

	w := &Word{
		Opts: opts.O_MASK,
	}

	i := 0

	for _, fl := range strings.Fields(str) {
		if f := forms.New(strings.ToLower(fl)); f != nil {
			i++
			w.Forms = append(w.Forms, f)
			w.Opts = w.Opts & f.Opts
			if i == 1 {
				w.Base = f.Name
			}
		}
	}

	if len(w.Forms) == 0 {
		return nil
	}

	return w
}

func (w *Word) String() string {
	maker := bytes.NewBuffer(nil)

	for i, f := range w.Forms {
		if i > 0 {
			maker.WriteRune(' ')
		}
		maker.WriteString(f.String())
	}

	return maker.String()
}

func (w *Word) Bytes() []byte {
	maker := bytes.NewBuffer(nil)

	for i, f := range w.Forms {
		if i > 0 {
			maker.WriteRune(' ')
		}
		maker.WriteString(f.String())
	}

	return maker.Bytes()
}
