/*
	创建一个N x N的布尔数组，其中当i和j互质时，a[i][j]为true
	func gcd(p int,q int) int {  
		使用2300多年前的欧几里得算法求解两数的最大公约数
		@param p 数一
		@param q 数二
		@return  最大公约数
	}
*/



package main

import (
	"fmt"
)

func main() {
	a := boolArray(5)
	fmt.Println(a)
}
func gcd(p int, q int) int {
	if q == 0 {
		return p
	}
	r := p % q
	return gcd(q, r)
}
func boolArray(N int) [][]bool {
	resu := make([][]bool, N)
	for row := 1; row < N+1; row++ {
		line := make([]bool, N)
		for col := 1; col < N+1; col++ {
			if gcd(row, col) == 1 {
				line[col-1] = true
			}
		}
		resu[row-1] = line
	}
	return resu
}
