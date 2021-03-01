// +build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func loadData() []string {
	rh, err := os.Open("data.txt")
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

func saveData(list []string) {
	wh, err := os.Create("data.txt")
	if err != nil {
		panic(err)
	}
	defer wh.Close()

	for _, str := range list {
		fmt.Fprintln(wh, str)
	}
}

func main() {
	list := loadData()
	if len(list) > 0 {
		saveData(list)
	}
}
