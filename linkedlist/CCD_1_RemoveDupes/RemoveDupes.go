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

func (l *linkedlist) append(newnode *listnode) {
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

func (l *linkedlist) removeDupes(){
	
}

func main() {
	fmt.Println("CCD Linkedlist 1 - Remove Dupes")		
}