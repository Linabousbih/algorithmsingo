package main

import ("fmt")

type Heap []int


func (h *Heap) Create() {
    test := true
    array := *h
    for test {
        test = false
        for i := 0; i < len(array)/2; i++ {
            leftChildIdx := 2*i + 1
            rightChildIdx := 2*i + 2

            if leftChildIdx >= len(array) {
                return
            }

            // Check if rightChildIdx is within bounds before accessing it
            if rightChildIdx < len(array) && array[rightChildIdx] < array[i] {
                if array[leftChildIdx] < array[rightChildIdx] {
                    swap(&array, leftChildIdx, i)
                    test = true
                } else {
                    swap(&array, rightChildIdx, i)
                    test = true
                }
            } else if array[leftChildIdx] < array[i] {
                swap(&array, leftChildIdx, i)
                test = true
            }

        }
        // fmt.Println(array)
    }
}


func (h* Heap) Insert(val int){
	*h=append(*h,val)
	array:=*h
	if len(array)==1{
		return
	}
	i:=len(array)-1
	test:=true
	for test && i>=0{
		test=false
		parent:=((i-1)/2)
		if array[parent]>array[i]{
			swap(&array,parent,i)
			i=parent
			test=true
		}
	}
	
}

func swap(array *Heap, a,b int){
	tmp:=(*array)[a]
	(*array)[a]=(*array)[b]
	(*array)[b]=tmp
}

func (h *Heap) Remove(index int){
	array:=*h
	*h=append(array[:index],array[index+1:]...)
	// fmt.Println(*h)
	h.Create()
}

func (h* Heap) RemoveValue(val int){
	array:=*h
	i:=0
	for i<len(array){
		if array[i]==val{
			h.Remove(i)
			return
		}
		i++
	}
}

func isMinHeap(heap Heap) bool {
	for i := 1; i < len(heap); i++ {
		parentIdx := (i - 1) / 2
		if parentIdx < 0 {
			return true
		}

		if heap[parentIdx] > heap[i] {
			return false
		}
	}
	return true
}

func main(){
	array:=[]int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41}
	heap:=Heap(array)
	heap.Create()
	fmt.Println(heap)
	heap.Insert(9)
	// fmt.Println(heap)
	heap.Remove(5)
	// fmt.Println(heap)
	// fmt.Println(heap[6],"  ",heap[6*2+1])
	a:=[]int{}
	ndHeap:=Heap(a)
	// fmt.Println("Second heap : ",ndHeap)
	ndHeap.Insert(48)
	ndHeap.Insert(12)
	ndHeap.Insert(24)
	ndHeap.Insert(7)
	ndHeap.Insert(8)
	ndHeap.Insert(-5)
	ndHeap.Insert(24)
	ndHeap.Insert(391)
	ndHeap.Insert(41)
	fmt.Println(ndHeap)
	heap.RemoveValue(12)
	fmt.Println(heap)
	b:=[]int{48, 12, 24, 7, 8, -5, 24}
	thirdHeap:=Heap(b)
	fmt.Println(isMinHeap(thirdHeap))
}