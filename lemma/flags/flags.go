package flags

import (
	"strconv"
	"strings"
)

type Flag uint64

const (
	F_NOUN Flag = 0x01 << iota
	F_ADJ
	F_VERB
	F_PRET

	F_SG
	F_MG

	F_MR
	F_GR
	F_SR

	F_IP
	F_RP
	F_DP
	F_VP
	F_TP
	F_PP

	F_F1
	F_F2
	F_F3

	F_NAME
	F_SNAME
	F_LNAME

	F_PROF

	F_ROMAN

	F_END
)

var (
	encode map[string]Flag
	decode map[Flag]string
)

func init() {

	encode = map[string]Flag{}
	decode = map[Flag]string{}

	list := []string{"noun", "adj", "verb", "pret", "sg", "mg", "mr", "gr", "sr", "ip", "rp", "dp", "vp", "tp", "pp",
		"f1", "f2", "f3", "fn", "sn", "ln", "prof", "roman"}

	cur := Flag(0x01)

	for _, v := range list {
		encode[v] = cur
		decode[cur] = v
		cur = cur << 1
	}
}

func New(txt string) Flag {

	res := Flag(0)

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

func (f Flag) String() string {

	has := false
	maker := strings.Builder{}

	for cur := Flag(1); cur < F_END; cur = cur << 1 {
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

func (f Flag) ToIntStr() string {
	return strconv.FormatUint(uint64(f), 10)
}

func FromIntStr(str string) Flag {
	val, _ := strconv.ParseUint(str, 10, 64)
	return Flag(val)
}
