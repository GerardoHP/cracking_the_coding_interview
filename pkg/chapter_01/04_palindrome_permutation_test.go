package chapter_01_test

import (
	"github.com/GerardoHP/cracking_the_coding_interview/pkg/chapter_01"
	"testing"
)

func TestPalindromePermutation(t *testing.T) {
	// Arrange
	str := "Tact Coa"

	// Act
	output := chapter_01.PalindromePermutation(str)

	// Assert
	if !output {
		t.Fatalf("expected to be true")
	}
}

func TestPalindromePermutationFalse(t *testing.T) {
	// Arrange
	str := "Tact Coat"

	// Act
	output := chapter_01.PalindromePermutation(str)

	// Assert
	if output {
		t.Fatalf("expected to be true")
	}
}
