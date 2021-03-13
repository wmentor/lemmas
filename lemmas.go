package lemmas

import (
	"io"

	"github.com/wmentor/lemmas/stat"
)

// text process analyse result record
type Keyword = stat.Record //nolint

// text processor interface
type Processor interface {
	TextProc(in io.Reader) []Keyword
}
