package main

import (
	"flag"

	"github.com/wmentor/lemmas/engine"
	"github.com/wmentor/log"
	"github.com/wmentor/serv"

	_ "github.com/wmentor/lemmas/controller"
)

func main() {

	var addr string
	var dataDir string

	flag.StringVar(&addr, "listen", ":8080", "listen address")
	flag.StringVar(&dataDir, "data", "./data", "data directory")

	flag.Parse()

	log.Open("name=lemmas path=./log period=day level=info keep=15 stderr=1")

	engine.Open(dataDir)

	serv.SetLogger(func(l *serv.LogData) {
		log.Infof("%s %s %d %s %.3fs", l.Addr, l.Method, l.StatusCode, l.RequestURL, l.Seconds)
	})

	serv.LoadTemplates("./templates")

	serv.Static("/fonts", "./htdocs/fonts")
	serv.Static("/css", "./htdocs/css")

	if err := serv.Start(addr); err != nil {
		panic(err)
	}
}
