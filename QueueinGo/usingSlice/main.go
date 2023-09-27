package main

import "fmt"

func add(q []int, val int) []int {
	q = append(q, val)
	return q
}

func remove(q []int) []int {
	if len(q) == 0 {
		fmt.Println("Queue is empty")
		return q
	}
	return q[1:]
}


func main() {
	queue := []int{}
	queue = add(queue, 5)
	queue = add(queue, 4)
	queue = add(queue, 3)
	fmt.Println(queue)
	queue = remove(queue)
	fmt.Println(queue[0])
	queue = remove(queue)
	fmt.Println(queue)

}
