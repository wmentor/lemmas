package engine

import (
	"strings"
	"testing"
)

func TestLoadParents(t *testing.T) {

	data := `
пушкин
покер
игра
`

	LoadKeywords(strings.NewReader(data))

}
