package main

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node)DFS(array[]string) []string{
	current:=n
	array=append(array, current.Name)
		for _,child:=range current.Children{
				array=child.DFS(array)
			}
	return array
}