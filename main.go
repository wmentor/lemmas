package main

import (
	"flag"

	"github.com/wmentor/lemmas/lemma"
	"github.com/wmentor/log"
	"github.com/wmentor/serv"

	_ "github.com/wmentor/lemmas/controller"
)

func main() {

	var addr string
	var lemmasData string

	flag.StringVar(&addr, "listen", ":8080", "listen address")
	flag.StringVar(&lemmasData, "lemmas", "./data/lemmas.db", "data likes lemmas.db file")

	flag.Parse()

	log.Open("name=lemmas path=./log period=day level=info keep=15 stderr=1")

	lemma.Open(lemmasData)

	serv.SetLogger(func(l *serv.LogData) {
		log.Infof("%s %s %d %s %.3fs", l.Addr, l.Method, l.StatusCode, l.RequestURL, l.Seconds)
	})

	serv.LoadTemplates("./templates")

	if err := serv.Start(addr); err != nil {
		panic(err)
	}
}
