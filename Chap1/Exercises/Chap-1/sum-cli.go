package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	// Exercice One and Two :
	/*arguments := os.Args
	sum := 0.0
	average := 0.0
	for i := 1; i < len(arguments); i++ {
		//argument, _ := strconv.ParseInt(arguments[i],10, 64)
		floatingArg, _ := strconv.ParseFloat(arguments[i], 64)
		sum += floatingArg

	}
	average = sum/(float64(len(arguments))-1)
	fmt.Println(average) */

	// Exercise Three
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "END" {
			fmt.Println("End of the program")
			break
		}
		fmt.Println(">> "+scanner.Text())
	}
}

// func validator(){

// }