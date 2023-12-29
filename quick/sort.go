package quick

// Time complexity:
//   - Ω(N log (N)) best-case: when chosen pivot divides slice into roughly-equal halves
//   - θ(N log (N)) avg case
//   - O(N^2) worst-case
// Space complexity: O(1) (auxiliary)
// Appropriate for large datasets
// Not appropriate for small datasets
// In-place sorting algorithm
// Unstable

// Lomuto uses the Lomuto partition algorithm for quicksort.
//
// https://en.wikipedia.org/wiki/Quicksort#Lomuto_partition_scheme
//
// Lomuto(input, 0, len(input)-1): sorts whole slice
//
// Lomuto(input, N, len(input)-1): sorts slice from index N to end, inclusive
func Lomuto(input []int, lo, hi int) []int {
	if lo >= hi || lo < 0 {
		return input
	}

	// Both statements below are shortcuts. Algo still works without these.
	if len(input) < 2 {
		return input
	}
	if len(input) == 2 {
		if input[0] > input[1] {
			input[0], input[1] = input[1], input[0]
		}
		return input
	}

	// Partition array and get pivot index
	p := partition(input, lo, hi)

	// Sort both partitions
	r := append(Lomuto(input, lo, p-1), Lomuto(input, p+1, hi)...)
	return r[len(r)-len(input):]
}

func partition(input []int, lo, hi int) int {
	pivot := input[hi] // Choose last element as pivot
	i := lo - 1        // Tmp pivot pos
	for j := lo; j < hi; j++ {
		if input[j] <= pivot {
			i += 1
			input[i], input[j] = input[j], input[i]
		}
	}

	// Move pivot el to correct pivot pos (between smaller and larger els)
	i += 1
	input[i], input[hi] = input[hi], input[i]
	return i
}
