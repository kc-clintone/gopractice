package piscine

// import (
// 	"sort"
// )

// func Abort(a, b, c, d, e int) int {
// 	num := []int{a, b, c, d, e}
// 	sort.Ints(num)
// 	return num[2]
// }

func Abort(a, b, c, d, e int) int {
	array := []int{a, b, c, d, e}
	counter := 5
	temp := 0

	for i := 0; i < counter - 1; i++ {
		temp = i
		for j := i + 1; j < counter; j++ {
			if array[j] < array[temp] {
				temp = j
			}
		}
		array[i], array[temp] = array[temp], array[i]
	}
	return array[2]
}