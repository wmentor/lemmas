package words

import (
	"strings"

	"github.com/wmentor/lemmas/engine/forms"
	"github.com/wmentor/lemmas/engine/opts"
)

type Word struct {
	Id    int64
	Opts  opts.Opts
	Forms []*forms.Form
}

func New(str string) *Word {

	w := &Word{
		Opts: opts.O_MASK,
	}

	for _, fl := range strings.Fields(str) {
		if f := forms.New(fl); f != nil {
			w.Forms = append(w.Forms, f)
			w.Opts = w.Opts & f.Opts
		}
	}

	if len(w.Forms) == 0 {
		return nil
	}

	return w
}

func (w *Word) String() string {
	maker := strings.Builder{}

	for i, f := range w.Forms {
		if i > 0 {
			maker.WriteRune(' ')
		}
		maker.WriteString(f.String())
	}

	return maker.String()
}
