// +build ignore

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	preAlloc int = 10240
)

var (
	shiftVal int = 0
)

func main() {

	var filename string
	flag.StringVar(&filename, "f", "", "process file name")
	flag.IntVar(&shiftVal, "shift", 0, "shift window")
	flag.Parse()

	if filename == "" {
		fmt.Println("use: generator -f filename")
		os.Exit(1)
	}

	list := loadFile(filename)
	saveFile(filename, list)
}

func pushPair(hash map[string]string, key string, value string) {
	if res, has := hash[key]; has {
		list := strings.Fields(res)
		if len(list) >= shiftVal {
			for _, c := range list[shiftVal:] {
				if c == value {
					return
				}
			}
		}
		list = append(list, value)
		hash[key] = strings.Join(list, " ")
	} else {
		hash[key] = value
	}
}

func loadFile(filename string) []string {
	rh, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open file %s error %v\n", filename, err)
		os.Exit(1)
	}
	defer rh.Close()

	br := bufio.NewReader(rh)

	hash := make(map[string]string, preAlloc)

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}
		list := strings.Fields(strings.ToLower(str))
		for _, v := range list {
			pushPair(hash, list[0], v)
		}
	}
	if len(hash) == 0 {
		return nil
	}

	list := make([]string, 0, len(hash))

	for _, v := range hash {
		list = append(list, v)
	}

	sort.Strings(list)
	return list
}

func saveFile(filename string, list []string) {
	wh, err := os.Create(filename)
	if err != nil {
		fmt.Printf("rewrite file %s error %v\n", filename, err)
		os.Exit(1)
	}
	defer wh.Close()

	for _, str := range list {
		fmt.Fprintln(wh, str)
	}
}
