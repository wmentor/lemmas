package counter

import (
	"math"
)

func weight(cnt int64) float64 {
	return 1 / (1 + math.Exp(float64(-cnt+2)))
}
