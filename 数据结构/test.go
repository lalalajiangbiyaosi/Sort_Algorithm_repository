package main

import (
	"fmt"
	"log"
)

type node struct {
	t        interface{}
	preNode  *node
	nextNode *node
}

type DoubleLink struct {
	phead  *node
	length int
}

func main() {
	db := NewDoubleLink()
	// db.Insertfront(1)
	// db.Insertlast(2)
	// db.Insert(3, 3)
	// db.Insertlast(4)
	// db.Insertlast(5)
	// db.Insert(99, 4)
	for i := 0; i < 10; i++ {
		db.Insertlast(i)
	}
	for i := 0; i < 5; i++ {
		db.Deletefront()
	}
	fmt.Println(db.phead)
	fmt.Println(db.phead.nextNode)
	fmt.Println(db.phead.nextNode.nextNode)
	fmt.Println(db.phead.nextNode.nextNode.nextNode)
	fmt.Println(db.phead.nextNode.nextNode.nextNode.nextNode)

}

func NewDoubleLink() *DoubleLink {
	return &DoubleLink{phead: &node{t: 0}, length: 0}
}
func (db *DoubleLink) Insertfront(t interface{}) bool {
	newNode := &node{t: t, preNode: db.phead, nextNode: db.phead}
	if db.phead.nextNode != nil {
		newNode.nextNode = db.phead.nextNode
		db.phead.nextNode.preNode = newNode
	}
	if db.phead.preNode == nil {
		db.phead.preNode = newNode
	}
	db.phead.nextNode = newNode
	db.length++
	return true
}
func (db *DoubleLink) Insertlast(t interface{}) bool {
	newNode := &node{t: t, nextNode: db.phead, preNode: db.phead}
	if db.phead.preNode != nil {
		newNode.preNode = db.phead.preNode
		db.phead.preNode.nextNode = newNode
	}
	if db.phead.nextNode == nil {
		db.phead.nextNode = newNode
	}
	db.phead.preNode = newNode
	db.length++
	return true
}

func (db *DoubleLink) Insert(t interface{}, index int) bool {
	if index > db.length {
		return db.Insertlast(t)
	}
	newNode := &node{t: t}
	oldNode, ok := db.getNode(index)
	fmt.Println("oldNode is ", oldNode)
	if !ok {
		log.Fatal("Fetching node geting wrong!")
	}
	newNode.nextNode, newNode.preNode = oldNode, oldNode.preNode
	oldNode.preNode.nextNode = newNode
	oldNode.preNode = newNode
	return true
}

func (db *DoubleLink) getNode(index int) (*node, bool) {
	if index > db.length || index < 0 {
		return nil, false
	}
	theNode := db.phead
	if index <= (db.length / 2) {
		for i := 0; i < index; i++ {
			theNode = theNode.nextNode
		}
		return theNode, true
	}
	index = db.length - index
	for index >= 0 {
		theNode = theNode.preNode
		index--
	}
	return theNode, true
}
func (db *DoubleLink) Deletefront() bool {
	if db.length == 0 {
		return false
	}
	db.phead.nextNode = db.phead.nextNode.nextNode
	db.phead.nextNode.preNode = db.phead
	db.length--
	return true
}
func (db *DoubleLink) Deletelast() bool {
	if db.length == 0 {
		return false
	}
	db.phead.preNode = db.phead.preNode.preNode
	db.phead.preNode.nextNode = db.phead
	db.length--
	return true
}
func (db *DoubleLink) isEmpty() bool {
	return db.length == 0
}
func (db *DoubleLink) Size() int {
	return db.length
}
