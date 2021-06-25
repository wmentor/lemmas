package dicts

import (
	"bufio"
	"embed"
	"strings"
)

// Iterator func for Each call
type EachFunc func(string) bool

var (
	dataset map[string]string
)

//go:embed dict_*.txt
var dictsFS embed.FS

func init() {
	dataset = make(map[string]string)

	flist, err := dictsFS.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, rec := range flist {
		loadDict(rec.Name())
	}
}

func loadDict(filename string) {

	size := len(filename)
	dict := strings.ReplaceAll(filename[5:size-4], "_", "")

	in, err := dictsFS.Open(filename)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	br := bufio.NewReader(in)
	value := ":" + dict

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		str = strings.ToLower(strings.TrimSpace(str))
		if str != "" {
			pushPair(str, value)
		}
	}
}

func pushPair(src, dest string) {
	if res, has := dataset[src]; has {
		list := strings.Split(res, " ")
		for _, c := range list {
			if c == dest {
				return
			}
		}
		list = append(list, dest)
		dataset[src] = strings.Join(list, " ")
	} else {
		dataset[src] = dest
	}
}

func InDict(dict string, name string) bool {
	search := ":" + dict
	if data, has := dataset[name]; has {
		for _, v := range strings.Fields(data) {
			if v == search {
				return true
			}
		}
	}
	return false
}

// Iterate over each word dict
func Each(src string, fn EachFunc) {
	if res, has := Get(src); has {
		for _, f := range res {
			if f != src && !fn(f) {
				return
			}
		}
	}
}

// Get dict names for src string if exists
func Get(src string) ([]string, bool) {
	if res, has := dataset[src]; has {
		return strings.Fields(res), true
	}
	return nil, false
}

// Get all dicts size
func Size() int {
	return len(dataset)
}
