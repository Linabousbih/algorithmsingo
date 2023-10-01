package main

import "fmt"

func sort (array [6]int) [6]int{
	j:=0
	var index int
	swap:=true
	for j<len(array) && swap{
		swap=false
		min:=array[j]
		for i:=j+1;i<len(array);i++{
			if array[i]<min{
				min=array[i]
				index=i
				swap=true
			}
		}
		if swap{
			array[index]=array[j]
			array[j]=min
		}
		j++
	}
	return array
}

func main(){
	array:=[6]int{20,12,10,15,2,9}

	array=sort(array)
	fmt.Println(array)
}