package linearprobing

import (
	"github.com/ravi-sankarp/prototypes/go-hash-map/hashing"
)

type entry struct {
	key   string
	value int
}

type hashmap struct {
	items    []entry
	capacity int
}

func (hm *hashmap) rehash() {
	copyItems := make([]entry, hm.capacity, hm.capacity)
	hm.capacity *= 2
	copy(copyItems, hm.items)
	hm.items = make([]entry, hm.capacity)
	for _, val := range copyItems {
		if val.key != "" {
			hm.set(val.key, val.value)
		}
	}
}

func (hashmap *hashmap) set(key string, value int) {
	index := hashing.GetIndex(key, hashmap.capacity)
	if hashmap.items[index].key == "" {
		hashmap.items[index] = entry{key, value}
		return
	}
	for i := (index + 1) % hashmap.capacity; i != index; i = (i + 1) % hashmap.capacity {
		if hashmap.items[i].key == "" {
			hashmap.items[i] = entry{key, value}
			return
		}
	}
	hashmap.rehash()
	hashmap.set(key, value)
}

func (hashmap *hashmap) get(key string) (int, bool) {
	index := hashing.GetIndex(key, hashmap.capacity)
	if hashmap.items[index].key == "" {
		return 0, false
	}
	for i := index; ; {
		if hashmap.items[i].key == key {
			return hashmap.items[i].value, true
		}
		i = (i + 1) % hashmap.capacity
		if i == index {
			break
		}
	}
	return 0, false
}

func newMap(capacity int) *hashmap {
	return &hashmap{
		items:    make([]entry, capacity),
		capacity: capacity,
	}
}
