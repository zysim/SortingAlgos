package cocktail

import (
	"testing"
)

func TestSort(t *testing.T) {
	N := 100
	input := make([]int, N)
	for i := range input {
		input[i] = N - i - 1
	}
	expected := make([]int, N)
	for i := range expected {
		expected[i] = i
	}
	actual := Sort(input)
	for i := range expected {
		if actual[i] != expected[i] {
			// Errorf by consequence of design doesn't break out of loops
			t.Fatalf("Cocktail sort failed. Output: %v", actual)
		}
	}
}
