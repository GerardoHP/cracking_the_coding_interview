package chapter_01

import "math"

// OneAway There Are three types of edits that can be performed on strings: insert a character, remove a character, or
// replace a character. Given ple strings, write a function to check if they are one edit (or zero edits) away.
func OneAway(str1, str2 string) bool {
	if math.Abs(float64(len(str1)-len(str2))) > 1 {
		return false
	}

	var bigArray *string
	var smallArray *string

	if len(str1) > len(str2) {
		bigArray = &str1
		smallArray = &str2
	} else {
		bigArray = &str2
		smallArray = &str1
	}

	return getResultFromTwoArrays(bigArray, smallArray)
}

func getResultFromTwoArrays(bigArray, smallArray *string) bool {
	i, j := 0, 0
	d := false
	for j < len(*smallArray) {
		if (*bigArray)[i] != (*smallArray)[j] {
			if d {
				return false
			}

			d = true
			if len(*bigArray) != len(*smallArray) {
				i++
			}
		}

		i++
		j++
	}

	return true
}
