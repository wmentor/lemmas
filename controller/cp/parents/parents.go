package parents

import (
	"net/http"
	"strings"

	"github.com/wmentor/lemmas/engine"
	"github.com/wmentor/serv"
)

func init() {
	serv.Register("GET", "/cp/parents", handler)
	serv.Register("POST", "/cp/parents", handler)
}

func handler(c *serv.Context) {

	if c.Method() == "POST" {
		engine.LoadParents(strings.NewReader(c.FormValue("data")))
	}

	c.SetContentType("text/html; charset=utf-8")
	c.WriteHeader(http.StatusOK)
	c.Render("cp/parents.jet", nil)
}
