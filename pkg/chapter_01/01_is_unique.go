package chapter_01

// IsUnique Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional
// data structures.
func IsUnique(input string) bool {
	d := make(map[rune]bool)
	for _, v := range input {
		if _, found := d[v]; !found {
			d[v] = true
		} else {
			return false
		}
	}

	return true
}
