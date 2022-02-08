package main

import (
	"fmt"
)
type Node struct{
	Value int
	Next *Node
}
var root = new(Node)

func addNote(t *Node, value int) int{
	if root == nil {
		t = &Node{value, nil}
		root = t
		return 0
	}
	if t.Value == value {
		fmt.Println("This value already exists")
		return -1
	}
	if t.Next == nil {
		t.Next = &Node{value, nil}
		return -2
	}
	return addNote(t,value)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty List!!")
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}

	fmt.Println()
}

func lookUp(t *Node, value int) bool {
	if root == nil {
		t = &Node{value, nil}
		root = t
		return false
	}
	if t.Value == value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookUp(t.Next, value)
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
func main(){

	fmt.Println(root)
	root = nil 
	addNote(root, 1)
	addNote(root, -1)
	traverse(root)
	addNote(root, 1)
	traverse(root)
	fmt.Println(size(root))

	if lookUp(root, -1) {
		fmt.Println("Node exists")
	}else{
		fmt.Println("Node DOES NOT exist")
	}

	if lookUp(root, 100) {
		fmt.Println("Node exists")
	}else{
		fmt.Println("Node DOES NOT exist")
	}
}