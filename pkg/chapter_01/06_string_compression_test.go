package chapter_01_test

import (
	"github.com/GerardoHP/cracking_the_coding_interview/pkg/chapter_01"
	"testing"
)

func TestStringCompression(t *testing.T) {
	// Arrange
	str := "aabcccccaaa"
	expected := "a2b1c5a3"

	// Act
	result := chapter_01.StringCompression(str)

	// Assert
	if expected != result {
		t.Fatalf("expected %v, but found %v \n", expected, result)
	}
}
