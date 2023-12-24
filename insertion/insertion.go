package insertion

// Time complexity:
//   - O(n) best-case
//   - O(n^2) avg case
//   - O(n^2) worst-case
// Space complexity: O(1)
// Adaptive - i.e. appropriate for partially-sorted datasets
// Algorithmic paradigm: incremental
// In-place sorting algorithm
// Stable

func Sort(input []map[string]int, key string) []map[string]int {
	output := make([]map[string]int, len(input))
	for i, e := range input {
		output[i] = map[string]int{}
		for k, v := range e {
			output[i][k] = v
		}
	}
	sorted := false
	for {
		for i := 0; i < len(output); i++ {
			if i == 0 {
				sorted = true
				continue
			}
			if output[i-1][key] > output[i][key] {
				sorted = false
				tmp := output[i]
				output[i] = output[i-1]
				output[i-1] = tmp
				if i > 1 {
					for j := i - 1; j > 0; j-- {
						if output[j-1][key] > tmp[key] {
							output[j] = output[j-1]
							output[j-1] = tmp
						}
					}
				}
			}
		}
		if sorted {
			break
		}
	}
	return output
}
