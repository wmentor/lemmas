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

func main() {

	var filename string
	flag.StringVar(&filename, "f", "", "process file name")
	flag.Parse()

	if filename == "" {
		fmt.Println("use: generator -f filename")
		os.Exit(1)
	}

	list := loadFile(filename)
	saveFile(filename, list)
}

func loadFile(filename string) []string {
	rh, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open file %s error %v\n", filename, err)
		os.Exit(1)
	}
	defer rh.Close()

	br := bufio.NewReader(rh)

	hash := make(map[string]bool, preAlloc)

	for {
		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}
		str = strings.Join(strings.Fields(strings.ToLower(str)), " ")
		if str == "" {
			continue
		}
		hash[str] = true
	}
	if len(hash) == 0 {
		return nil
	}

	list := make([]string, 0, len(hash))

	for k := range hash {
		list = append(list, k)
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
