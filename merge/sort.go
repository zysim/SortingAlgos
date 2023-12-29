package merge

func Sort(input []int) []int {
	if len(input) < 2 {
		return input
	}

	if len(input) == 2 {
		lo := input[0]
		hi := input[1]

		if lo <= hi {
			return []int{lo, hi}
		}

		return []int{hi, lo}
	}

	lo, hi := split(input)

	lo = Sort(lo)
	hi = Sort(hi)

	output := make([]int, len(input))

	loIndex := 0
	hiIndex := 0

	for i := 0; i < len(output); i++ {
		if loIndex == len(lo) {
			output[i] = hi[hiIndex]
			hiIndex += 1
			continue
		}
		if hiIndex == len(hi) {
			output[i] = lo[loIndex]
			loIndex += 1
			continue
		}
		if lo[loIndex] <= hi[hiIndex] {
			output[i] = lo[loIndex]
			loIndex += 1
			continue
		}
		if lo[loIndex] > hi[hiIndex] {
			output[i] = hi[hiIndex]
			hiIndex += 1
		}
	}

	return output
}

func split(from []int) ([]int, []int) {
	mid := len(from) / 2
	return from[:mid], from[mid:]
}
