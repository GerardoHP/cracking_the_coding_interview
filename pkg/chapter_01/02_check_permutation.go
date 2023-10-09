package chapter_01

// CheckPermutation Given two strings, write a method to decide if one is a permutation of the other
func CheckPermutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	m1, m2 := stringToMap(str1), stringToMap(str2)
	for k, v := range m1 {
		if v != m2[k] {
			return false
		}
	}

	return true
}

func stringToMap(str string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range str {
		if _, found := m[r]; !found {
			m[r] = 1
		} else {
			m[r]++
		}
	}

	return m
}
