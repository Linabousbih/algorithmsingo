package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Append data to the linked list
func (ll *LinkedList) Append(data int) {
	newNodeAddress := &Node{val: data, next: nil}
	if ll.head == nil {
		ll.head = newNodeAddress
	} else {
		currentNode := ll.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNodeAddress
	}
}

// Display values of a linked list
func (ll *LinkedList) Display() {
	currentNode := ll.head
	for currentNode != nil {
		fmt.Printf("%v -> ", currentNode.val)
		currentNode = currentNode.next
	}
	fmt.Println("nil")
}

func main() {
	myLinkedList := new(LinkedList)
	myLinkedList.Append(5)
	myLinkedList.Append(6)
	myLinkedList.Append(1)
	myLinkedList.Display()

}
