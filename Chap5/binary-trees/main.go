package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct{
	Left *Tree
	Value int
	Right *Tree
}

func traverse(tree *Tree) {
	if tree == nil {
		return
	}

	traverse(tree.Left)
	fmt.Println(tree.Value, " ")
	traverse(tree.Right)
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < n*2; i++ {
		temp := rand.Intn(n*2)
		t = insert(t, temp)
	}
	return t
}

func insert(t *Tree, value int) *Tree {
	if t == nil {
		return &Tree{nil, value, nil}
	}

	if value == t.Value {
		return t
	}

	if value < t.Value {
		t.Left = insert(t.Left ,value)
		return t
	}

	t.Right = insert(t.Right, value)
	return t
}

func main() {
	tree := create(40)
	fmt.Println("Value of the root is ", tree.Value)
	traverse(tree)	
}