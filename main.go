package main

import (
	"fmt"
)

func main() {
	poly1 := []int{0, 1, 3, 1, 2, 1}
	poly2 := []int{8, 4, 2, 4, 1, 2, 4, 5, 2, 0, 1, 0, 0, 1}
	polyResult := multiply(poly1, poly2)
	fmt.Println(polyResult)
}

func multiply(poly1 []int, poly2 []int) []int {
	result := make([]int, len(poly1)+len(poly2))
	for i, coefficient := range poly1 {
		for j, coefficient2 := range poly2 {
			result[i+j] = coefficient * coefficient2
		}
	}
	return result
}
