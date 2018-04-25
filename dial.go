package main

import (
	"sort"
	// "errors"
	"bufio"
	"fmt"
	"os"
	// "sort"
	"io"
	"strconv"
	"strings"
)

func main() {

	// slice := []int{50,60,40,50,30,55,66,77,99,65,79,37}
	// sort.Ints(slice)
	// fmt.Println(slice)
	// fmt.Println(IndexOf(99,slice))
	var l float64 = 1.0
	var n float64 = 0
	c := l / n
	fmt.Println(c)
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

	// p, err := readInts(os.Stdin)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// sort.Ints(p)
	// fmt.Println(p)

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
	return findEdge(p, key, tempN, 1) - findEdge(p, key, tempN, -1)
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
