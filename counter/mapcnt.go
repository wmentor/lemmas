package counter

import (
	"sort"
)

type mapcnt struct {
	data  map[string]int64
	total int64
}

// return all known keys
func (m *mapcnt) Names() []string {
	res := make([]string, 0, len(m.data))
	for k := range m.data {
		res = append(res, k)
	}
	return res
}

// return sum of all keys counters
func (m *mapcnt) Total() int64 {
	return m.total
}

// return key number
func (m *mapcnt) Size() int {
	return len(m.data)
}

// inc single key counter
func (m *mapcnt) Inc(key string) {
	m.data[key]++
	m.total++
}

// reset all data
func (m *mapcnt) Reset() {
	m.data = make(map[string]int64)
	m.total = 0
}

// iterate over each key
func (m *mapcnt) Each(fn EachFunc) {
	for k, v := range m.data {
		fn(k, v)
	}
}

// iterate over each key order by frequency
func (m *mapcnt) EachFreq(fn EachFunc) {

	list := make([]string, 0, len(m.data))

	for k := range m.data {
		list = append(list, k)
	}

	sort.Slice(list, func(i int, j int) bool {
		vi := m.data[list[i]]
		vj := m.data[list[j]]
		return vi > vj || vi == vj && list[i] < list[j]
	})

	for _, k := range list {
		fn(k, m.data[k])
	}
}
