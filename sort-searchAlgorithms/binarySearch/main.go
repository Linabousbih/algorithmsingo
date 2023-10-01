package main

import "fmt"


func binSearch(array []int, x int) bool {
	
	if len(array) == 0 {
		return false
	}
	
	mid := len(array) / 2
	
	if array[mid] == x {
		return true
	} else if x < array[mid] {
		return binSearch(array[:mid], x)
	} else {
		return binSearch(array[mid+1:], x)
	}
}

func binSearchIndex(array []int, x int)int{
	low:=0
	high:=len(array)-1
	for low<=high{
		// fmt.Println(low,high)
		if array[(low+high)/2]==x{
			return (low+high)/2
		} else if x>array[(low+high)/2]{
			low=((low+high)/2)+1
		} else if x<array[(low+high)/2]{
			high=((low+high)/2)-1
		}
	}
	return -1
}

func main(){
	arr:=[]int{3,4,7,8,9,10}
	ok:=binSearch(arr,9)
	i:=binSearchIndex(arr,10)
	fmt.Println(ok,i)
}