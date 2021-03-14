package lemmas

import (
	"io"

	"github.com/wmentor/lemmas/stat"
)

// fetch result function type
type EachResultFunc = stat.EachResultFunc

// text processor interface
type Processor interface {
	AddText(in io.Reader)
	FetchResult(EachResultFunc)
}
