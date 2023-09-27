package main

import "fmt"

type Node struct{
	value int
	Next *Node
}

type Queue struct{
	Head *Node
	Tail *Node
}

func (queue *Queue) Enqueue(val int){
	newElement:=&Node{value: val, Next: nil}
	if queue.Head==nil{
		queue.Head=newElement
		queue.Tail=newElement
	} else {
		newElement.Next=queue.Head
		queue.Head=newElement
	}
}

func (queue *Queue) Dequeue() int{
	if queue==nil{
		fmt.Println("Queue is empty")
		return 0
	}
	dequeued:=queue.Tail.value
	current:=queue.Head
	for current.Next!=queue.Tail{
		current=current.Next
	}
	current.Next=nil
	queue.Tail=current
	
	return dequeued
}

func (queue *Queue)Show(){
	current:=queue.Head
	for current!=nil{
		fmt.Print(current.value,"->")
		current=current.Next
	}
	fmt.Println("nil")
}

func main(){
	q:=new(Queue)
	q.Enqueue(8)
	q.Enqueue(7)
	q.Enqueue(6)
	// fmt.Println(q.Tail.value)
	v:=q.Dequeue()
	fmt.Println(v)
	q.Show()	
	q.Enqueue(4)
	q.Enqueue(5)
	q.Show()	
}