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
	hash := calculateHash(key)
	for _, v := range ht.vals[hash] {
		if v.Key == key {
			return v.Value, true
		}
	}

	return "", false
}

func (ht *HashTable) Length() int {
	return ht.count
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
