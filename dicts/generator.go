// +build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func loadData(filename string) []string {
	rh, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer rh.Close()

	br := bufio.NewReader(rh)

	var list []string

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}
		str = strings.Join(strings.Fields(strings.ToLower(str)), " ")
		if str == "" {
			continue
		}
		list = append(list, str)
	}
	if len(list) == 0 {
		return nil
	}

	sort.Strings(list)
	return list
}

func saveData(filename string, list []string) {
	wh, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer wh.Close()

	for _, str := range list {
		fmt.Fprintln(wh, str)
	}
}

func main() {

	for _, filename := range []string{"dict_lastnames.txt", "dict_names.txt"} {
		list := loadData(filename)
		if len(list) > 0 {
			saveData(filename, list)
		}
	}
}
