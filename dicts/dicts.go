package dicts

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	_ "embed" //nolint
)

// Iterator func for Each call
type EachFunc func(string) bool

var (
	dataset map[string]string
)

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

//go:embed dict_m_patronymics.txt
var dictMPatronymics []byte

//go:embed dict_w_patronymics.txt
var dictWPatronymics []byte

func init() {
	dataset = make(map[string]string)

	loadEmbed("cities", bytes.NewReader(dictCities))
	dictCities = nil

	loadEmbed("companies", bytes.NewReader(dictCompanies))
	dictCompanies = nil

	loadEmbed("countries", bytes.NewReader(dictCountries))
	dictCountries = nil

	loadEmbed("mlastnames", bytes.NewReader(dictMLastnames))
	dictMLastnames = nil

	loadEmbed("mnames", bytes.NewReader(dictMNames))
	dictMNames = nil

	loadEmbed("mpatronymics", bytes.NewReader(dictMPatronymics))
	dictMPatronymics = nil

	loadEmbed("roman", bytes.NewReader(dictRoman))
	dictRoman = nil

	loadEmbed("wlastnames", bytes.NewReader(dictWLastnames))
	dictWLastnames = nil

	loadEmbed("wnames", bytes.NewReader(dictWNames))
	dictWNames = nil

	loadEmbed("wpatronymics", bytes.NewReader(dictWPatronymics))
	dictWPatronymics = nil

}

func loadEmbed(dict string, in io.Reader) {
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
