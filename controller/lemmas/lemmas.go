package lemmas

import (
	"strings"

	"github.com/wmentor/lemmas/lemma"
	"github.com/wmentor/serv"
)

func init() {
	serv.Register("GET", "/lemmas", page)
	serv.Register("POST", "/lemmas", proc)
}

func page(c *serv.Context) {
	c.SetContentType("text/html; charset=utf-8")
	c.WriteHeader(200)
	c.Render("lemmas_new.jet", nil)
}

func proc(c *serv.Context) {

	data := c.FormValue("text")

	maker := strings.Builder{}

	lemma.Process(strings.NewReader(data), func(l string) {
		maker.WriteString(strings.ReplaceAll(l, " ", "|"))
		maker.WriteRune(' ')
	})

	c.SetContentType("text/html; charset=utf-8")
	c.WriteHeader(200)
	c.Render("lemmas_proc.jet", map[string]interface{}{"data": maker.String()})
}
