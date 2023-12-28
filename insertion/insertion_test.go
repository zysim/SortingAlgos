package insertion

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSortInt(t *testing.T) {
	input := make([]int, 1000)
	for i := 0; i < len(input); i++ {
		input[i] = i + 1
	}
	rand.Shuffle(len(input), func(i, j int) { input[i], input[j] = input[j], input[i] })

	output := SortInt(input)

	for i, o := range output {
		if o != i+1 {
			t.Error("SortInt incorrect")
		}
	}
}

func TestSortDict(t *testing.T) {
	input := []map[string]int{
		{
			"a": 11,
			"b": 2,
		},
		{
			"a": 5,
			"b": 1,
		},
		{
			"a": 13,
			"b": 4,
		},
		{
			"a": 12,
			"b": 3,
		},
		{
			"a": 6,
			"b": 2,
		},
	}

	expectedA := []int{5, 6, 11, 12, 13}
	expectedB := []int{1, 2, 2, 3, 4}

	output := SortDict(input, "a")

	for i, o := range output {
		if o["a"] != expectedA[i] {
			t.Error("Insertion sort failure")
		}
	}

	// Test sort stability; input order of elements w/ same key
	// shouldn't change
	output = SortDict(SortDict(input, "a"), "b")
	for i, o := range output {
		if o["a"] != expectedA[i] || o["b"] != expectedB[i] {
			t.Error("Insertion sort stability failure")
		}
	}
}

func BenchmarkSortInt(b *testing.B) {
	N := 1000
	worstCase := make([]int, N)
	for i := N; i > 0; i-- {
		worstCase[N-i] = i
	}
	bestCase := make([]int, N)
	for i := 0; i < N; i++ {
		bestCase[i] = i
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
			name:  "Worst case",
			input: worstCase,
		},
		{
			name:  "Best case",
			input: bestCase,
		},
	}
	for i, a := range avgCases {
		testCases = append(testCases, struct {
			name  string
			input []int
		}{
			name:  fmt.Sprintf("Average case %d\n", i+1),
			input: a,
		})
	}

	for _, tc := range testCases {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SortInt(tc.input)
			}
		})
	}
}

func BenchmarkSortDict(b *testing.B) {
	worstCase := make([]map[string]int, 1000)
	for i := 0; i < len(worstCase); i++ {
		worstCase[i] = map[string]int{"a": 1000 - i}
	}
	bestCase := make([]map[string]int, 1000)
	for i := 0; i < len(bestCase); i++ {
		bestCase[i] = map[string]int{"a": i}
	}
	avgCase := make([]map[string]int, 1000)
	for i := 0; i < len(avgCase); i++ {
		avgCase[i] = map[string]int{"a": rand.Intn(1000)}
	}

	testCases := []struct {
		name  string
		input []map[string]int
	}{
		{
			name:  "Worst case",
			input: worstCase,
		},
		{
			name:  "Best case",
			input: bestCase,
		},
		{
			name:  "Average case",
			input: avgCase,
		},
	}

	for _, tc := range testCases {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SortDict(tc.input, "a")
			}
		})
	}
}
