package hashing

import "testing"

func TestHash(t *testing.T) {
	hash1 := Hash("hello")
	hash2 := Hash("hello2")
	t.Log(hash1, hash2)
	if hash1 != hash2 {
		t.Fatal("Inconsitent Hash")
	}
	t.Log("SUCCESS")
}

func TestIndexSpreading(t *testing.T) {
	capacity := 64
	index1 := GetIndex("key-1", capacity)
	index2 := GetIndex("key-2", capacity)
	index3 := GetIndex("key-3", capacity)
	index4 := GetIndex("key-4", capacity)
	index5 := GetIndex("key-5", capacity)
	t.Log(index1, index2, index3, index4, index5)
	indexMap := map[int]int{
		index1: index1,
		index2: index2,
		index3: index3,
		index4: index4,
		index5: index5,
	}
	if len(indexMap) <= 3 {
		t.Fatal("Uneven distribution")
	}
	t.Log("SUCCESS")
}
