package text

import (
	"strings"

	"github.com/wmentor/lemmas"
	"github.com/wmentor/serv"
	"github.com/wmentor/tokens"
)

func init() {
	serv.Register("GET", "/text", page)
	serv.Register("POST", "/text", proc)
}

func page(c *serv.Context) {
	c.SetContentType("text/html; charset=utf-8")
	c.WriteHeader(200)
	c.Render("text_new.jet", nil)
}

func proc(c *serv.Context) {

	data := c.FormValue("text")

	maker := strings.Builder{}

	tokens.Process(strings.NewReader(data), func(t string) {
		if lemmas.CanProcess(t) {
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
	c.Render("text_proc.jet", map[string]interface{}{"data": maker.String()})
}
