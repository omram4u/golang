package testgitsrc

import "fmt"

type Node struct {
	next *Node
	prev *Node
	val  int
}
type Linklist struct {
	root *Node
	tail *Node
}

func (l *Linklist) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{val: val}
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev

}
func (l *Linklist) PrintNode() {
	x := l.root
	for x != nil {
		fmt.Printf("%v->", x.val)
		x = x.next
	}
	fmt.Println("hello")
}
func (l *Linklist) PrintRevers() {
	x := l.tail
	for x != nil {
		fmt.Printf("%v->", x.val)
		x = x.prev
	}
	fmt.Println("olleh")
}
func (l *Linklist) RemoveNode(node *Node) {
	x := node.prev
	if l.root == node {
		if l.root.next == nil {
			fmt.Println("one node")
			l.root = nil
			l.tail = l.root
			return
		}
		fmt.Println("hello root")
		l.root = node.next
		l.root.prev = nil
		return
	}

	if l.tail == node {
		//fmt.Println("hello tail")
		x.next = nil
		l.tail.prev = nil
		l.tail = x
		return
	}
	x.next = x.next.next
}
func Dprint() {
	fmt.Println("hello doblelinke")
}
