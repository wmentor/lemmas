package engine

import (
	"os"

	"github.com/wmentor/lemmas/engine/storage"
	"github.com/wmentor/log"
)

var (
	dataDir   string
	formsFile string
)

func Open(dir string) {

	if dir == "" {
		if dir = os.Getenv("GOPATH"); dir != "" {
			dir = dir + "/src/github.com/wmentor/lemmas/data"
		} else {
			dir = "./data"
		}
	}

	dataDir = dir

	formsFile = dataDir + "/forms.txt"

	if rh, err := os.Open(formsFile); err == nil {
		defer rh.Close()
		storage.FormsLoad(rh)
	} else {
		log.Errorf("read file %s error: %s", formsFile, err.Error())
	}

}

func Save() {
	if wh, err := os.Create(formsFile); err == nil {
		defer wh.Close()
		storage.FormsSave(wh)
	} else {
		log.Errorf("write file %s error: %s", formsFile, err.Error())
	}
}
