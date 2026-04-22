package main

import (
	"math/rand"
)

type Entry struct {
	key   string
	value int
}

type HashMap struct {
	items    [][]Entry
	capacity int
}

var HASH_SECRET int = rand.Int()

func hash(key string) int {
	h := HASH_SECRET
	for i, _ := range key {
		h = int(key[i]) + (h<<5)*(h>>len(key))
	}
	return h ^ (h >> 16)
}

func (hm *HashMap) getIndex(key string) int {
	return hash(key) & (hm.capacity - 1)
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
