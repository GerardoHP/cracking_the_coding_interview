package chapter_01

import (
	"fmt"
	"strings"
)

// StringCompression Implement a method to perform basic string compression using the counts of repeated characters.
// For example, the string aabccccaaa would become a2b1c5a3. If the "compressed" string would not become smaller than the
// original string, your method should return the original string. You can assume the string has only uppercase and
// lowercase letters (a - z).
func StringCompression(str string) string {
	strBldr := strings.Builder{}
	r := str[0]
	c := 1
	t := 0
	for i := 1; i < len(str); i++ {
		if str[i] == r {
			c++
		} else {
			strBldr.WriteString(string(r))
			strBldr.WriteString(fmt.Sprint(c))
			t += c
			r = str[i]
			c = 1
		}
	}

	if t != len(str) {
		strBldr.WriteString(string(r))
		strBldr.WriteString(fmt.Sprint(c))
	}

	return strBldr.String()
}
