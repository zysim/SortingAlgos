package cocktail

func Sort(input []int) []int {
	rev := false
	sorted := true
	from := 0
	til := len(input) - 1
	for {
		if !rev {
			sorted = sortAsc(input, from, til)
			til--
		} else {
			sorted = sortDesc(input, from, til)
			from++
		}
		if sorted {
			break
		}
		rev = !rev
	}
	return input
}

func sortAsc(input []int, from, til int) bool {
	sorted := true
	for i := from; i < til; i++ {
		if input[i] > input[i+1] {
			sorted = false
			input[i], input[i+1] = input[i+1], input[i]
		}
	}
	return sorted
}

func sortDesc(input []int, from, til int) bool {
	sorted := true
	for i := til; i > from; i-- {
		if input[i-1] > input[i] {
			sorted = false
			input[i-1], input[i] = input[i], input[i-1]
		}
	}
	return sorted
}
