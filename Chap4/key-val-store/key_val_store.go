/*
	-The program is a simple key-value store
	-Will have some simple actions (commands) = ADD, DELETE, LOOKUP and CHANGE
	-STOP will exit the program
	-PRINT will print all store items
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Element struct {
	Name string
	Surname string
	Id string
}

var DATA = make(map[string]Element)

func ADD(k string, elem Element) bool{
	if k == "" {
		return false
	}

	if LOOKUP(k) == nil {
		DATA[k] = elem
		return true
	}

	return false
}

func LOOKUP(k string) *Element{
	_, target := DATA[k]
	if target {
		found := DATA[k]
		return &found
	}
	return nil
}

func DELETE(k string) bool{
	if LOOKUP(k) != nil{
		delete(DATA, k)
		return true
	}
	return false
}

func CHANGE(k string, edit Element) bool {
	if LOOKUP(k) != nil {
		DATA[k] = edit
		return true		
	}
	return false
}

func PRINT(){
	for key, value := range DATA {
		fmt.Printf("Key: %s; Value: %v\n", key, value)
	}
}
func main(){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)
		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0]{
		case "ADD":
			element := Element{tokens[2], tokens[3], tokens[4]}
			if !ADD(tokens[1], element) {
				fmt.Printf("error adding the element\n")
			}
		case "DELETE":
			if !DELETE(tokens[1]) {
				fmt.Printf("error deleting the element\n")
			}
		case "LOOKUP":
			element := LOOKUP(tokens[1])
			if element == nil {
				fmt.Printf("no element found\n")
			}else{
				fmt.Printf("%v\n", *element)
			}
		case "CHANGE":
			newElement := Element{tokens[2], tokens[3], tokens[4]}
			if !CHANGE(tokens[1], newElement) {
				fmt.Printf("error editting the element\n")
			}
		case "PRINT":
			PRINT()
		case "STOP":
			return
		default:
			fmt.Printf("unknown command - error\n")
		}

	}
}