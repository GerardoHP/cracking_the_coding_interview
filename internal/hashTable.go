package internal

import (
	"math"
	"strings"
)

type hashElement struct {
	Key   string
	Value string
}

type HashTable struct {
	vals  [tableSize][]hashElement
	count int
}

const (
	base         uint64 = 62
	characterSet        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	tableSize           = 5
)

func (ht *HashTable) AddOrUpdate(key, value string) bool {
	hash := calculateHash(key)
	for i, v := range ht.vals[hash] {
		if v.Key == key {
			ht.vals[hash][i].Value = value
			return true
		}
	}

	ht.vals[hash] = append(ht.vals[hash], hashElement{key, value})
	ht.count++
	return true
}

func (ht *HashTable) Get(key string) (string, bool) {
	val, found, _, _ := internalGet(*ht, key)
	return val, found
}

func (ht *HashTable) Length() int {
	return ht.count
}

func (ht *HashTable) Remove(key string) bool {
	_, found, index, hash := internalGet(*ht, key)
	if found {
		c := make([]hashElement, len(ht.vals[hash]))
		copy(c, ht.vals[hash])
		if len(c) == 1 {
			ht.vals[hash] = make([]hashElement, 0)
			ht.count--
			return found
		}

		ht.vals[hash] = make([]hashElement, 0)
		ht.vals[hash] = append(ht.vals[hash], c[:index]...)
		ht.vals[hash] = append(ht.vals[hash], c[index+1:]...)
		ht.count--
	}

	return found
}

func internalGet(ht HashTable, key string) (string, bool, int, int) {
	hash := calculateHash(key)
	for i, v := range ht.vals[hash] {
		if v.Key == key {
			return v.Value, true, i, hash
		}
	}

	return "", false, -1, hash
}

func calculateHash(str string) int {
	decodeStr := stringToBase62(str)
	return decodeStr % tableSize
}

func stringToBase62(str string) int {
	var val uint64
	for k, v := range str {
		pow := len(str) - (k + 1)
		pos := strings.IndexRune(characterSet, v)
		if pos == -1 {
			continue
		}

		val += uint64(pos) * uint64(math.Pow(float64(base), float64(pow)))
	}

	return int(val)
}
