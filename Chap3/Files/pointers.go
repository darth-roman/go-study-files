package main

// Code segment for Pointers 1

import "fmt"

func getPointer(n *int){
	*n = *n * *n
}

func returnPointer(n int) *int{
	v := n*n
	return &v
}
func main(){

	// Part 2
	i := -10
	j := 25

	pI := &i
	pJ := &j
	
	fmt.Println("pI memory:", pI)
	fmt.Println("pJ memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pJ value:", *pJ)
	
	// Part 3
	*pI = 1234
	*pI--
	fmt.Println("i value:", i)
	
	// Part 4
	getPointer(pJ)
	fmt.Println("j value:", j)

	k:= returnPointer(12)
	fmt.Println(*k)
	fmt.Println(k)





}