package counter

// Counter key record info
type Key struct {
	Name    string
	Counter int64
	Weight  float64
}

// Counter type
type Counter interface {
	Inc(key string)
	KeyNames() []string
	Keys() []*Key
}

// Create new counter
func New() Counter {
	return mapcnt(make(map[string]int64))
}
