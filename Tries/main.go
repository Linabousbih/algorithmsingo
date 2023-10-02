package main

import "fmt"

// import "fmt"

type TrieNode struct {
	Children map[rune]*TrieNode
	isEnd bool
}

type Trie struct{
	root *TrieNode
}

func newTrie() Trie{
	root:=TrieNode{Children:make(map[rune]*TrieNode)}

	return Trie{root: &root}
}

func (t* Trie) Insert(word string){
	current:=t.root
	for _,char:=range word{
		if current.Children[char]==nil{
			current.Children[char]=&TrieNode{Children:make(map[rune]*TrieNode)}
		}
			current=current.Children[char]
	}
	current.isEnd=true
	}
func (t*Trie) Search(word string) bool{
	current:=t.root
	for _,char:=range word{
		if current.Children[char]==nil{
			return false
		}
			current=current.Children[char]
	}
	return current.isEnd
	}

func (t*TrieNode) Display(){
	current:=t
	if current.isEnd{
		fmt.Println()
		fmt.Print(" ")
	}
	for c,char:=range current.Children{
		fmt.Print(string(c))
		char.Display()
	}
}

func main(){
	t:=newTrie()
	t.Insert("word")
	t.Insert("wo")
	fmt.Println(t.Search("wo"))
	t.Insert("age")
	t.Insert("agent")
	t.Insert("and")
	t.Insert("ant")
	t.root.Display()
}