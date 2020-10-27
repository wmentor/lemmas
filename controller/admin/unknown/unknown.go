package unknown

import (
	"strings"

	"github.com/wmentor/lemmas/lemma"
	"github.com/wmentor/serv"
	"github.com/wmentor/tokens"
)

func init() {
	serv.Register("GET", "/admin/unknown", page)
	serv.Register("POST", "/admin/unknown", proc)
}

func page(c *serv.Context) {
	c.SetContentType("text/html; charset=utf-8")
	c.WriteHeader(200)
	c.Render("admin/unknown_form.jet", nil)
}

func proc(c *serv.Context) {

	data := c.FormValue("text")

	maker := strings.Builder{}

	tokens.Process(strings.NewReader(data), func(t string) {
		if lemma.CanProcess(t) {
			maker.WriteString("<span class=\"oldtok\">")
			maker.WriteString(t)
			maker.WriteString("</span> ")
		} else {
			maker.WriteString("<span class=\"newtok\">")
			maker.WriteString(t)
			maker.WriteString("</span> ")
		}
	})

	c.SetContentType("text/html; charset=utf-8")
	c.WriteHeader(200)
	c.Render("admin/unknown_res.jet", map[string]interface{}{"data": maker.String()})
}
