package chapter_01_test

import (
	"github.com/GerardoHP/cracking_the_coding_interview/pkg/chapter_01"
	"testing"
)

func TestUrlify(t *testing.T) {
	// Arrange
	expected := "Mr%20John%20Smith"
	str := "Mr John Smith    "
	i := 13

	// Act
	r := chapter_01.Urlify([]rune(str), i)

	// Assert
	if expected != r {
		t.Fatalf("expected %v, but found %v \n", expected, r)
	}
}
