package fixed

import (
	"net/http"
	"strings"

	"github.com/wmentor/lemmas/controller/generic"
	"github.com/wmentor/lemmas/engine"
	"github.com/wmentor/serv"
)

func init() {
	serv.Register("GET", "/cp/words/load/fixed", handler)
	serv.Register("POST", "/cp/words/load/fixed", handler)
}

func handler(c *serv.Context) {

	if c.Method() == "POST" {
		engine.LoadFixed(strings.NewReader(strings.ReplaceAll(c.FormValue("words"), ",", " ")))
	}

	vars := generic.DefaultVars(c)

	c.SetContentType("text/html; charset=utf-8")
	c.WriteHeader(http.StatusOK)
	c.Render("cp/words/load_fixed.jet", vars)
}
