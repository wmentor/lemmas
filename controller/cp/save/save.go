package save

import (
	"net/http"

	"github.com/wmentor/lemmas/engine"
	"github.com/wmentor/serv"
)

func init() {
	serv.Register("GET", "/cp/save", handler)
}

func handler(c *serv.Context) {

	go engine.Save()

	c.SetContentType("text/plain; charset=utf-8")
	c.WriteHeader(http.StatusOK)
	c.WriteString("OK")
}
