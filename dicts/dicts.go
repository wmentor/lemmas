package dicts

import (
	"errors"
)

//go:generate go run generator.go

type Dict interface {
	Get(word string) int
}

var (
	ErrDictNotFound error = errors.New("dictionary not found")

	allDicts map[string]Dict
)

func init() {
	allDicts = map[string]Dict{}
	loader()
}

type dict map[string]int

func (d dict) Get(word string) int {
	return d[word]
}

func GetDict(name string) (Dict, error) {
	if d, has := allDicts[name]; has {
		return d, nil
	}
	return nil, ErrDictNotFound
}
