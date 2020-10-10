package forms

import (
	"strings"

	"github.com/wmentor/lemmas"
	"github.com/wmentor/serv"

	_ "github.com/wmentor/lemmas/server/controller/forms/raw"
	_ "github.com/wmentor/lemmas/server/controller/forms/save"
)

func init() {
	serv.Register("GET", "/forms", page)
	serv.Register("POST", "/forms", add)
}

func page(c *serv.Context) {
	c.SetContentType("text/html; charset=utf-8")
	c.WriteHeader(200)
	c.Render("forms.jet", nil)
}

func add(c *serv.Context) {

	f := strings.ToLower(strings.TrimSpace(c.FormValue("form")))
	b := strings.ToLower(strings.TrimSpace(c.FormValue("base")))

	if f != "" && b != "" {
		lemmas.AddForm(f, b)
	}

	page(c)
}
