package storage

import (
	"strings"
	"testing"
)

func TestStorage(t *testing.T) {

	data := `
хранилище хранилище
хранилища хранилища хранилище
хранилищу хранилище
хранилищем хранилище
хранилищ хранилища
хранилищам хранилища
хранилищами хранилища
хранилищах хранилища
  `

	Load(strings.NewReader(data))

	tBI := func(form string, wait string) {
		var list []string
		EachBase(form, func(t string) bool {
			list = append(list, t)
			return true
		})
		res := strings.Join(list, " ")
		if res != wait {
			t.Fatalf("Invalid bases for: %s", form)
		}
	}

	tBI("хранилище", "хранилище")
	tBI("хранилища", "хранилища хранилище")
	tBI("хранилищ", "хранилища")
	tBI("12", "12")

	Set("хранилища")

	tBI("хранилища", "")

	Set("хранилища", "хранилища", "хранилище")

	tBI("хранилища", "хранилища хранилище")
}
