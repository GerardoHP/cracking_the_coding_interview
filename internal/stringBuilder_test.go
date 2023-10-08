package internal_test

import (
	"github.com/GerardoHP/cracking_the_coding_interview/internal"
	"testing"
)

func TestStringBuilder_Append(t *testing.T) {
	// Arrange
	strBuilder := internal.StringBuilder{}
	expectedResult := "aaa"

	// Act
	strBuilder.Append("a")
	strBuilder.Append("a")
	strBuilder.Append("a")
	result := strBuilder.String()

	// Assert
	if expectedResult != result {
		t.Fatalf("expected %v, but found %v \n", expectedResult, result)
	}
}
