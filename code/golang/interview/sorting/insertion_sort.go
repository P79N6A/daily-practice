package sorting

func less(i, j int) bool {
	if i < j {
		return true
	}
	return false
}
func exchange(i, j int, data []int) {
	data[i], data[j] = data[j], data[i]
}

func InsertionSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i; j > 0 && less(data[j], data[j-1]); j-- {
			exchange(j, j-1, data)
		}
	}
}
