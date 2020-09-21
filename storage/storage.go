package storage

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/wmentor/lemmas/log"
	"github.com/wmentor/tokens"
)

var (
	data map[string]string = map[string]string{}
)

func Load(in io.Reader) {

	br := bufio.NewReader(in)
	res := map[string]string{}

	var list []string = make([]string, 0, 10)

	fn := func(t string) {
		list = append(list, t)
	}

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		list = list[:0]

		tokens.Process(strings.NewReader(str), fn)

		if len(list) > 1 {
			res[list[0]] = strings.Join(list[1:], " ")
		} else if len(list) == 1 {
			log.Log("ERROR", "invalid string: %s", strings.TrimSpace(str))
		}
	}

	data = res
}

func LoadFile(filename string) error {
	if rh, e := os.Open(filename); e == nil {
		defer rh.Close()
		Load(rh)
	} else {
		return e
	}
	return nil
}

func Save(out io.Writer) {
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	list := make([]string, 0, len(data))

	for k := range data {
		list = append(list, k)
	}

	sort.Sort(sort.StringSlice(list))

	for i, k := range list {
		v := data[k]

		bw.WriteString(k)
		bw.WriteRune(' ')
		bw.WriteString(v)
		bw.WriteRune('\n')

		if (i+1)%100 == 0 {
			bw.Flush()
		}
	}
}

func SaveFile(filename string) error {
	if wh, err := os.Create(filename); err == nil {
		defer wh.Close()
		Save(wh)
	} else {
		return err
	}
	return nil
}
