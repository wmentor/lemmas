package opts

import (
	"strings"
)

type Opts uint64

const (
	O_RU Opts = 0x01 << iota
	O_EN

	O_NOUN // noun
	O_ADJ  // adj
	O_ADV  // adverb
	O_VERB // verb
	O_PRET // pretext
	O_PRON // pronoun
	O_PART // particle
	O_CONJ // conjuction
	O_NUM  // number

	O_SG // single
	O_MG // multi

	O_MR
	O_GR
	O_SR

	O_IP
	O_RP
	O_DP
	O_VP
	O_TP
	O_PP

	O_INF    // infinitive
	O_PAST   // past
	O_PRES   // present
	O_FUTURE // future
	O_IMP    // imperative

	O_F1 // face 1
	O_F2 // face 2
	O_F3 // face 3

	O_FN // first name
	O_SN // second name
	O_LN // last name

	O_ROMAN // roman number

	O_END // LAST BIT
)

var (
	encode map[string]Opts
	decode map[Opts]string
)

func init() {

	encode = map[string]Opts{}
	decode = map[Opts]string{}

	list := []string{"ru", "en",
		"noun", "adj", "adv", "verb", "pret", "pron", "part", "conj", "num", "sg", "mg", "mr", "gr", "sr",
		"ip", "rp", "dp", "vp", "tp", "pp",
		"inf", "past", "pres", "fut", "imp",
		"f1", "f2", "f3", "fn", "sn", "ln", "roman"}

	cur := Opts(0x01)

	for _, v := range list {
		encode[v] = cur
		decode[cur] = v
		cur = cur << 1
	}
}

func New(txt string) Opts {

	res := Opts(0)

	for {

		if idx := strings.Index(txt, "."); idx >= 0 {
			sub := txt[:idx]
			txt = txt[idx+1:]
			res = res | encode[sub]
		} else {
			res = res | encode[txt]
			break
		}

	}

	return res
}

func (f Opts) String() string {

	has := false
	maker := strings.Builder{}

	for cur := Opts(1); cur < O_END; cur = cur << 1 {
		if f&cur != 0 {
			if has {
				maker.WriteRune('.')
			}
			maker.WriteString(decode[cur])
			has = true
		}
	}

	return maker.String()
}
