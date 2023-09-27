package main

import "fmt"

type Node struct{
	num int
	next *Node
}

type Pile struct{
	head *Node
}

func (pile *Pile) empiler(new int){
	newNode:=&Node{num:new}
	if pile.head==nil{
		pile.head=newNode
	} else {
		newNode.next=pile.head
		pile.head=newNode	
	}
	
}

func (pile *Pile) depiler() int{
	var depiled int
	if pile.head!=nil{
		depiled=pile.head.num
		pile.head=pile.head.next
	}
	return depiled
}

func (pile *Pile) show() {
    current := pile.head
    for current != nil {
        fmt.Printf("%v -> ", current.num)
        current = current.next
    }
    fmt.Println("nil")
}
func main(){
    myStack := &Pile{}
    myStack.empiler(2)
    myStack.empiler(4)
    myStack.empiler(5)
    myStack.empiler(3)
    myStack.show()
	myStack.depiler()
	myStack.depiler()
    myStack.show()
}

