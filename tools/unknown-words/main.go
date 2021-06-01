package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/wmentor/html"
	"github.com/wmentor/lemmas/counter"
	"github.com/wmentor/lemmas/forms"
	"github.com/wmentor/tokens"
)

func main() {

	var source string
	flag.StringVar(&source, "source", "", "source file or url")
	flag.Parse()

	if source == "" {
		fmt.Println("use: unknown-words -source [filename|url]")
		os.Exit(1)
	}

	var input io.Reader

	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		parser := html.New()

		opts := &html.GetOpts{
			Agent:   "Lemmas",
			Timeout: time.Second * 10,
		}

		if err := parser.Get(source, opts); err != nil {
			panic(err)
		}

		input = bytes.NewReader(parser.Text())

	} else {

		rh, err := os.Open(source)
		if err != nil {
			panic(err)
		}
		defer rh.Close()

		input = rh
	}

	skip := map[string]bool{
		"{":  true,
		"}":  true,
		"«":  true,
		"»":  true,
		"\"": true,
		"'":  true,
		".":  true,
		",":  true,
		":":  true,
		"?":  true,
		"!":  true,
		"%":  true,
		"$":  true,
		"-":  true,
		"(":  true,
		")":  true,
		"*":  true,
		"/":  true,
		";":  true,
		"[":  true,
		"]":  true,
	}

	cnt := counter.New()

	tokens.Process(input, func(t string) {

		if _, has := forms.Get(t); !has && !skip[t] && !strings.ContainsAny(t, "./:_#1234567890%") {
			cnt.Inc(t)
		}

	})

	i := 0

	cnt.EachFreq(func(key string, val int64) {
		i++
		if i < 500 {
			fmt.Println(key, val)
		}
	})
}
