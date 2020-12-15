package tokens

import (
	"github.com/wmentor/serv"
	"github.com/wmentor/tokens"
)

func init() {
	serv.Register("POST", "/api/tokens", handler)
}

func handler(c *serv.Context) {

	result := make([]string, 0, 1024)

	tokens.Process(c.Body(), func(t string) {
		result = append(result, t)
	})

	c.SetContentType("application/json; charset=utf-8")
	c.WriteHeader(200)
	c.WriteJson(result)
}
