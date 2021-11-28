// author : hafizmfadli
// 29 November 2021

package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

// PushBack add element to end of linked list
func (l *LinkedList) PushBack(key interface{}) {
	node := &Node{
		prev: nil,
		next: nil,
		key:  key,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}
	l.length++
}

// Display print linked list element from head to tail
func (l *LinkedList) Display() {
	temp := l.head
	for {
		if temp != nil {
			fmt.Printf("%v ", temp.key)
			temp = temp.next
		} else {
			break
		}
	}
	fmt.Println()
}

// Reverse reverse linked list value
func (l *LinkedList) Reverse() {
	var keys []interface{}
	temp := l.head
	for {
		if temp != nil {
			keys = append(keys, temp.key)
			temp = temp.next
		} else {
			break
		}
	}
	temp = l.tail
	for _, v := range keys {
		if temp != nil {
			temp.key = v
			temp = temp.prev
		}
	}
}

func main() {
	myList := LinkedList{}
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	myList.PushBack(4)
	myList.PushBack(5)
	myList.PushBack(6)
	myList.Reverse()
	myList.Display()
	myList.Reverse()
	myList.Display()
	fmt.Println(myList.length)
}
