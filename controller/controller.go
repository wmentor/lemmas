package controller

import (
	"net/http"

	"github.com/wmentor/lemmas/controller/generic"
	"github.com/wmentor/serv"

	_ "github.com/wmentor/lemmas/controller/api"
	_ "github.com/wmentor/lemmas/controller/cp"
)

func init() {
	serv.Register("GET", "/", handler)
}

func handler(c *serv.Context) {

	vars := generic.DefaultVars(c)

	c.SetContentType("text/html; charset=utf-8")
	c.WriteHeader(http.StatusOK)
	c.Render("cp/main.jet", vars)
}
