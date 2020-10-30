package groups

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/wmentor/lemmas/lemma/flags"
	"github.com/wmentor/lemmas/lemma/forms"
	"github.com/wmentor/tokens"
)

var (
	padezh []flags.Flag = []flags.Flag{flags.F_IP, flags.F_RP, flags.F_DP, flags.F_VP, flags.F_TP, flags.F_PP}
)

func LoadFile(filename string) error {
	rh, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer rh.Close()

	Load(rh)
	return nil
}

func Load(in io.Reader) {
	forms.Reset()

	br := bufio.NewReader(in)

	var list []string = make([]string, 0, 16)

	ff := func(t string) {
		list = append(list, t)
	}

	for {

		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		list = list[:0]

		tokens.Process(strings.NewReader(str), ff)

		if len(list) < 2 {
			continue
		}

		f := flags.New(list[0])

		if f&flags.F_NOUN != 0 {
			list = list[1:]
			if len(list) == 6 {
				for i, form := range list {
					cf := padezh[i]
					forms.Add(form, (cf | f).ToIntStr(), list[0], (f | flags.F_IP).ToIntStr())
				}
			} else if len(list) == 1 {
				cf := f | flags.F_IP | flags.F_RP | flags.F_DP | flags.F_VP | flags.F_TP | flags.F_PP
				forms.Add(list[0], cf.ToIntStr(), list[0], cf.ToIntStr())
			}
		}
	}
}
