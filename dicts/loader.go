package dicts

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	_ "embed" //nolint
)

var dictionaries map[string]Dict

//go:embed dict_m_names.txt
var dictMNames []byte

//go:embed dict_w_names.txt
var dictWNames []byte

//go:embed dict_countries.txt
var dictCountries []byte

//go:embed dict_roman.txt
var dictRoman []byte

//go:embed dict_w_lastnames.txt
var dictWLastnames []byte

//go:embed dict_m_lastnames.txt
var dictMLastnames []byte

//go:embed dict_cities.txt
var dictCities []byte

//go:embed dict_companies.txt
var dictCompanies []byte

func init() {
	dictionaries = make(map[string]Dict)

	dictionaries["cities"] = loadDict(bytes.NewReader(dictCities))
	dictionaries["companies"] = loadDict(bytes.NewReader(dictCompanies))
	dictionaries["countries"] = loadDict(bytes.NewReader(dictCountries))
	dictionaries["m_lastnames"] = loadDict(bytes.NewReader(dictMLastnames))
	dictionaries["m_names"] = loadDict(bytes.NewReader(dictMNames))
	dictionaries["roman"] = loadDict(bytes.NewReader(dictRoman))
	dictionaries["w_lastnames"] = loadDict(bytes.NewReader(dictWLastnames))
	dictionaries["w_names"] = loadDict(bytes.NewReader(dictWNames))

	dictMNames = nil
	dictWNames = nil
	dictCities = nil
	dictCountries = nil
	dictRoman = nil
}

func loadDict(in io.Reader) Dict {
	br := bufio.NewReader(in)

	res := mdict(make(map[string]bool))

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		str = strings.ToLower(strings.TrimSpace(str))
		if str != "" {
			res[str] = true
		}
	}

	return res
}

func InDict(dict string, name string) bool {
	if dict, has := dictionaries[dict]; has {
		return dict.Has(name)
	}
	return false
}
