package storage

import (
	"testing"
)

func TestKeywords(t *testing.T) {

	KeywordAdd("пушкин")
	KeywordAdd("орел")

	if IsKeyword("12312csfasdfas") {
		t.Fatal("Unknown keyword")
	}

	if !IsKeyword("орел") {
		t.Fatal("Keyword not found")
	}
}
