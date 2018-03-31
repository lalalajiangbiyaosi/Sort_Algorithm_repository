package main

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}
func main() {
	const n = 10000
	rightmost := make(chan int)
	right := rightmost
	left := rightmost
	for i := 0; i < n; i++ {
		left = make(chan int)
		go f(left, right)
		right = left
	}
	go func(c chan int) {
		c <- 1
	}(rightmost)
	fmt.Println(<-right)
}
// func main() {
// 	const n = 10000
// 	leftmost := make(chan int)
// 	right := leftmost
// 	left := leftmost
// 	for i := 0; i < n; i++ {
// 		right = make(chan int)
// 		go f(left, right)
// 		left = right
// 	}
// 	go func(c chan int) { c <- 1 }(right)
// 	fmt.Println(<-leftmost)
// }
