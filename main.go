package main

import (
	"fmt"

	"github.com/zysim/SortingAlgos/merge"
)

func main() {
	// input := []map[string]int{
	// 	{"first": 11, "second": 5},
	// 	{"first": 5, "second": 4},
	// 	{"first": 13, "second": 2},
	// 	{"first": 12, "second": 1},
	// 	{"first": 6, "second": 1},
	// }

	// input := []int{11, 5, 13, 12, 6}

	fmt.Println(merge.Sort([]int{38, 27, 43, 10}))
}
