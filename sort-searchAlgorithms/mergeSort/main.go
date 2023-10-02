package main

import "fmt"

func merge(a,b[]int)[]int{
	for i:=0;i<len(a);i++{
		b=oneMerge(b,a[i])
	}
	return b
}

func oneMerge(array []int, x int) []int {
    i:=binarySearch(array,x)

    result := make([]int, len(array)+1)
    copy(result[:i], array[:i])   // Copy elements before insertion point
    result[i] = x                 // Insert the new element
    copy(result[i+1:], array[i:]) // Copy elements after insertion point

    return result
}

func binarySearch(array []int,x int)int{
	
	start:=0
	end:=len(array)-1
	for start<=end{	
		m:=(start+end)/2
		if array[m]<x{
			start=m+1
		} else{
			end=m-1
		}
	}
	return start
}

func mergeSort(a []int)[]int {
	n:=len(a)
	if len(a)>1{
		x:=mergeSort(a[0:n/2])
		y:=mergeSort(a[n/2:])
	return merge(x,y)
	}
	return a
}

func main(){
	a:=[]int{0,8,4,1,23,45,0,-9,5}
	c:=mergeSort(a)
	fmt.Println(c)
}