package bubble

// Time complexity:
//   - O(n) best-case
//   - O(n^2) avg case
//   - O(n^2) worst-case
// Space complexity: O(n) total; O(1) auxiliary
// Appropriate for partially-sorted datasets
// In-place sorting algorithm
// Stable

func Sort(input []int) (output []int) {
	l := len(input)
	output = make([]int, l)
	copy(output, input)
	for i := 0; i <= l/2; i++ {
		for j := 0; j < l-i-1; j++ {
			if output[j] > output[j+1] {
				output[j], output[j+1] = output[j+1], output[j]
				continue
			}
		}
	}
	return output
}
