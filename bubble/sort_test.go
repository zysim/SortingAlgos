package bubble

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	input := []int{3, 4, 1, 5, 10, 9, 2, 8, 6, 7}
	actual := fmt.Sprint(Sort(input))
	if actual != "[1 2 3 4 5 6 7 8 9 10]" {
		t.Errorf(
			"Sort incorrect.\nExpected: %s\nGot: %s\n",
			"[1 2 3 4 5 6 7 8 9 10]",
			actual,
		)
	}
}

func BenchmarkSort_BySize(b *testing.B) {
	testCases := []struct {
		name string
		size int
	}{
		{
			name: "10",
			size: 10, // 62.82 ns/op
		},
		{
			name: "1e3",
			size: 1e3, // 551.551 us/op
		},
		{
			name: "1e4",
			size: 1e4, // 53.091474 ms/op
		},
		{
			name: "1e5",
			size: 1e5, // 5.591990750 s/op
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
				Sort(input)
			}
		})
	}
}
