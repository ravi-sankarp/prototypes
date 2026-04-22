package nestedarrayaddressing

import "github.com/ravi-sankarp/prototypes/go-hash-map/hashing"

type entry struct {
	key   string
	value int
}

type hashmap struct {
	items    [][]entry
	capacity int
}

func (hashmap *hashmap) set(key string, value int) {
	index := hashing.GetIndex(key, hashmap.capacity)
	bucket := hashmap.items[index]

	for i, e := range bucket {
		if e.key == key {
			bucket[i].value = value
			return
		}
	}
	hashmap.items[index] = append(hashmap.items[index], entry{key, value})
}

func (hm *hashmap) get(key string) (int, bool) {
	bucket := hm.items[hashing.GetIndex(key, hm.capacity)]
	for _, e := range bucket {
		if e.key == key {
			return e.value, true
		}
	}
	return 0, false
}

func newMap(capacity int) *hashmap {
	return &hashmap{
		items:    make([][]entry, capacity),
		capacity: capacity,
	}
}
