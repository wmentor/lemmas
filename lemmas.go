package lemmas

import (
	"io"

	"github.com/wmentor/lemmas/stat"
)

// fetch result function type
type EachResultFunc = stat.EachResultFunc

// text processor interface
type Processor interface {
	AddText(in io.Reader)       // add text data to process
	AddHTML(in io.Reader)       // add text data from HTML
	FetchResult(EachResultFunc) // return result callback
	Tokens() int64              // tokens counter
	Reset()                     // reset data and reinit object
	ReadingTime() int64         // calc reading time in seconds
}
