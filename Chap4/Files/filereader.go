package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// // Select a particular column from a line of text
// // Read a text from a text file line by line
// /**
// * Two CLI arguments to operate.
// * 1st arg : A required column number
// * 2nd arg : File path to process
//  */

// func main(){
// 	arguments := os.Args
// 	if len(arguments) < 2{
// 		fmt.Println("usage: regex column <file1> [<file2> [<file2> [... <fileN>]]]")
// 		os.Exit(1)
// 	}
// 	// strconv.Atoi() is for parsing the integer from the command line argument
// 	// of the start line
// 	temp, err := strconv.Atoi(arguments[1])

// 	// parsing went wrong
// 	if err != nil {
// 		fmt.Println("Column value is not an integer:", temp)
// 		return
// 	}

// 	// The starting column line
// 	column := temp

// 	// negative number means no line valid obv
// 	if column < 0 {
// 		fmt.Println("Invalid Column number!")
// 		os.Exit(1)
// 	}

// 	// loop through the file lines
// 	for _, filename := range arguments[2:]{
// 		fmt.Println("\t\t", filename)
// 		f, err := os.Open(filename)
// 		if err != nil {
// 			fmt.Printf("error opening the file %s\n",err)
// 			continue
// 		}
// 		defer f.Close()
// 		r := bufio.NewReader(f)
// 		for {
// 			line, err := r.ReadString('\n')

// 			if err == io.EOF {
// 				break
// 			} else if err != nil {
// 				fmt.Printf("error reading file %s", err)
// 			}

// 			data := strings.Fields(line)
// 			if len(data) >= column {
// 				fmt.Println((data[column-1]))
// 			}
// 		}
// 	}

// }

// // Notes
// // /*
// // 	-Pattern matching is a technique for searching a string
// // 	for some set of characters based on specific search pattern
// // 	that is based on RegEx and grammars

// // 	- Can be used to replace and delete and manupilate the extracted
// // 	string/data

// // 	-Go package that is responsible for RegEx is "regexp"

// // */