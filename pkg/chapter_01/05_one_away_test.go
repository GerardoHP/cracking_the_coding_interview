package chapter_01_test

import (
	"github.com/GerardoHP/cracking_the_coding_interview/pkg/chapter_01"
	"testing"
)

type testOneAwayArgs struct {
	str1     string
	str2     string
	expected bool
	result   bool
	testId   int
}

func newTestOneAwayArgs(str1, str2 string, expected bool, testId int) *testOneAwayArgs {
	return &testOneAwayArgs{
		str1:     str1,
		str2:     str2,
		expected: expected,
		testId:   testId,
	}
}
func TestOneAway(t *testing.T) {
	tableTests := []*testOneAwayArgs{
		newTestOneAwayArgs("pale", "ple", true, 1),
		newTestOneAwayArgs("pales", "pale", true, 2),
		newTestOneAwayArgs("pale", "bale", true, 3),
		newTestOneAwayArgs("pale", "bake", false, 4),
	}

	for _, v := range tableTests {
		r := chapter_01.OneAway(v.str1, v.str2)
		if r != v.expected {
			t.Fatalf("expected to be %v, but found %v in %v \n", v.expected, r, v.testId)
		}
	}
}
