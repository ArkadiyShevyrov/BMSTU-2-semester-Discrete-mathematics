package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int32{0, 1}
	b := []int32{0, 1}
	p := 100
	fmt.Println(add(a, b, p))
}
func add(a, b []int32, p int) []int32 {
	n := int(math.Max(float64(len(a)), float64(len(b))))
	arrayA := make([]int32, 0, n)
	for _, value := range a {
		arrayA = append(arrayA, value)
	}
	for len(arrayA) < n {
		arrayA = append(arrayA, 0)
	}

	arrayB := make([]int32, 0, n)
	for _, value := range b {
		arrayB = append(arrayB, value)
	}
	for len(arrayB) < n {
		arrayB = append(arrayB, 0)
	}
	c := make([]int32, n+1)
	transfer := 0
	var i int
	for i = 0; i < n; i++ {
		s := arrayA[i] + arrayB[i] + int32(transfer)
		if s < int32(p) {
			c[i] = s
			transfer = 0
		} else {
			c[i] = s - int32(p)
			transfer = 1
		}
	}
	if transfer == 1 {
		c[i] = 1
		return c
	} else {
		return c[:i]
	}
}
