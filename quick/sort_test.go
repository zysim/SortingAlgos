package quick

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestLomuto(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected string
	}{
		{
			name:     "Even-numbered",
			input:    []int{3, 4, 1, 5, 10, 9, 2, 8, 6, 7},
			expected: "[1 2 3 4 5 6 7 8 9 10]",
		},
		{
			name:     "Odd-numbered",
			input:    []int{3, 4, 1, 5, 9, 2, 8, 6, 7},
			expected: "[1 2 3 4 5 6 7 8 9]",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual := fmt.Sprint(Lomuto(tc.input, 0, len(tc.input)-1))
			if actual != tc.expected {
				t.Errorf(
					"Sort incorrect.\nExpected: %s\nGot: %s\n",
					tc.expected,
					actual,
				)
			}
		})
	}
}

func BenchmarkLomuto_BySize(b *testing.B) {
	testCases := []struct {
		name string
		size int
	}{
		{
			name: "10",
			size: 10, // 311.2 ns/op
		},
		{
			name: "1e3",
			size: 1e3, // 1.720129 ms/op
		},
		{
			name: "1e4",
			size: 1e4, // 224.013292 ms/op
		},
		{
			name: "1e5",
			size: 1e5, // 31.246184333 s/op
		},
	}
	for _, tc := range testCases {
		tc := tc
		input := make([]int, tc.size)
		for i := range input {
			input[i] = tc.size - i
		}
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Lomuto(input, 0, tc.size-1)
			}
		})
	}
}

func BenchmarkLomuto_ByCase(b *testing.B) {
	N := 1000
	bestCase := make([]int, N)
	for i := range bestCase {
		bestCase[i] = i
	}
	worstCase := make([]int, N)
	for i := range worstCase {
		worstCase[i] = N - i
	}
	avgCases := make([][]int, 5)
	for i, a := range avgCases {
		a = make([]int, N)
		copy(a, bestCase)
		rand.Shuffle(N, func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})
		avgCases[i] = a
	}

	// Algo sorts in-place. Need to create a copy
	input := make([]int, N)
	b.Run("Best-case", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy(input, bestCase)
			Lomuto(input, 0, N-1)
		}
	})
	b.Run("Worst-case", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy(input, worstCase)
			Lomuto(input, 0, N-1)
		}
	})
	for i, a := range avgCases {
		b.Run(fmt.Sprintf("Average case %d", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				copy(input, a)
				Lomuto(input, 0, N-1)
			}
		})
	}
}
