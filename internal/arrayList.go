package internal

import "fmt"

const (
	initialSize = 2
)

type ArrayList struct {
	array    []int
	Capacity int
	Count    int
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		array:    make([]int, initialSize),
		Capacity: initialSize,
		Count:    0,
	}
}

func (al *ArrayList) Add(val int) {
	if al.Capacity == al.Count {
		al.resize()
	}

	al.array[al.Count] = val
	al.Count++
}

func (al *ArrayList) Get(index int) (int, error) {
	if index < -1 || index >= al.Count {
		return -1, fmt.Errorf("index out of bounds")
	}

	return al.array[index], nil
}

func (al *ArrayList) Delete(index int) error {
	if index < -1 || index >= al.Count {
		return fmt.Errorf("index out of bounds")
	}

	newArray := make([]int, al.Capacity)
	copy(newArray[:index], al.array[:index])
	copy(newArray[index:], al.array[index+1:])
	al.array = newArray
	al.Count--
	return nil
}

func (al *ArrayList) resize() {
	newCapacity := al.Capacity * 2
	newArray := make([]int, newCapacity)
	copy(newArray[:al.Count], al.array[:al.Count])
	al.array = newArray
	al.Capacity = newCapacity
}
