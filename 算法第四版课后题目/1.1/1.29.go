/*
	func count(key int,a []int) int 计算a中key元素的数量
	func rank (key int,a []int) int 返回a中比key小的元素的数量
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	// "errors"
	"os"
	"strconv"
	"strings"
)

func main() {
	c := []int{23, 50, 10, 99, 18, 23, 98, 84, 11, 10, 48, 77, 13, 54, 98, 77, 77, 77, 68}
	fmt.Println(rank(77, c))
	fmt.Println(count(77, c))

	args := os.Args
	file, err := os.Open(args[1])
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fileOnenum, err := readInts(file)
	sort.Ints(fileOnenum)
	fileOnenum = deleteSameInts(fileOnenum)
	if err != nil {
		panic(err)
	}
	p, err := readInts(os.Stdin)
	if err != nil {
		panic(err)
	}
	for _, v := range p {
		if IndexOf(v, fileOnenum) < 0 {
			fmt.Println(v)
		}
	}
}

func rank(key int, p []int) int {
	sort.Ints(p)
	tempN := IndexOf(key, p)
	return findEdge(p, key, tempN, -1)
}
func findEdge(p []int, key int, index int, step int) int {
	for {
		if p[index+step] != key {
			return index
		}
		index = index + step
	}
}
func count(key int, p []int) int {
	sort.Ints(p)
	tempN := IndexOf(key, p)
	return findEdge(p, key, tempN, 1) - findEdge(p, key, tempN, -1) + 1
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
