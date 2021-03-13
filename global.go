package lemmas

const (
	bufferSize int = 5
)

var (
	eos map[string]bool = map[string]bool{".": true, "?": true, "!": true, "…": true}
)
