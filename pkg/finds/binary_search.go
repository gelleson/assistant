package finds

func BinarySearch(searchTerm int, arr []int) int {
	left := 0

	right := len(arr) - 1

	for left <= right {
		m := (left + right) / 2

		if arr[m] > searchTerm {

			left = m + 1
		} else if arr[m] < searchTerm {
			right = m - 1

		} else {
			return arr[m]
		}
	}

	return -1

}
