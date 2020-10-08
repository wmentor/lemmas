package save

import (
	"github.com/wmentor/lemmas"
	"github.com/wmentor/serv"
)

func init() {
	serv.Register("GET", "/forms/save", handler)
}

func handler(c *serv.Context) {
	lemmas.Save()
	c.WriteRedirect("/forms")
}
