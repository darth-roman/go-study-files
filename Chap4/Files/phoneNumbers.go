package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"regexp"
// )

// func findNumbers(input string) string{
// 	phoneGrammar := "(0[5|6|7]-[0-9][0-9]-[0-9][0-9]-[0-9][0-9]-[0-9][0-9])"
// 	matchMe := regexp.MustCompile(phoneGrammar)
// 	return matchMe.FindString(input)
// }
// func main(){
// 	args := os.Args
// 	if len(args) < 2 {
// 		fmt.Printf("warning: not enough args\n")
// 		os.Exit(1)
// 	}

// 	for _, filename := range args[1:] {
// 		f, err := os.Open(filename)
// 		if err != nil {
// 			fmt.Printf("error: error opening the file, %s", err)
// 			os.Exit(1)
// 		}
// 		defer f.Close()

// 		reader:= bufio.NewReader(f)
// 		for {
// 			line, err := reader.ReadString('\n')
// 			if err == io.EOF {
// 				break
// 			}else if err != nil {
// 				fmt.Printf("error: error reading the file, %s", err)
// 				break
// 			}

// 			phone := findNumbers(line)
// 			fmt.Println(phone)
// 		}

// 	}
// }