package raw

import (
	"bufio"
	"strings"

	"github.com/wmentor/lemmas"
	"github.com/wmentor/serv"
	"github.com/wmentor/tokens"
)

func init() {
	serv.Register("GET", "/forms/raw", page)
	serv.Register("POST", "/forms/raw", save)
}

func page(c *serv.Context) {
	c.SetContentType("text/html; charset=utf-8")
	c.WriteHeader(200)
	c.Render("forms_raw.jet", nil)
}

func save(c *serv.Context) {

	data := c.FormValue("data")

	br := bufio.NewReader(strings.NewReader(data))

	var list []string

	fn := func(w string) {
		list = append(list, w)
	}

	for {

		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		list = list[:0]

		tokens.Process(strings.NewReader(str), fn)

		if len(list) > 1 {
			lemmas.AddForm(list[0], (list[1:])...)
		}
	}

	page(c)
}
