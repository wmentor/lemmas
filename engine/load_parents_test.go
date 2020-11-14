package engine

import (
	"strings"
	"testing"
)

func TestLoadParents(t *testing.T) {

	data := `
пушкин > поэт > человек
покер > азертная игра > игра
`

	LoadWords(strings.NewReader(data))

}
