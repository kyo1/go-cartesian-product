package cartesian

func All(set []string) func() []string {
	var positions []int = []int{0}

	return func() (element []string) {
		for _, pos := range positions {
			element = append(element, set[pos])
		}

		for i := 0; i < len(positions); i++ {
			positions[i]++
			if positions[i] == len(set) {
				positions[i] = 0
				if i == len(positions)-1 {
					positions = append(positions, 0)
					break
				}
				continue
			}
			break
		}

		return element
	}
}

func Product(set []string, length int) (product [][]string) {
	var positions []int = make([]int, length)

	for {
		var pair []string
		for _, pos := range positions {
			pair = append(pair, set[pos])
		}
		product = append(product, pair)

		var count int // Count the number of carry
		for i := 0; i < length; i++ {
			positions[i]++
			if positions[i] == len(set) {
				positions[i] = 0
				count++
				continue
			}
			break
		}

		if count == length {
			break
		}
	}
	return product
}
