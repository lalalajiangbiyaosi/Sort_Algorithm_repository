package main

import (
	"fmt"
)

var (
	deposits = make(chan int)
	balances = make(chan int)
)

func Deposit(mount int) {
	deposits <- mount
}
func Balance() int {
	return <-balances
}
func teller() {
	var deposit int
	for {
		select {
		case mount := <-deposits:
			deposit += mount
		case balances <- deposit:

		}
	}
}
func init() {
	go teller()
}
func main() {
	Deposit(100)
	Deposit(200)
	fmt.Println(Balance())
}
