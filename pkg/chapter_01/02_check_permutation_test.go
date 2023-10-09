package chapter_01_test

import (
	"testing"

	"github.com/GerardoHP/cracking_the_coding_interview/pkg/chapter_01"
)

func TestCheckPermutation_Different_Size(t *testing.T) {
	// Arrange
	str1 := "aoeu"
	str2 := "aoeui"

	// Act
	result := chapter_01.CheckPermutation(str1, str2)

	// Assert
	if result {
		t.Fatalf("expected result to be false")
	}
}

func TestCheckPermutation_true(t *testing.T) {
	// Arrange
	str1 := "abc"
	str2 := "bca"

	// Act
	result := chapter_01.CheckPermutation(str1, str2)

	// Assert
	if !result {
		t.Fatalf("expected result to be true")
	}
}

func TestCheckPermutation_false(t *testing.T) {
	// Arrange
	str1 := "abc"
	str2 := "bcc"

	// Act
	result := chapter_01.CheckPermutation(str1, str2)

	// Assert
	if result {
		t.Fatalf("expected result to be false")
	}
}
