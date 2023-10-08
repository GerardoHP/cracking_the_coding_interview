package internal_test

import (
	"github.com/GerardoHP/cracking_the_coding_interview/internal"
	"testing"
)

func TestArrayList_Add(t *testing.T) {
	// arrange
	al := internal.NewArrayList()
	expectedCount := 3
	expectedCapacity := 4

	// act
	al.Add(1)
	al.Add(2)
	al.Add(3)

	// assert
	if expectedCount != al.Count {
		t.Fatalf("expected count %v, but found %v \n", expectedCount, al.Count)
	}
	if expectedCapacity != al.Capacity {
		t.Fatalf("expected capacity %v, but found %v \n", expectedCapacity, al.Capacity)
	}
}

func TestArrayList_Get(t *testing.T) {
	// arrange
	al := internal.NewArrayList()
	expectedVal := 3

	// act
	al.Add(1)
	al.Add(2)
	al.Add(expectedVal)
	val, _ := al.Get(2)
	_, err := al.Get(4)

	// assert
	if expectedVal != val {
		t.Fatalf("expected val %v, but found %v \n", expectedVal, val)
	}

	if err == nil {
		t.Fatalf("error expected to be nil")
	}
}

func TestArrayList_Delete(t *testing.T) {
	// arrange
	al := internal.NewArrayList()
	expectedVal := 3
	expectedCount := 2

	// act
	al.Add(1)
	al.Add(2)
	al.Add(expectedVal)
	_ = al.Delete(1)
	val, _ := al.Get(1)
	err := al.Delete(20)

	// assert
	if val != expectedVal {
		t.Fatalf("expected val %v, but found %v \n", expectedVal, val)
	}

	if l := al.Count; expectedCount != l {
		t.Fatalf("expected count %v, but found %v \n", expectedCount, l)
	}

	if err == nil {
		t.Fatalf("error expected to be nil")
	}
}
