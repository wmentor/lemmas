package counter

// each keyword iterator func
type EachFunc func(key string, counter int64)

// Counter type
type Counter interface {
	Inc(key string)       // inc key name counter by 1
	Names() []string      // all key names
	Each(fn EachFunc)     // each in random order
	EachFreq(fn EachFunc) // each by frequency
	Reset()               // remove all data
	Total() int64         // all key counters summa
	Size() int            // key number
}

// Create new counter
func New() Counter {
	return &mapcnt{data: make(map[string]int64)}
}
