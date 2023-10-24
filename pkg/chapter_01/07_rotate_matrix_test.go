package chapter_01_test

import (
	"github.com/GerardoHP/cracking_the_coding_interview/pkg/chapter_01"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	m := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	c := [][]int{
		{13, 9, 5, 1},
		{14, 10, 6, 2},
		{15, 11, 7, 3},
		{16, 12, 8, 4},
	}

	chapter_01.RotateMatrix(&m)
	if !AreSlicesEqual(m, c) {
		t.Fatalf("expected to be the same array")
	}
}

func AreSlicesEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
