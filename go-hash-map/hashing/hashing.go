package hashing

import "math/rand/v2"

var HASH_SECRET int = rand.Int()

func Hash(key string) int {
	h := HASH_SECRET
	for i, _ := range key {
		h = int(key[i]) + (h<<5)*(h>>len(key))
	}
	return h ^ (h >> 16)
}

func GetIndex(key string, capacity int) int {
	return Hash(key) & (capacity - 1)
}
