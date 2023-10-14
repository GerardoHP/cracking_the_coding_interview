package extras

import "math"

func SumOfTwo(arr1, arr2 []int, result int) (bool, int, int) {
	m := make(map[int]int)
	for k, v := range arr1 {
		m[v] = k
	}

	for k, v := range arr2 {
		val := int(math.Abs(float64(result - v)))
		if _, found := m[val]; found {
			return true, m[val], k
		}

	}

	return false, -1, -1
}
