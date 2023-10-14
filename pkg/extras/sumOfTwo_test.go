package extras_test

import (
	"github.com/GerardoHP/cracking_the_coding_interview/pkg/extras"
	"testing"
)

func TestSumOfTwo_FoundValue(t *testing.T) {
	arr1 := []int{1, 2, 83}
	arr2 := []int{1, 2, 83}
	search := 84

	found, i1, i2 := extras.SumOfTwo(arr1, arr2, search)

	if !found && i1 != 2 && i2 != 0 {
		t.Fatalf("not expected to fail")
	}
}

func TestSumOfTwo_NotFoundValue(t *testing.T) {
	arr1 := []int{1, 2, 83}
	arr2 := []int{1, 2, 83}
	search := 87

	found, _, _ := extras.SumOfTwo(arr1, arr2, search)

	if found {
		t.Fatalf("not expected to fail")
	}
}
