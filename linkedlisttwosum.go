package main

import "log"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(value int) {

	newNode := &Node{value: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}

	current = newNode

}

func (l *LinkedList) Print() {

	current := l.head
	for current != nil {
		log.Println(current.value)
		current = current.next
	}

	log.Println("still empty")
}

// func main() {

// 	ll := LinkedList{}
// 	ll.Print()
// }
