package chapter_01_test

import (
	"github.com/GerardoHP/cracking_the_coding_interview/pkg/chapter_01"
	"testing"
)

func TestIsUnique_true(t *testing.T) {
	// Arrange
	unique := "quick brown"

	// Act
	result := chapter_01.IsUnique(unique)

	// Assert
	if !result {
		t.Fatalf("expected to bo true")
	}
}

func TestIsUnique_false(t *testing.T) {
	// Arrange
	notUnique := "this is not unique"

	// Act
	result := chapter_01.IsUnique(notUnique)

	// Assert
	if result {
		t.Fatalf("expected to be false")
	}
}
