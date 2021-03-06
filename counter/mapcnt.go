package counter

import (
	"sort"
)

type mapcnt struct {
	data map[string]int64
}

// return all known keys
func (m *mapcnt) KeyNames() []string {
	res := make([]string, 0, len(m.data))
	for k := range m.data {
		res = append(res, k)
	}
	return res
}

// return []*Keys ordered by counter value desc
func (m *mapcnt) Keys() []*Key {

	res := make([]*Key, 0, len(m.data))

	for k, v := range m.data {
		res = append(res, &Key{Name: k, Counter: v, Weight: weight(v)})
	}

	sort.Slice(res, func(i int, j int) bool {
		v1 := res[i]
		v2 := res[j]
		return v1.Counter > v2.Counter || v1.Counter == v2.Counter && v1.Name < v2.Name
	})

	return res
}

func (m *mapcnt) Inc(key string) {
	m.data[key]++
}

func (m *mapcnt) Reset() {
	m.data = make(map[string]int64)
}
