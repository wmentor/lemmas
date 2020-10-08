package main

import (
	"flag"

	"github.com/wmentor/lemmas"
	"github.com/wmentor/serv"

	_ "github.com/wmentor/lemmas/server/controller"
)

func main() {

	var addr string

	flag.StringVar(&addr, "listen", ":8080", "listen address")

	flag.Parse()

	lemmas.Open("")

	serv.LoadTemplates("./templates")

	if err := serv.Start(addr); err != nil {
		panic(err)
	}
}
