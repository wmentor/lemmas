package stat

import (
	"net/http"

	"github.com/wmentor/lemmas/controller/generic"
	"github.com/wmentor/serv"
)

func init() {
	serv.Register("GET", "/cp/stat", handler)
}

func handler(c *serv.Context) {

	vars := generic.DefaultVars(c)

	c.SetContentType("text/html; charset=utf-8")
	c.WriteHeader(http.StatusOK)
	c.Render("cp/words/stat.jet", vars)
}
