package dicts

import (
	"bufio"
	"bytes"
	_ "embed" //nolint
	"strconv"
	"strings"
)

//go:embed dict_names.txt
var dictNames []byte

//go:embed dict_lastnames.txt
var dictLastnames []byte

func loader() {
	allDicts["names"] = dictFromBytes(dictNames)
	dictNames = nil

	allDicts["lastnames"] = dictFromBytes(dictLastnames)
	dictLastnames = nil
}

func dictFromBytes(in []byte) Dict {
	br := bufio.NewReader(bytes.NewReader(in))
	rd := dict(map[string]int{})

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		list := strings.Fields(strings.ToLower(str))
		if len(list) != 2 {
			continue
		}

		code, _ := strconv.Atoi(list[1])
		rd[list[0]] = code
	}

	return rd
}
