package main

import (
	"fmt"
)

var array = []int{
	10, -64, -127, -36, -155, 29, -100, -114, -50, 2,
	-9, 16, -166, -21, -78, -58, -73, 35, -121, -97,
	-73, -27, -149, -160, -137, -8, -46, -95, -107, -125,
	-98, 21, 27, 30, -57, -8, -13, -142, -87, 13,
	-175, -167,
}

func main() {
	fmt.Println(array)
	qsort(len(array), less, swap)
	fmt.Println(array)
}

func qsort(n int,
	less func(i, j int) bool,
	swap func(i, j int)) {
	quicksort(0, n-1, less, swap)
}

func quicksort(low int, high int,
	less func(i, j int) bool,
	swap func(i, j int)) {
	if low < high {
		p := partition(low, high, less, swap)
		quicksort(low, p, less, swap)
		quicksort(p+1, high, less, swap)
	}
}

func partition(low int,
	high int,
	less func(i, j int) bool,
	swap func(i, j int)) int {
	i := low - 1

	for j := low; j < high; j++ {
		if !less(high, j) {
			i++
			swap(i, j)
		}
	}
	swap(i+1, high)
	return i
}

func less(i, j int) bool {
	return array[i] < array[j]
}

func swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}
