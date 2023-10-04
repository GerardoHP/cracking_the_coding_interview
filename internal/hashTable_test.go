package internal_test

import (
	"github.com/GerardoHP/cracking_the_coding_interview/internal"
	"testing"
)

func TestHashTable_AddOrUpdate(t *testing.T) {
	// arrange
	table := internal.HashTable{}
	table.AddOrUpdate("b-ba", "primer valor")
	table.AddOrUpdate("bb+a", "segundo valor")
	table.AddOrUpdate("bb-a", "tercer valor")
	expectedCount := 3

	// act, assert
	if l := table.Length(); l != expectedCount {
		t.Fatalf("expected %v, but found %v \n", expectedCount, l)
	}
}

func TestHashTable_Get(t *testing.T) {
	// arrange
	key := "bb+a"
	value := "primer valor"
	table := internal.HashTable{}
	table.AddOrUpdate(key, value)
	table.AddOrUpdate("bb-a", "segundo valor")

	// act, assert
	if v, found := table.Get(key); !found || v != value {
		t.Fatalf("expected value %v, but nothing was found or found %v \n", value, v)
	}
}

func TestHashTable_Get2(t *testing.T) {
	// arrange
	table := internal.HashTable{}

	// act, assert
	if v, found := table.Get("nothing"); found {
		t.Fatalf("not expected any value, but found %v \n", v)
	}
}

func TestHashTable_AddOrUpdate2(t *testing.T) {
	// arrange
	updatedValue := "nuevo valor"
	key := "bb+a"
	table := internal.HashTable{}
	table.AddOrUpdate("b-ba", "primer valor")
	table.AddOrUpdate(key, "segundo valor")
	table.AddOrUpdate(key, updatedValue)
	expectedCount := 2

	// act, assert
	if l := table.Length(); l != expectedCount {
		t.Fatalf("expected %v, but found %v \n", expectedCount, l)
	}

	if v, found := table.Get(key); !found {
		t.Fatalf("expected value")
	} else {
		if v != updatedValue {
			t.Fatalf("expected value %v, but found %v \n", updatedValue, v)
		}
	}
}

func TestHashTable_Delete(t *testing.T) {
	// arrange
	table := internal.HashTable{}
	key := "b-ba"
	value := "primer valor"
	table.AddOrUpdate(key, value)
	table.AddOrUpdate("bb+a", "segundo valor")
	table.AddOrUpdate("bb-a", "tercer valor")
	table.Remove(key)
	expectedCount := 2

	// act, assert
	if l := table.Length(); l != expectedCount {
		t.Fatalf("expected %v, but found %v \n", expectedCount, l)
	}

	if val, found := table.Get(key); found {
		t.Fatalf("not expecte any value but found %v \n", val)
	}
}

func TestHashTable_Remove(t *testing.T) {
	// arrange
	table := internal.HashTable{}
	key := "bb+a"
	value := "segundo valor"
	table.AddOrUpdate("b-ba", "primer valor")
	table.AddOrUpdate(key, value)
	table.AddOrUpdate("bb-a", "tercer valor")
	table.Remove(key)
	expectedCount := 2

	// act, assert
	if l := table.Length(); l != expectedCount {
		t.Fatalf("expected %v, but found %v \n", expectedCount, l)
	}

	if val, found := table.Get(key); found {
		t.Fatalf("not expecte any value but found %v \n", val)
	}
}
