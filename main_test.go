package main

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/zysim/SortingAlgos/bubble"
	"github.com/zysim/SortingAlgos/insertion"
	"github.com/zysim/SortingAlgos/merge"
	"github.com/zysim/SortingAlgos/quick"
)

func BenchmarkAlgos(b *testing.B) {
	input := loadInput()
	l := len(input) - 1

	algos := []struct {
		name    string
		handler func([]int) []int
	}{
		{
			name:    "Quicksort w/ Lomuto",
			handler: func(i []int) []int { return quick.Lomuto(i, 0, l) },
		},
		{
			name:    "Bubblesort",
			handler: bubble.Sort,
		},
		{
			name:    "Insertionsort",
			handler: insertion.SortInt,
		},
		{
			name:    "Mergesort",
			handler: merge.Sort,
		},
	}

	for _, algo := range algos {
		algo := algo
		b.Run(algo.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				algo.handler(input)
			}
		})
	}
}

func loadInput() (input []int) {
	raw, _ := os.ReadFile("test_input.txt")
	contents := strings.Split(string(raw), " ")
	for _, c := range contents {
		i, _ := strconv.Atoi(c)
		input = append(input, i)
	}
	return
}
