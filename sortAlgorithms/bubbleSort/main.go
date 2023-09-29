package main
import "fmt"

func sort(array[8]int) [8]int{
	swaps:=1
	for swaps >0{
		// fmt.Println("A new try now")
		swaps=0
		for i:=0;i<len(array)-1;i++{
			if array[i]>array[i+1]{
				swaps++
				tmp:=array[i]
				array[i]=array[i+1]
				array[i+1]=tmp
			}
		}
	}
	return array
}

func main(){
	array:=[8]int{5,6,0,3,2,4,9,-3}
	array=sort(array)
	fmt.Println(array)
}