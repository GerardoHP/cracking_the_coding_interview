package chapter_01_test

import (
	"github.com/GerardoHP/cracking_the_coding_interview/pkg/chapter_01"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	// Arrange
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 0},
		{7, 8, 9},
	}
	expected := [][]int{
		{1, 2, 0},
		{0, 0, 0},
		{7, 8, 0},
	}

	// Act
	result := chapter_01.ZeroMatrix(arr)

	// Assert
	if !AreSlicesEqual(expected, result) {
		t.Fatalf("expected to be equal to %v \n", expected)
	}
}

func TestZeroMatrix2(t *testing.T) {
	// Arrange
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Act
	result := chapter_01.ZeroMatrix(arr)

	// Assert
	if !AreSlicesEqual(expected, result) {
		t.Fatalf("expected to be equal to %v \n", expected)
	}
}
