package sorts

type CocktailSort struct {
}

func (c CocktailSort) Sort(arr []int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}

		right -= 1

		for i := right; i < left; i-- {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}

		left += 1
	}

	return arr
}
