package counter

import (
	"testing"
)

func TestWeight(t *testing.T) {

	tF := func(x int64, downBorder float64, upBorder float64) {
		res := weight(x)
		if res < downBorder || res > upBorder {
			t.Fatalf("weight(%v) failed", x)
		}
	}

	tF(0, 0.1, 0.15)
	tF(1, 0.25, 0.35)
	tF(2, 0.49, 0.51)
	tF(3, 0.65, 0.75)
	tF(4, 0.8, 0.9)
}
