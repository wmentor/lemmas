package forms

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	_ "embed" //nolint
)

// Iteroator func for Each call
type EachFunc func(string) bool

//go:embed data.txt
var dataFile []byte

var (
	formData map[string]string
)

func init() {
	formData = make(map[string]string)
	loadData(bytes.NewReader(dataFile))
	dataFile = nil
}

func loadData(in io.Reader) {

	br := bufio.NewReader(in)

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		list := strings.Fields(strings.ToLower(str))

		for i, f := range list {
			pushPair(f, list[0])
			if i == 0 {
				pushPair("#"+list[0], list[0])
			}
		}
	}
}

func pushPair(src, dest string) {
	if res, has := formData[src]; has {
		list := strings.Split(res, " ")
		for _, c := range list {
			if c == dest {
				return
			}
		}
		list = append(list, dest)
		formData[src] = strings.Join(list, " ")
	} else {
		formData[src] = dest
	}
}

// Iterate over current and each known base forms
func Each(src string, fn EachFunc) {
	if fn(src) {
		if res, has := Get(src); has {
			for _, f := range res {
				if f != src && !fn(f) {
					return
				}
			}
		}
	}
}

// Get base forms for src string if exists
func Get(src string) ([]string, bool) {
	if res, has := formData[src]; has {
		return strings.Fields(res), true
	}
	return nil, false
}
