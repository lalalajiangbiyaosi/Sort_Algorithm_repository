/*
	func IndexOf 为 BinarySearch
	func deleteSameInts 删去白名单中的所有重复元素

*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func main() {
	c := []int{23, 50, 10, 99, 18, 23, 98, 84, 11, 10, 48, 77, 13, 54, 98, 77, 77, 77, 68}
	sort.Ints(c)
	fmt.Println(c)
	fmt.Println(IndexOf(77, c))
	c = deleteSameInts(c)
	fmt.Println(IndexOf(77, c))
}
func readInts(r io.Reader) (p []int, err error) {
	inputReader := bufio.NewReader(r)
	for {
		line, err := inputReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimRight(line, "\n")
		resu, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		p = append(p, resu)
	}
	return p, nil
}
func IndexOf(n int, codeslice []int) (index int) {
	var lo int
	var hi = len(codeslice) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if n > codeslice[mid] {
			lo = mid + 1
		} else if n < codeslice[mid] {
			hi = mid - 1
		} else {
			return mid
		}

	}
	return -1
}

func deleteSameInts(p []int) []int {
	var pointer int
	resu := []int{p[pointer]}
	for _, v := range p {
		if v == p[pointer] {
			continue
		} else {
			resu = append(resu, v)
			pointer++
		}
	}
	return resu
}
