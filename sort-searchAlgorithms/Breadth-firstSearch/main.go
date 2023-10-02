package main



type Node struct {
	Name     string
	Children []*Node
}

func (n *Node)bfs() []string{
	current:=n 
	queue:=make([]*Node,0)
	result:=make([]string,0)
	queue=append(queue, current)
	
	for len(queue)>0{
		if len(queue[0].Children)>0{
			for i:=0;i<len(queue[0].Children);i++{
				queue=append(queue,queue[0].Children[i])
			}
		}
		result=append(result,dequeue(&queue))
	}
	return result

	}

func dequeue(q *[]*Node) string{
	v:=(*q)[0].Name
	(*q)=(*q)[1:]
	return v
}
func main(){

}