package main

import (
	"fmt"
)

func main() {
	input := []int{58, 43, 26, 11, 18, 55, 34, 90, 100, 12, 500, 10000, 78, 29, 103, 1, 32, 8}
	B := Insert_sort(input)
	C := Selection_sort(input)
	fmt.Println(B)
	fmt.Println(C)
}
func Insert_sort(A []int) []int {

	for j := 1; j < len(A); j++ {
		i := j - 1
		for i >= 0 && A[i] > A[i+1] {
			A[i], A[i+1] = A[i+1], A[i]
			i = i - 1
		}
	}
	return A
}
func Selection_sort(A []int) []int {
	minVal := 0
	for j := 0; j < len(A); j++ {
		minVal = A[j]
		minValIndex := j
		i := j + 1
		for i < len(A) {
			if A[i] < minVal {
				minVal = A[i]
				minValIndex = i

			}
			i = i + 1E
		}
		if minValIndex != j {
			A[j], A[minValIndex] = A[minValIndex], A[j]
		}
	}
	return A
}
