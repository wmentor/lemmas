package storage

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/wmentor/log"
)

var (
	mt sync.RWMutex

	nextId   int64
	formsIdx map[string]string
	wordsIdx map[string]string

	dir string
)

func init() {

	dir = "."

	nextId = 1
	formsIdx = make(map[string]string)
	wordsIdx = make(map[string]string)

}

func Open(db string) {

	dir = db

}

func loadNextId(dir string) int64 {

	filename := dir + "/nextid.txt"

	if data, err := ioutil.ReadFile(filename); err != nil {
		log.Errorf("load %s failed: %s", filename, err.Error())
		return 1
	} else {
		if nid, _ := strconv.ParseInt(string(data), 16, 64); nid > 0 {
			return nid
		}
	}

	return 1
}

func loadForms(dir string) map[string]string {

	res := make(map[string]string)

	filename := dir + "/forms.txt"

	if rh, err := os.Open(filename); err == nil {
		defer rh.Close()

		br := bufio.NewReader(rh)

		for {
			str, err := br.ReadString('\n')
			if err != nil && str == "" {
				break
			}

			if str = strings.TrimSpace(str); len(str) > 0 {
				if idx := strings.IndexRune(str, '|'); idx > 0 {
					res[str[:idx]] = str[idx+1:]
				}
			}
		}

	} else {
		log.Errorf("load %s failed: %s", filename, err.Error())
	}

	return res
}

func loadWords(dir string) map[string]string {
	res := make(map[string]string)

	filename := dir + "/words.txt"

	if rh, err := os.Open(filename); err == nil {
		defer rh.Close()

		br := bufio.NewReader(rh)

		for {
			str, err := br.ReadString('\n')
			if err != nil && str == "" {
				break
			}

			if str = strings.TrimSpace(str); len(str) > 0 {
				if idx := strings.IndexRune(str, '|'); idx > 0 {
					res[str[:idx]] = str[idx+1:]
				}
			}
		}

	} else {
		log.Errorf("load %s failed: %s", filename, err.Error())
	}

	return res
}
