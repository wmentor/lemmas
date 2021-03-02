package keywords

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	_ "embed" //nolint
)

//go:generate go run generator.go

//go:embed data.txt
var dataFile []byte

var (
	data map[string]string
)

func init() {
	data = make(map[string]string)
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

		if list := strings.Fields(strings.ToLower(str)); len(list) > 1 {
			for _, f := range list[1:] {
				pushPair(list[0], f)
			}
		}
	}
}

func pushPair(src, dest string) {
	if res, has := data[src]; has {
		list := strings.Split(res, " ")
		for _, c := range list {
			if c == dest {
				return
			}
		}
		list = append(list, dest)
		data[src] = strings.Join(list, " ")
	} else {
		data[src] = dest
	}
}

// Get all keyword for source data
func Get(src string) ([]string, bool) {
	if res, has := data[src]; has {
		return strings.Fields(res), true
	}
	return nil, false
}
