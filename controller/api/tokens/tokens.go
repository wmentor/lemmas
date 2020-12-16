package tokens

import (
	"bytes"
	"io"

	"github.com/wmentor/html"
	"github.com/wmentor/serv"
	"github.com/wmentor/tokens"
)

func init() {
	serv.Register("POST", "/api/tokens", handler)
}

func handler(c *serv.Context) {

	result := make([]string, 0, 1024)

	var in io.Reader

	if c.HasQueryParam("html") {
		parser := html.New()
		parser.Parse(c.Body())
		in = bytes.NewReader(parser.Text())
	} else {
		in = c.Body()
	}

	tokens.Process(in, func(t string) {
		result = append(result, t)
	})

	c.SetContentType("application/json; charset=utf-8")
	c.WriteHeader(200)
	c.WriteJson(result)
}
