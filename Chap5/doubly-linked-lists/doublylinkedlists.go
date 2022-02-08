package main

import "fmt"

type Node struct{
	Value int
	Previous *Node
	Next *Node
}
var root = new(Node)

func addNode(val int, t *Node) int{
	if root == nil {
		t = &Node{val , nil, nil}
		root = t
		return 0
	}

	if t.Value == val {
		fmt.Println("Item already exists")
		return -1
	}

	if t.Next == nil {
		temp := t 
		t.Next = &Node{val, temp, nil}
		return 2
	}

	return addNode(val, t.Next)
}

func traverse(t *Node){
	if t == nil {
		fmt.Println("-> Empty List")
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func reverse(t *Node){
	if t == nil {
		fmt.Println("-> Empty List")
	}

	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}

	for temp.Previous != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}
	fmt.Printf("%d -> ", temp.Value)
	fmt.Println()
}

func size(t *Node) int{
	if t == nil {
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func lookUp(val int, t *Node) bool{
	if t == nil {
		return false
	}

	if val == t.Value {
		return false
	}

	if t.Next == nil {
		return false
	}

	return lookUp(val, t.Next)
}
func main(){
	fmt.Println(root)
	traverse(root)
	addNode(1, root)
	addNode(1, root)
	traverse(root)
	addNode(10, root)
	addNode(2, root)
	traverse(root)
	fmt.Println(lookUp(100, root))
	fmt.Println(size(root))
	reverse(root)
}