package main

type Entry struct {
	key   string
	value int
}

type HashMap struct {
	items    [][]Entry
	capacity int
}

func hash(key string) int {
	h := 0
	for i, _ := range key {
		h += h*i + i
	}
	return h
}

func (hm *HashMap) getIndex(key string) int {
	return hash(key) % hm.capacity
}

func (hashMap *HashMap) set(key string, value int) {
	index := hashMap.getIndex(key)
	bucket := hashMap.items[index]

	for i, e := range bucket {
		if e.key == key {
			bucket[i].value = value
			return
		}
	}
	hashMap.items[index] = append(hashMap.items[index], Entry{key, value})
}

func (hm *HashMap) get(key string) (int, bool) {
	bucket := hm.items[hm.getIndex(key)]
	for _, e := range bucket {
		if e.key == key {
			return e.value, true
		}
	}
	return 0, false
}

func NewMap(capacity int) *HashMap {
	return &HashMap{
		items:    make([][]Entry, capacity),
		capacity: capacity,
	}
}
func main() {}
