package engine

import (
	"bufio"
	"os"
	"time"

	"github.com/wmentor/lemmas/engine/forms"
	"github.com/wmentor/lemmas/engine/meta"
	"github.com/wmentor/log"
)

var (
	dataDir   string
	formsFile string
	fixedFile string
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

	formsFile = dataDir + "/forms.txt"
	fixedFile = dataDir + "/fixed.txt"
	metaFile = dataDir + "/meta.txt"

	writeAccess(func() {

		forms.Reset()
		meta.Reset()

		if rh, err := os.Open(formsFile); err == nil {
			defer rh.Close()
			log.Infof("read %s", formsFile)

			br := bufio.NewReader(rh)

			forms.LoadForms(br)

		} else {
			log.Errorf("read file %s error: %s", formsFile, err.Error())
		}

		if rh, err := os.Open(fixedFile); err == nil {
			defer rh.Close()
			log.Infof("read %s", fixedFile)

			br := bufio.NewReader(rh)

			forms.LoadFixed(br)

		} else {
			log.Errorf("read file %s error: %s", formsFile, err.Error())
		}

		if rh, err := os.Open(metaFile); err == nil {
			defer rh.Close()
			log.Infof("read %s", metaFile)

			br := bufio.NewReader(rh)

			meta.Load(br)

		} else {
			log.Errorf("read file %s error: %s", metaFile, err.Error())
		}

	})
}

func Save() {

	readAccess(func() {

		if wh, err := os.Create(formsFile); err == nil {
			defer wh.Close()
			forms.SaveForms(wh)

		} else {
			log.Errorf("write file %s error: %s", formsFile, err.Error())
		}

		if wh, err := os.Create(fixedFile); err == nil {
			defer wh.Close()
			forms.SaveFixed(wh)

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
