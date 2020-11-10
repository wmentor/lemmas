package storage

import (
	"bytes"
	"strconv"
	"strings"
	"sync"

	"github.com/wmentor/kv"
	"github.com/wmentor/lemmas/engine/words"
)

type FindResult struct {
	Form  string
	Words []*words.Word
}

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

func TestOpen() {
	kv.Open("test=1")
}

func Close() {
	kv.Close()
}

func nextId() int64 {

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

func delForm(form string, wordId int64) {

	key := formToKey(form)
	wIdStr := strconv.FormatInt(wordId, 10)

	if val := kv.Get(key); len(val) > 0 {

		buf := bytes.NewBuffer(nil)
		j := 0

		for _, ids := range strings.Fields(string(val)) {
			if ids != wIdStr {
				if j > 0 {
					buf.WriteRune(' ')
				}
				buf.WriteString(ids)
				j++
			}
		}

		if j > 0 {
			kv.Set(key, buf.Bytes())
		} else {
			kv.Set(key, nil)
		}
	}
}

func AddWord(wstr string) int64 {
	if w := words.New(wstr); w != nil {

		mt.Lock()
		defer mt.Unlock()

		wid := nextId()

		key := wordIdToKey(wid)
		kv.Set(key, w.Bytes())

		for _, f := range w.Forms {
			addForm(f.Name, wid)
		}

		return wid
	}

	return 0
}

func getWord(wid int64) *words.Word {
	if data := kv.Get(wordIdToKey(wid)); len(data) > 0 {
		return words.New(string(data))
	}
	return nil
}

func DelWord(wid int64) {
	mt.Lock()
	defer mt.Unlock()

	if w := getWord(wid); w != nil {
		kv.Set(wordIdToKey(wid), nil)

		for _, f := range w.Forms {
			delForm(f.Name, wid)
		}
	}
}

func Find(form string) *FindResult {

	mt.RLock()
	defer mt.RUnlock()

	fr := &FindResult{Form: form}
	key := formToKey(form)
	data := kv.Get(key)
	if len(data) > 0 {
		for _, id := range strings.Fields(string(data)) {
			if wid, _ := strconv.ParseInt(id, 10, 64); wid > 0 {
				if w := getWord(wid); w != nil {
					fr.Words = append(fr.Words, w)
				}
			}
		}
	}
	return fr
}
