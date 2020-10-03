package sorts

type BubbleSort struct {
}

func (s BubbleSort) Sort(un []int) []int {
	sortedSlice := un

	total := len(un)

	for i := 0; i < total-1; i++ {
		for j := 0; j < total-1-i; j++ {
			f := 0
			if sortedSlice[j] > sortedSlice[j+1] {
				sortedSlice[j+1], sortedSlice[j] = sortedSlice[j], sortedSlice[j+1]
				f = 1
			}

			if f == 0 {
				continue
			}
		}
	}

	return sortedSlice
}
