package opts

import (
	"testing"
)

func TestOpts(t *testing.T) {

	tF := func(txt string, flag Opts) {
		f := New(txt)
		if f != flag {
			t.Fatalf("New failed for: %s", txt)
		}
		if flag.String() != txt {
			t.Fatalf("String failed for: %s", txt)
		}
	}

	tF("", 0)
	tF("mr", O_MR)
	tF("noun.sg.mr.vp", O_NOUN|O_SG|O_MR|O_VP)
	tF("noun.mg.sr.pp", O_NOUN|O_MG|O_SR|O_PP)
	tF("noun.ln", O_NOUN|O_LN)
	tF("roman", O_ROMAN)
}
