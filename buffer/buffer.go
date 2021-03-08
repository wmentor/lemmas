package buffer

type Buffer interface {
	Push(string)
	Get(int) string
	Shift(int)
	Empty() bool
	Full() bool
	Len() int
}

type buffer struct {
	records []string
	used    int
	size    int
	first   int
}

func New(size int) Buffer {
	return &buffer{
		used:    0,
		first:   0,
		size:    size,
		records: make([]string, size),
	}
}

func (b *buffer) Push(str string) {
	b.records[(b.first+b.used)%b.size] = str
	if b.used < b.size {
		b.used++
	} else {
		b.first = (b.first + 1) % b.size
	}
}

func (b *buffer) Get(idx int) string {
	return b.records[(b.first+idx)%b.size]
}

func (b *buffer) Shift(n int) {
	if n >= b.used {
		b.first = 0
		b.used = 0
	} else {
		b.first = (b.first + n) % b.size
		b.used -= n
	}
}

func (b *buffer) Empty() bool {
	return b.used == 0
}

func (b *buffer) Full() bool {
	return b.used == b.size
}

func (b *buffer) Len() int {
	return b.used
}
