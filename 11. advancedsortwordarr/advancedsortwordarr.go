package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	size := 0
	for i := range array {
		size = i + 1
	}
	var temp string
	for i := 0; i < size-1; i++ {
		for j := 0; j < (size - i - 1); j++ {
			if f(array[j], array[j+1]) > 0 {
				temp = array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}
