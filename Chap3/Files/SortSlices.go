package main

// import (
// 	"fmt"
// 	"sort"
// )

// type aStructure struct{
// 	person string
// 	height int
// 	weight int
// }

// func main(){
// 	mySlice := make([]aStructure, 0)
// 	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
// 	mySlice = append(mySlice, aStructure{"Bill", 135, 45})
// 	mySlice = append(mySlice, aStructure{"John", 155, 45})
// 	mySlice = append(mySlice, aStructure{"George", 144, 50})
// 	mySlice = append(mySlice, aStructure{"David", 134, 40})

// 	fmt.Println("0:", mySlice)

// 	sort.Slice(mySlice, func(i, j int) bool{
// 		return mySlice[i].height < mySlice[j].height
// 	})

// 	fmt.Println("<:", mySlice)

// 	sort.Slice(mySlice, func(i, j int) bool{
// 		return mySlice[i].height > mySlice[j].height
// 	})

// 	fmt.Println(">:", mySlice)

// }