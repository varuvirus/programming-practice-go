package main

import (
    "fmt"
)

type listnode struct{
	data int
	next *listnode
}

type linkedlist struct {
	head *listnode
	length int
}

func (l *linkedlist) prepend(newnode *listnode) {
	newnode.next = l.head
	l.head = newnode
	l.length++
}

func (l *linkedlist) append (newnode *listnode) {
	if(l.head == nil) {
		l.head = newnode
	} else {	
		current := l.head
		for current != nil {
			if(current.next != nil) {
				current = current.next
			} else {
				current.next = newnode
				l.length++
				break 
			}
		}
	}
}

func (l *linkedlist) print(){
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}

	fmt.Println()
}

func (l *linkedlist) reverse() *linkedlist {
	current := l.head
	var previous *listnode
	var temp *listnode 

	for current != nil {
		temp = current.next
		current.next = previous
		previous = current 
		current = temp
	}

	return &linkedlist{head: previous}	
}

func main() {

	fmt.Println("Hello, World!")

	mylinkedlist := linkedlist{}

	node1 := &listnode{data:10}
	node2 := &listnode{data:5}
	node3 := &listnode{data:1}

	mylinkedlist.prepend(node1)
	mylinkedlist.prepend(node2)
	mylinkedlist.prepend(node3)
	
	node4 := &listnode{data:15}
	node5 := &listnode{data:20}
	node6 := &listnode{data:25}

	mylinkedlist.append(node4)
	mylinkedlist.append(node5)
	mylinkedlist.append(node6)

	fmt.Println("printing forward ... ")
	mylinkedlist.print()

	var reversedlistPointer *linkedlist = mylinkedlist.reverse()

	fmt.Println("printing reverse ... ")
	reversedlistPointer.print()

}
