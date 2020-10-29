package flags

import (
	"testing"
)

func TestFlags(t *testing.T) {

	tF := func(txt string, flag Flag) {
		f := New(txt)
		if f != flag {
			t.Fatalf("New failed for: %s", txt)
		}
		if flag.String() != txt {
			t.Fatalf("String failed for: %s", txt)
		}
	}

	tE := func(txt string) {
		f := New(txt)
		v := f.ToIntStr()
		f = FromIntStr(v)
		if f.String() != txt {
			t.Fatalf("ToIntStr/FromIntStr failed for: %s", txt)
		}
	}

	tF("", 0)
	tF("mr", F_MR)
	tF("noun.sg.mr.vp", F_NOUN|F_SG|F_MR|F_VP)
	tF("noun.mg.sr.pp", F_NOUN|F_MG|F_SR|F_PP)
	tF("noun.ln", F_NOUN|F_LNAME)
	tF("roman", F_ROMAN)

	tE("noun.sg.mr.vp")
	tE("noun.ln")
}
