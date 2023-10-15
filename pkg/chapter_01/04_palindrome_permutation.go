package chapter_01

import "strings"

// PalindromePermutation Given a string, write a function to check if it is a permutation of a palindrome. A palindrome
// is a word of phrase that is the same forwards and backwards. A permutation is a rearrangement of letters. The
// palindrome does not need to be limited to just dictionary words.
func PalindromePermutation(str string) bool {
	str = strings.ToLower(str)
	m := make(map[rune]int)
	for _, v := range str {
		if v == ' ' {
			continue
		}

		if _, found := m[v]; !found {
			m[v] = 0
		}

		m[v]++
	}

	odd := false
	for _, v := range m {
		if v%2 == 1 {
			if odd {
				return false
			}

			odd = true
		}
	}

	return true
}
