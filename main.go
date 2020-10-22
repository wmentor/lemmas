package main

import (
	"flag"

	"github.com/wmentor/lemmas/lemma"
	"github.com/wmentor/serv"

	_ "github.com/wmentor/lemmas/controller"
)

func main() {

	var addr string

	flag.StringVar(&addr, "listen", ":8080", "listen address")

	flag.Parse()

	lemma.Open("")

	serv.LoadTemplates("./templates")

	if err := serv.Start(addr); err != nil {
		panic(err)
	}
}
