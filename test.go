package main

import (
	"fmt"
	// "sort"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	// c := []int{23,50,10,99,18,23,98,84,11,10,48,77,13,54,98,77,77,77,68}
	n, _ := ioutil.ReadAll(os.Stdin)
	stackString := NewStack(10)
	stackInts := NewStack(10)
	for _, v := range n {
		switch string(v) {
		case "(", " ":
			continue
		case "+", "-", "*", "/":
			stackString.Push(v)
		case ")":
			ops := stackString.Pop().(byte)
			v := stackInts.Pop().(float64)
			switch string(ops) {
			case "+":
				v = stackInts.Pop().(float64) + v
			case "-":
				v = stackInts.Pop().(float64) - v
			case "*":
				v = stackInts.Pop().(float64) * v
			case "/":
				v = stackInts.Pop().(float64) / v
			}
			stackInts.Push(v)
		default:
			v, _ := strconv.ParseFloat(string(v), 64)
			stackInts.Push(v)
		}
	}
	fmt.Println(stackInts.Pop())
}

type Stack struct {
	a    []interface{}
	size uint
}

func NewStack(size uint) *Stack {
	return &Stack{a: make([]interface{}, size), size: 0}
}
func (s *Stack) Push(content interface{}) {
	s.a[s.size] = content
	s.size = s.size + 1
	// s.a[size++] = a
}
func (s *Stack) Pop() interface{} {
	s.size = s.size - 1
	return s.a[s.size]
}
func (s *Stack) Size() int {
	return int(s.size)
}
func (s *Stack) isEmpty() bool {
	return s.size == 0
}
