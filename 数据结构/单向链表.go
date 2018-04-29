package singleLink

type node struct {
	t    interface{}
	next *node
}

type SingleLink struct {
	count int
	pnext *node
}

func NewSingleLink(t interface{}) *SingleLink {
	// nd := &node{t: t}
	// return &SingleLink{count: 1, pnext: nd}
	return &SingleLink{count: 0}
}
func (sl *SingleLink) Size() int {
	return sl.count
}
func (sl *SingleLink) isEmpty() bool {
	return sl.count == 0
}
func (sl *SingleLink) getNode(index int) *node {
	if index > sl.count || index < 0 {
		return nil
	}
	prenode := sl.pnext
	for temp := 0; temp < index; temp++ {
		prenode = prenode.next
	}
	return prenode
}
func (sl *SingleLink) Insert(index int, t interface{}) bool {
	prenode := sl.getNode(index - 1)
	if prenode != nil {
		newNode := &node{t: t, next: prenode.next}
		prenode.next = newNode
		sl.count++
		return true
	}
	return false
}
func (sl *SingleLink) Delete(index int) bool {
	if sl.isEmpty() {
		return false
	}
	prenode := sl.getNode(index - 1)
	delNode := prenode.next
	prenode.next = delNode.next
	sl.count--
	return true
}
