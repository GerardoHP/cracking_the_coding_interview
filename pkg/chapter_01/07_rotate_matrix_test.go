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

	r := chapter_01.RotateMatrix(m)
	if len(r) <= 0 {
		t.Fatalf("expected to be more than one")
	}
}
