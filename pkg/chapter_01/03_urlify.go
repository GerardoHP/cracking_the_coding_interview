package chapter_01

// Urlify Write a method to replace all spaces in a string with '%20'. You may assume that the string has sufficient space at
// the end to the hold the additional characters, and that you are given the "true" length of the string. (Note: If
// implementing in Java, please use a character array so that you can perform this operation in place.)
func Urlify(str []rune, size int) string {
	index := len(str) - size
	copy(str[index:], str[:len(str)])
	for i := len(str) - 1; i >= 0 && index > 0; i-- {
		if str[i] == ' ' {
			copy(str[index-2:i], str[index:i])
			index -= 2
			i -= 2
			copy(str[i:], []rune("%20"))
		}
	}

	return string(str)
}
