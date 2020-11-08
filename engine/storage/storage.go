package storage

import (
	"bytes"
	"strconv"
	"strings"
	"sync"

	"github.com/wmentor/kv"
	"github.com/wmentor/lemmas/engine/words"
)

var (
	mt sync.RWMutex

	idCnt []byte = []byte{0}
)

func Open(db string) error {

	if _, err := kv.Open("global=1 path=" + db); err != nil {
		return err
	}

	return nil
}

func Close() {
	kv.Close()
}

func nextId() int64 {
	mt.Lock()
	defer mt.Unlock()

	val := kv.Get(idCnt)

	id, _ := strconv.ParseInt(string(val), 10, 64)
	if id == 0 {
		id = 1
	}

	kv.Set(idCnt, []byte(strconv.FormatInt(id+1, 10)))

	return id
}

func formToKey(form string) []byte {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte(1)
	buf.WriteString(form)
	return buf.Bytes()
}

func wordIdToKey(wid int64) []byte {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte(2)
	buf.WriteString(strconv.FormatInt(wid, 10))
	return buf.Bytes()
}

func addForm(form string, wordId int64) {

	key := formToKey(form)
	wIdStr := strconv.FormatInt(wordId, 10)

	val := kv.Get(key)

	if len(val) > 0 {
		for _, ws := range strings.Fields(string(val)) {
			if ws == wIdStr {
				return
			}
		}

		buf := bytes.NewBuffer(nil)

		buf.Write(val)
		buf.WriteByte(' ')
		buf.WriteString(wIdStr)

		kv.Set(key, buf.Bytes())

	} else {
		kv.Set(key, []byte(wIdStr))
	}

}

func AddWord(wstr string) {
	if w := words.New(wstr); w != nil {
		wid := nextId()
		key := wordIdToKey(wid)
		kv.Set(key, w.Bytes())

		for _, f := range w.Forms {
			addForm(f.Name, wid)
		}
	}
}
