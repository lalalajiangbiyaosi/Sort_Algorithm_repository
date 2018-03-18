package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{58, 43, 26, 11, 18, 55, 34, 90, 100, 12, 500, 10000, 78, 29, 103, 1, 32, 8}
	B := Insert_sort(input)
	C := Selection_sort(input)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println("into merge_sort!")
	fmt.Println(Merge_sort(input))
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
			i = i + 1
		}
		if minValIndex != j {
			A[j], A[minValIndex] = A[minValIndex], A[j]
		}
	}
	return A
}
func Merge_sort(A []int) []int {
	//	fmt.Println("Into mergo_sort!")
	if len(A) < 2 {
		return A
	}
	//	fmt.Println(A)
	X := A[0 : len(A)/2]
	Y := A[len(A)/2:]
	n := []int{}
	m := []int{}
	//	fmt.Println(X, Y)
	resultSlice := []int{}
	switch len(X) {
	case 1:
		n = X
	default:
		n = Merge_sort(X)
	}
	switch len(Y) {
	case 1:
		m = Y
	default:
		m = Merge_sort(Y)
	}
	i, j := 0, 0
	tempN := n[i]
	tempM := m[j]
	for k := 0; k < len(n)+len(m); k++ {

		if tempN < tempM {
			resultSlice = append(resultSlice, n[i])

			i = i + 1
			if i < len(n) {
				tempN = n[i]
			} else {
				tempN = math.MaxInt64
			}
		} else {
			resultSlice = append(resultSlice, m[j])
			j = j + 1
			if j < len(m) {
				tempM = m[j]
			} else {
				tempM = math.MaxInt64
			}
		}
	}
	return resultSlice
	//	fmt.Println(X)
	//	fmt.Println(Y)
	//	return A
}
