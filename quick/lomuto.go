package quick

// https://en.wikipedia.org/wiki/Quicksort#Lomuto_partition_scheme
// Lomuto(input, 0, len(input)-1): sorts whole slice
// Lomuto(input, N, len(input)-1): sorts slice from index N to end, inclusive
func Lomuto(input []int, lo, hi int) []int {
	if lo >= hi || lo < 0 {
		return input
	}

	// Partition array and get pivot index
	p := partition(input, lo, hi)

	// Sort both partitions
	r := append(append(Lomuto(input, lo, p-1), input[p]), Lomuto(input, p+1, hi)...)
	return r[len(r)-len(input):]
}

func partition(input []int, lo, hi int) int {
	pivot := input[hi] // Choose last element as pivot
	i := lo - 1        // Tmp pivot pos
	for j := lo; j < hi; j++ {
		if input[j] <= pivot {
			i += 1
			tmp := input[i]
			input[i] = input[j]
			input[j] = tmp
		}
	}

	// Move pivot el to correct pivot pos (btw smaller and larger els)
	i += 1
	tmp := input[hi]
	input[hi] = input[i]
	input[i] = tmp
	return i
}
