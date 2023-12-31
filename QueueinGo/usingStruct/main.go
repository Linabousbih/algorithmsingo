package main

import "fmt"

type Queue struct{
	Elements []int
	Size int
}


func (q *Queue) Enqueue(val int){
	if len(q.Elements)==q.Size{
		fmt.Println("The queue is full")
		return
	}
	q.Elements=append(q.Elements, val)
}

func (q* Queue) Dequeue() int{
	if len(q.Elements)==0{
		fmt.Println("The queue is empty")
		return 0
	}
	dequeued:=q.Elements[0]
	q.Elements=q.Elements[1:]
	return dequeued
}

func (q* Queue) show(){
	fmt.Println(q.Elements)
}

func main(){
	q:=Queue{Size: 5}
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(3)
	q.Enqueue(6)
	q.Enqueue(1)
	q.Enqueue(9)
	q.Enqueue(2)
	q.show()
	v:=q.Dequeue()
	fmt.Println(v)
	q.show()
	_=q.Dequeue()
	_=q.Dequeue()
}