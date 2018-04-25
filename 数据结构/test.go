package main

import "fmt"

type List struct {
	t    interface{}
	next *List
}
type Queue struct {
	start, end *List
	length     uint
}

func main() {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	fmt.Println(q)
	q.forEach(printList)
}
func printList(l *List) {
	fmt.Println(l)
}
func NewQueue() *Queue {
	return &Queue{}
}
func (q *Queue) forEach(f func(*List)) {
	if q.length == 0 {
		return
	}
	node := q.start
	for node.next != nil {
		f(node)
		node = node.next
	}
	f(node)
}
func (q *Queue) Push(t interface{}) bool {

	l := &List{t: t}
	if q.length == 0 {
		q.start, q.end = l, l
		q.length++
		return true
	}
	q.end.next = l
	q.end = l
	q.length++
	return true
}
func (q *Queue) Pop() (interface{}, bool) {
	if q.length == 0 {
		return nil, false
	}
	temp := q.start
	q.start = q.start.next
	q.length--
	return temp, true
}
func (q *Queue) isEmpty() bool {
	return q.length == 0
}
func (q *Queue) Size() uint {
	return q.length
}
