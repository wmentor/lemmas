package unknown

import (
	"net/http"
	"strings"

	"github.com/wmentor/lemmas/controller/generic"
	"github.com/wmentor/lemmas/engine/forms"
	"github.com/wmentor/serv"
	"github.com/wmentor/tokens"
)

func init() {
	serv.Register("GET", "/cp/words/unknown", handler)
	serv.Register("POST", "/cp/words/unknown", handler)
}

func handler(c *serv.Context) {

	vars := generic.DefaultVars(c)

	if c.Method() == "POST" {

		maker := strings.Builder{}

		tokens.Process(strings.NewReader(c.FormValue("data")), func(t string) {
			if forms.Has(t) {
				maker.WriteString(`<span class="found">`)
				maker.WriteString(t)
				maker.WriteString("</span> ")
			} else {
				maker.WriteString(`<span class="unknown">`)
				maker.WriteString(t)
				maker.WriteString("</span> ")
			}
		})

		c.SetContentType("text/html; charset=utf-8")
		c.WriteHeader(http.StatusOK)
		vars["data"] = maker.String()
		c.Render("cp/words/unknown_result.jet", vars)

		return
	}

	c.SetContentType("text/html; charset=utf-8")
	c.WriteHeader(http.StatusOK)
	c.Render("cp/words/unknown_check.jet", vars)
}
