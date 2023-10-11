package chapter_01

// Urlify Write a method to replace all spaces in a string with '%20'. You may assume that the string has sufficient space at
// the end to the hold the additional characters, and that you are given the "true" length of the string. (Note: If
// implementing in Java, please use a character array so that you can perform this operation in place.)
func Urlify(str []rune, size int) string {
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			t := make([]rune, len(str[i+1:]))
			copy(t[:], str[i+1:])
			str[i] = '%'
			i++
			str[i] = '2'
			i++
			str[i] = '0'
			i++
			copy(str[i:], t[:])
		}
	}

	return string(str)
}
