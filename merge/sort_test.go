package merge

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {
	input := []int{3, 1, 4, 2, 6, 5, 10, 9, 7, 8}
	actual := fmt.Sprintf("%v", Sort(input))

	if actual != "[1 2 3 4 5 6 7 8 9 10]" {
		t.Errorf("Slice not sorted correctly. Got: %s\n", actual)
	}
}

func BenchmarkSort(b *testing.B) {
	N := 1000
	worstCase := make([]int, N)
	for i := range worstCase {
		worstCase[i] = N - i
	}
	bestCase := make([]int, N)
	for i := range bestCase {
		bestCase[i] = i + 1
	}
	avgCases := make([][]int, 3)
	for i, a := range avgCases {
		a = make([]int, N)
		copy(a, bestCase)
		rand.Shuffle(N, func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})
		avgCases[i] = a
	}

	testCases := []struct {
		name  string
		input []int
	}{
		{
			name:  "Worst-case",
			input: worstCase,
		},
		{
			name:  "Best-case",
			input: bestCase,
		},
	}
	for i := range avgCases {
		testCases = append(testCases, struct {
			name  string
			input []int
		}{
			name:  fmt.Sprintf("Average case %d", i+1),
			input: avgCases[i],
		})
	}

	for _, tc := range testCases {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Sort(tc.input)
			}
		})
	}
}

// Benchmark shows that mergesort scales linearly in time, with size
func BenchmarkSortSize(b *testing.B) {
	first := make([]int, 10)
	second := make([]int, 1e3)
	third := make([]int, 1e4)
	fourth := make([]int, 1e5)

	for i := range first {
		first[i] = i + 1
	}
	for i := range second {
		second[i] = i + 1
	}
	for i := range third {
		third[i] = i + 1
	}
	for i := range fourth {
		fourth[i] = i + 1
	}

	testCases := []struct {
		name  string
		input []int
	}{
		{
			name:  "Size 10",
			input: first,
		},
		{
			name:  "Size 1000",
			input: second,
		},
		{
			name:  "Size 1e4",
			input: third,
		},
		{
			name:  "Size 1e5",
			input: fourth,
		},
	}

	for _, tc := range testCases {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Sort(tc.input)
			}
		})
	}
}
