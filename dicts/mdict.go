package dicts

type Dict interface {
	Has(string) bool
}

type mdict map[string]bool

func (m mdict) Has(str string) bool {
	return m[str]
}
