// This files will cover the Loops and the Arrays and Slices
package main

// import "fmt"

// func printSlice(x []int){
// 	for _, number := range x{
// 		fmt.Print(number, " ")
// 	}
// 	fmt.Println()
// }

// func main(){
// 	// The Go lang doesn't cover the while loop
// 	// You can do it with just the for loop

// 	// Simple for loop
// 	// for i := 0; i < 10; i++ {
// 	// 	if i/2 == 0{
// 	// 		continue
// 	// 	}
// 	// 	fmt.Println(i)
// 	// }

// 	// While loop
// 	// fmt.Println()
// 	// i := 10
// 	// for {
// 	// 	if i < 0{
// 	// 		break
// 	// 	}

// 	// 	fmt.Print(i, " ")
// 	// 	i--
// 	// }
// 	// fmt.Println()

// 	// Arrays
// 	// fmt.Println("=== Arrays ===")
// 	// name_of_array := [Array_Size]ArrayType{elem0, elem1, ...}
// 	// anArray := [5]int{0, -2, -4, 3, 4}
// 	// for i := 0; i < len(anArray); i++ {
// 	// 	fmt.Printf("Elem %d of value: %d \n",i,anArray[i])
// 	// }
// 	// for i, value := range anArray {
// 	// 	fmt.Printf("Elem %d of value: %d \n",i,value)
// 	// }

// 	// fmt.Println("=== Slices ===")
// 	// Same ass arrays way of declaration
// 	// aSlice := []int{1,2,3,4,5}
// 	// Zero value of a slices is nil
// 	// integers := make([]int,20)
// 	// fmt.Println(len(integers))

// 	// Temporary, i.e doesnt edit the slice
// 	// integers = append(integers, 23)
// 	// for i := 0; i < len(integers); i++ {
// 	// 	fmt.Println(integers[i])
// 	// }

// 	//
// 	// s2 := integers[1:4]
// 	// fmt.Println(s2)

// 	// Slices expands automatically

// 	// aSlice := []int{-1, 0, 4}
// 	// fmt.Printf("aSlice: ")
// 	// printSlice(aSlice)
// 	// fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

// 	// aSlice = append(aSlice, -100)

// 	// fmt.Printf("aSlice: ")
// 	// printSlice(aSlice)
// 	// fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

// 	// aSlice = append(aSlice, 23)
// 	// aSlice = append(aSlice, 21)
// 	// aSlice = append(aSlice, 25)
// 	// aSlice = append(aSlice, 26)
// 	// aSlice = append(aSlice, 27)
// 	// aSlice = append(aSlice, 28)
// 	// aSlice = append(aSlice, 29)
// 	// aSlice = append(aSlice, 30)
// 	// aSlice = append(aSlice, 31)

// 	// fmt.Printf("aSlice: ")
// 	// printSlice(aSlice)
// 	// fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

// }