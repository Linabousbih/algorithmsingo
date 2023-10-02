package main

// import (
// 	"fmt"
// )

// type Node struct {
// 	Val      string
// 	Children []*Node
// }

// type Trie struct {
// 	root *Node
// }

// func (n *Node) DFS() {
// 	n.dfs(0)
// }

// func (n *Node) dfs(level int) {
// 	indent := ""
// 	for i := 0; i < level; i++ {
// 		indent += "  "
// 	}
// 	fmt.Printf("%s%s\n", indent, n.Val)

// 	for _, child := range n.Children {
// 		child.dfs(level + 1)
// 	}
// }

// func (trie *Trie) Insert(s string) {
// 	parent := trie.root
// 	for i := 0; i < len(s); i++ {
// 		if index, ok := BFS(parent.Children, s[i]); ok {
// 			fmt.Println("I exists")
// 			parent = parent.Children[index]
// 		} else {
// 			fmt.Println("I don't exist")
// 			newNode := &Node{Val: string(s[i])}
// 			parent.Children = append(parent.Children, newNode)
// 			parent = newNode
// 		}
// 	}
// }

// func BFS(tree []*Node, c byte) (int, bool) {
// 	for i := 0; i < len(tree); i++ {
// 		if tree[i].Val == string(c) {
// 			return i, true
// 		}
// 	}
// 	return -1, false
// }

// func main() {
// 	t := &Trie{
// 		root: &Node{},
// 	}
// 	fmt.Println(t.root.Val,t.root.Children)
// 	t.Insert("and")
// 	// t.Insert("ant")
// 	// t.Insert("age")
// 	// t.Insert("agent")
// 	fmt.Println(t.root.Val,t.root.Children)
// 	t.root.DFS()
// }
