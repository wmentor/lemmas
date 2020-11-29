package engine

import (
	"bufio"
	"compress/gzip"
	"os"
	"time"

	"github.com/wmentor/lemmas/engine/forms"
	"github.com/wmentor/lemmas/engine/meta"
	"github.com/wmentor/log"
)

var (
	dataDir   string
	formsFile string
	metaFile  string
	needSave  bool
)

func init() {
	go monitor()
}

func Open(dir string) {

	if dir == "" {
		if dir = os.Getenv("GOPATH"); dir != "" {
			dir = dir + "/src/github.com/wmentor/lemmas/data"
		} else {
			dir = "./data"
		}
	}

	dataDir = dir

	formsFile = dataDir + "/forms.txt.gz"
	metaFile = dataDir + "/meta.txt.gz"

	writeAccess(func() {

		if rh, err := os.Open(formsFile); err == nil {
			defer rh.Close()

			br := bufio.NewReader(rh)

			if gz, err := gzip.NewReader(br); err == nil {
				forms.Reset()
				forms.Load(gz)
			} else {
				log.Errorf("read file %s error: %s", formsFile, err.Error())
			}

		} else {
			log.Errorf("read file %s error: %s", formsFile, err.Error())
		}

		if rh, err := os.Open(metaFile); err == nil {
			defer rh.Close()

			br := bufio.NewReader(rh)

			if gz, err := gzip.NewReader(br); err == nil {
				meta.Reset()
				meta.Load(gz)
			} else {
				log.Errorf("read file %s error: %s", metaFile, err.Error())
			}
		} else {
			log.Errorf("read file %s error: %s", metaFile, err.Error())
		}

	})
}

func Save() {

	readAccess(func() {

		if wh, err := os.Create(formsFile); err == nil {
			defer wh.Close()
			forms.Save(wh)
		} else {
			log.Errorf("write file %s error: %s", formsFile, err.Error())
		}

		if wh, err := os.Create(metaFile); err == nil {
			defer wh.Close()
			meta.Save(wh)
		} else {
			log.Errorf("write file %s error: %s", metaFile, err.Error())
		}

	})
}

func monitor() {

	for {
		time.Sleep(time.Minute)
		if needSave {
			needSave = false
			Save()
			log.Info("save data")
		}
	}

}
